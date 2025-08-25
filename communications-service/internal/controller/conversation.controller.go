package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
)

var ConversationControllerRouter = new(cConversationController)

type cConversationController struct {
}

// @Summary Create a new conversation
// @Description Create a new conversation with the given parameters
// @Tags Conversation
// @Accept json
// @Produce json
// @Param request body vo.CreateConversationRequest true "Create conversation request"
// @Success      201  {object}  response.ResponseData{data=model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/create [post]
func (c *cConversationController) CreateConversation(ctx *gin.Context) {
	var params vo.CreateConversationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	conversation, err := service.ConversationService().CreateConversation(ctx, model.ConversationInput{
		Title:       params.Title,
		Description: params.Description,
		UserIDS:     params.UserIDS,
		Type:        params.Type,
		Background:  params.Background,
		Emoji:       params.Emoji,
		IsDeleted:   params.IsDeleted,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, conversation)
}

// @Summary Get conversation by ID
// @Description Get conversation details by ID
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param conversation_id path int true "Conversation ID"
// @Success 200 {object} response.ResponseData{data=model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/get-by-id/{conversation_id} [get]
func (c *cConversationController) GetConversationByID(ctx *gin.Context) {
	var params vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	conversation, err := service.ConversationService().GetConversationById(ctx, params.ConversationID)
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, conversation)
}

// @Summary Get conversations by user ID
// @Description Get all conversations for a specific user
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param user_id path int true "User ID"
// @Param limit query int false "Limit" default(50)
// @Param page query int false "Page" default(1)
// @Success 200 {object} response.ResponseData{data=[]model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversations not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/get-by-user-id/{user_id} [get]
func (c *cConversationController) GetConversationByUserID(ctx *gin.Context) {
	var params vo.UserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var query vo.PageRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	fmt.Println("Fetching conversations for user ID:", params.UserID)
	conversations, err := service.ConversationService().GetConversationByUserId(ctx, params.UserID, query)
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, conversations)
}

// @Summary Update a conversation
// @Description Update conversation details
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param conversation_id path int true "Conversation ID"
// @Param request body vo.UpdateConversationRequest true "Update conversation request"
// @Success 200 {object} response.ResponseData{data=model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/update [put]
func (c *cConversationController) UpdateConversation(ctx *gin.Context) {
	var conversationID vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&conversationID); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var params vo.UpdateConversationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if params.ConversationID != conversationID.ConversationID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "Conversation ID mismatch")
		return
	}
	conversation, err := service.ConversationService().UpdateConversation(ctx, model.UpdateConversationInput{
		ID:          params.ConversationID,
		Title:       params.Title,
		Description: params.Description,
		Type:        params.Type,
		Background:  params.Background,
		Emoji:       params.Emoji,
		IsDeleted:   params.IsDeleted,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, conversation)
}

// @Summary Add users to a conversation
// @Description Add users to an existing conversation
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param conversation_id path int true "Conversation ID"
// @Param request body vo.AddUserToConversationRequest true "Add users to conversation request"
// @Success 200 {object} response.ResponseData{data=model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/add-users/{conversation_id} [patch]
func (c *cConversationController) AddUserToConversation(ctx *gin.Context) {
	var params vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var request vo.AddUserToConversationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if params.ConversationID != request.ConversationID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "Conversation ID mismatch")
		return
	}
	conversation, err := service.ConversationService().AddUserToConversation(ctx, params.ConversationID, request.UserIDS)
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, conversation)
}

// @Summary Remove a user from a conversation
// @Description Remove a user from an existing conversation
// @Tags Conversation
// @Accept json
// @Produce json
// @Param conversation_id path int true "Conversation ID"
// @Param user_id path int true "User ID"
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Success 200 {object} response.ResponseData{data=model.ConversationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/remove-users/{conversation_id}/{user_id} [patch]
func (c *cConversationController) RemoveUserFromConversation(ctx *gin.Context) {
	var params vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var userID vo.UserIDRequest
	if err := ctx.ShouldBindUri(&userID); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	conversation, err := service.ConversationService().RemoveUserFromConversation(ctx, params.ConversationID, userID.UserID)
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, conversation)
}

// @Summary Soft delete a conversation
// @Description Soft delete a conversation by ID
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param conversation_id path int true "Conversation ID"
// @Success 200 {object} response.ResponseData{data=string} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/soft-delete/{conversation_id} [put]
func (c *cConversationController) SoftDeleteConversation(ctx *gin.Context) {
	var params vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	conversation, err := service.ConversationService().SoftDeleteConversation(ctx, model.SoftDeleteConversationInput{
		ID: params.ConversationID,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, conversation)
}

// @Summary Delete a conversation
// @Description Permanently delete a conversation by ID
// @Tags Conversation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param conversation_id path int true "Conversation ID"
// @Success 200 {object} response.ResponseData{data=string} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Conversation not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /conversation/delete/{conversation_id} [delete]
func (c *cConversationController) DeleteConversation(ctx *gin.Context) {
	var params vo.ConversationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	err := service.ConversationService().DeleteConversation(ctx, params.ConversationID)
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, "Conversation deleted successfully")
}
