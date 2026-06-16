package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
)

const (
	mediaStoragePath = "./storages/media"
	maxUploadSize    = 10 << 20 // 10 MB
)

var MediaControllerRouter = new(cMediaController)

type cMediaController struct{}

// @Summary Upload a media file
// @Description Upload a media file (image, video, document) associated with a message or conversation
// @Tags Media
// @Accept multipart/form-data
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param file formData file true "Media file"
// @Param user_id formData int true "User ID"
// @Param message_id formData int false "Message ID"
// @Param conversation_id formData int false "Conversation ID"
// @Success 201 {object} response.ResponseData{data=model.MediaOutput} "Media uploaded successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /media/upload [post]
func (c *cMediaController) UploadMedia(ctx *gin.Context) {
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxUploadSize)

	var params vo.UploadMediaRequest
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "file is required: "+err.Error())
		return
	}
	defer file.Close()

	mimeType := header.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	ext := filepath.Ext(header.Filename)
	storedName := fmt.Sprintf("%s%s", uuid.New().String(), strings.ToLower(ext))

	if err := os.MkdirAll(mediaStoragePath, os.ModePerm); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, "failed to create storage directory: "+err.Error())
		return
	}

	destPath := filepath.Join(mediaStoragePath, storedName)
	if err := ctx.SaveUploadedFile(header, destPath); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, "failed to save file: "+err.Error())
		return
	}

	accessURL := fmt.Sprintf("/ticket-communication/api/v1/media/files/%s", storedName)

	out, err := service.MediaService().CreateMedia(ctx, model.MediaInput{
		MessageID:      params.MessageID,
		ConversationID: params.ConversationID,
		UserID:         params.UserID,
		FileName:       storedName,
		OriginalName:   header.Filename,
		MimeType:       mimeType,
		Size:           header.Size,
		URL:            accessURL,
	})
	if err != nil {
		os.Remove(destPath)
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, out)
}

// @Summary Get media by ID
// @Description Get media metadata by ID
// @Tags Media
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Media ID"
// @Success 200 {object} response.ResponseData{data=model.MediaOutput} "Media retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Media not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /media/get-by-id/{id} [get]
func (c *cMediaController) GetMediaByID(ctx *gin.Context) {
	var params vo.MediaIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MediaService().GetMediaById(ctx, params.ID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get media by message ID
// @Description Get all media files associated with a message
// @Tags Media
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param message_id path int true "Message ID"
// @Success 200 {object} response.ResponseData{data=[]model.MediaOutput} "Media retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /media/message/{message_id} [get]
func (c *cMediaController) GetMediaByMessageID(ctx *gin.Context) {
	var params vo.MediaByMessageIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MediaService().GetMediaByMessageId(ctx, params.MessageID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get media by conversation ID
// @Description Get all media files associated with a conversation
// @Tags Media
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param conversation_id query int true "Conversation ID"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} response.ResponseData{data=[]model.MediaOutput} "Media retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /media/conversation [get]
func (c *cMediaController) GetMediaByConversationID(ctx *gin.Context) {
	var params vo.MediaByConversationIDRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MediaService().GetMediaByConversationId(ctx, params.ConversationID, params.Limit, params.Offset)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Delete media
// @Description Delete a media file by ID
// @Tags Media
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Media ID"
// @Success 204 {object} response.ResponseData "Media deleted successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Media not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /media/delete/{id} [delete]
func (c *cMediaController) DeleteMedia(ctx *gin.Context) {
	var params vo.MediaIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	media, err := service.MediaService().GetMediaById(ctx, params.ID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	filePath := filepath.Join(mediaStoragePath, media.FileName)
	if err := service.MediaService().DeleteMedia(ctx, params.ID); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	os.Remove(filePath)
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, nil)
}
