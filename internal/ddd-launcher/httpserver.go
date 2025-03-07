package launcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	service "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-application-service"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-application-service/diary"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-application-service/stickynote"
	cache "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-cache"
	gateway "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-gateway"
	persistence "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/db"
	web "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-interface-web"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/ginx/middleware"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/log"
)

// @title           ddd-archetype-go
// @version         1.0
// @description     This is the open api doc for ddd-archetype-go

// @host      :5050
// @BasePath  /

func SetupHTTPServer(cfgPath string) {
	cfg := LoadConfig(cfgPath)

	mdb := InitMySQL(cfg.DB)
	rdb := InitRedis(cfg.RDB)

	server := &httpApi{}
	r := server.setupRouter(mdb, rdb)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           r.Handler(),
		ReadHeaderTimeout: time.Second * 5,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("error occurs in http server")
	}
}

type httpApi struct{}

func (h *httpApi) setupRouter(mdb *gorm.DB, rdb redis.Cmdable) *gin.Engine {
	r := gin.New()
	r.Use(middleware.WithRequestAndTrace(), gin.Recovery())

	diaryDB := db.NewDiaryDB(mdb)
	diaryRepo := persistence.NewDiaryEntityRepoImpl(diaryDB, cache.NewDiaryEntityCache(rdb))
	diaryCmdSvc := diary.NewDiaryCommandService(&diary.DiaryEntityFactoryImpl{}, diaryRepo)
	diaryQuerySvc := diary.NewDiaryQueryService(diaryDB)
	diary := web.NewDiaryController(diaryCmdSvc, diaryQuerySvc)

	openaiGateway := gateway.NewOpenAIGatewayImpl()
	stickyNoteDB := db.NewStickyNoteDB(mdb)
	stickyNoteGenSvc := service.NewGenerateContentDomainServiceImpl(openaiGateway, stickyNoteDB)
	stickyNoteRepo := persistence.NewStickyNoteEntityRepoImpl(stickyNoteDB, cache.NewStickyNoteEntityCache(rdb))
	stickyNoteCmdSvc := stickynote.NewStickyNoteCommandService(&stickynote.StickyNoteEntityFactoryImpl{}, stickyNoteRepo)
	stickyNoteQuerySvc := stickynote.NewStickyNoteQueryService(stickyNoteGenSvc)
	stickyNote := web.NewStickyNoteController(stickyNoteCmdSvc, stickyNoteQuerySvc)

	v1 := r.Group("/api/v1")
	{
		diaryV1 := v1.Group("/diary")
		diaryV1.POST("/create", diary.Create)
		diaryV1.POST("/pageList", diary.PageList)
	}

	{
		stickyNoteV1 := v1.Group("/stickynote")
		stickyNoteV1.POST("/create", stickyNote.Create)
		stickyNoteV1.POST("/modify", stickyNote.Modify)
		stickyNoteV1.POST("/generateDiaryContent", stickyNote.GenerateDiaryContent)
	}
	return r
}
