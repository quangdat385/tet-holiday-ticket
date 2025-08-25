package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
)

var MessageControllerRouter = new(cMessageController)

type cMessageController struct{}

// @Summary Get message by ID
// @Description Get message by ID
// @Tags Message
// @Accept json
// @Produce json
// @Param id path int true "Message ID"
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Success 200 {object} response.ResponseData{data=model.MessageOutput} "Message retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Message not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/get-by-id/{id} [get]
func (c *cMessageController) GetMessageByID(ctx *gin.Context) {
	var params vo.MessageIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MessageService().GetMessageById(ctx, params.ID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get messages by conversation ID
// @Description Get messages by conversation ID
// @Tags Message
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param conversation_id query int true "Conversation ID"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} response.ResponseData{data=[]model.MessageOutput} "Messages retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Messages not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/conversation/{conversation_id} [get]
func (c *cMessageController) GetMessagesByConversationID(ctx *gin.Context) {
	var params vo.MessagesByConversationIDRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MessageService().GetMessagesByConversationId(ctx, params.ConversationID, params.Limit, params.Offset)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get messages by user ID
// @Description Get messages by user ID
// @Tags Message
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param user_id query int true "User ID"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} response.ResponseData{data=[]model.MessageOutput} "Messages retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Messages not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/user [get]
func (c *cMessageController) GetMessagesByUserID(ctx *gin.Context) {
	var params vo.MessagesByUserIDRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MessageService().GetMessageByUserId(ctx, params.UserID, params.Limit, params.Offset)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Create a new message
// @Description Create a new message
// @Tags Message
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param message body model.MessageInput true "Message input"
// @Success 201 {object} response.ResponseData{data=model.MessageOutput} "Message created successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/create [post]
func (c *cMessageController) CreateMessage(ctx *gin.Context) {
	var input vo.CreateMessageRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.MessageService().CreateMessage(ctx, model.MessageInput{
		ConversationID: input.ConversationID,
		UserID:         input.UserID,
		Status:         input.Status,
		Message:        input.Message,
		Type:           input.Type,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, out)
}

// @Summary Update message status
// @Description Update message status
// @Tags Message
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Message ID"
// @Param payload body vo.UpdateMessageStatusRequest true "New status"
// @Success 201 {object} response.ResponseData{data=model.MessageOutput} "Message status updated successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Message not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/update/{id} [patch]
func (c *cMessageController) UpdateMessageStatus(ctx *gin.Context) {
	var params vo.MessageIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var body vo.UpdateMessageStatusRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if body.MessageID != params.ID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "Message ID mismatch")
		return
	}
	out, err := service.MessageService().UpdateMessageStatus(ctx, body.MessageID, body.UserID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, out)
}

// @Summary Delete a message
// @Description Delete a message
// @Tags Message
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Message ID"
// @Success 204 {object} response.ResponseData "Message deleted successfully"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Message not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /message/delete/{id} [delete]
func (c *cMessageController) DeleteMessage(ctx *gin.Context) {
	var params vo.MessageIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	err := service.MessageService().DeleteMessage(ctx, params.ID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, nil)
}
