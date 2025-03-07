package web

import (
	"github.com/gin-gonic/gin"

	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-application-service/diary"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/response"
)

type DiaryController struct {
	cmdSvc   *diary.DiaryCommandService
	querySvc *diary.DiaryQueryService
}

func NewDiaryController(cmdSvc *diary.DiaryCommandService, querySvc *diary.DiaryQueryService) *DiaryController {
	return &DiaryController{cmdSvc: cmdSvc, querySvc: querySvc}
}

// Create godoc
// @Summary Create a diary
// @Description Create a diary
// @Tags diary service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param CreateCommand body diary.CreateCommand true "The command to create a diary"
// @Success 200 {object} diary.CreateView
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/diary/create [post]
func (c *DiaryController) Create(ctx *gin.Context) {
	var cmd diary.CreateCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		response.GinParamError(ctx, err)
		return
	}

	view, err := c.cmdSvc.CreateDiary(ctx, cmd)
	if err != nil {
		response.GinError(ctx, err)
		return
	}
	response.GinSuccess(ctx, view)
}

// PageList godoc
// @Summary list diaries
// @Description list diaries in page according to the last id and page size
// @Tags diary service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param Query body diary.Query true "The command to page list"
// @Success 200 {object} diary.QueryView
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/diary/pageList [post]
func (c *DiaryController) PageList(ctx *gin.Context) {
	var cmd diary.Query
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		response.GinParamError(ctx, err)
		return
	}

	view, err := c.querySvc.PageList(ctx, cmd)
	if err != nil {
		response.GinError(ctx, err)
		return
	}
	response.GinSuccess(ctx, view)
}
