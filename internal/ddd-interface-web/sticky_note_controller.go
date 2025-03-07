package web

import (
	"github.com/gin-gonic/gin"

	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-application-service/stickynote"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/response"
)

type StickyNoteController struct {
	cmdSvc   *stickynote.StickyNoteCommandService
	querySvc *stickynote.StickyNoteQueryService
}

func NewStickyNoteController(
	cmdSvc *stickynote.StickyNoteCommandService,
	querySvc *stickynote.StickyNoteQueryService,
) *StickyNoteController {
	return &StickyNoteController{
		cmdSvc:   cmdSvc,
		querySvc: querySvc,
	}
}

// Create godoc
// @Summary Create a sticky note
// @Description Create a sticky note
// @Tags sticky note service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param CreateCommand body stickynote.CreateCommand true "The command to create a sticky note"
// @Success 200 {object} stickynote.CreateView
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/stickynote/create [post]
func (c *StickyNoteController) Create(ctx *gin.Context) {
	var cmd stickynote.CreateCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		response.GinParamError(ctx, err)
		return
	}

	view, err := c.cmdSvc.CreateStickyNote(ctx, cmd)
	if err != nil {
		response.GinError(ctx, err)
		return
	}
	response.GinSuccess(ctx, view)
}

// Modify godoc
// @Summary Modify a sticky note
// @Description Modify a sticky note
// @Tags sticky note service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param ModifyCommand body stickynote.ModifyCommand true "The command to modify a sticky note"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/stickynote/modify [post]
func (s *StickyNoteController) Modify(ctx *gin.Context) {
	var cmd stickynote.ModifyCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		response.GinParamError(ctx, err)
		return
	}
	err := s.cmdSvc.ModifyStickyNote(ctx, cmd)
	if err != nil {
		response.GinError(ctx, err)
		return
	}
	response.GinSuccess(ctx, gin.H{})
}

// GenerateDiaryContent godoc
// @Summary Generate content for a diary
// @Description Generate content for a diary
// @Tags sticky note service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param GenerateContentCommand body stickynote.GenerateContentCommand true "The command to generate content for a diary"
// @Success 200 {object} stickynote.GenerateContentView
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/stickynote/generateDiaryContent [post]
func (s *StickyNoteController) GenerateDiaryContent(ctx *gin.Context) {
	var cmd stickynote.GenerateContentCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		response.GinParamError(ctx, err)
		return
	}
	view, err := s.querySvc.GenerateDiaryContent(ctx, cmd)
	if err != nil {
		response.GinError(ctx, err)
		return
	}
	response.GinSuccess(ctx, view)
}
