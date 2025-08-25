package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
)

var NotificationControllerRouter = new(cNotificationController)

type cNotificationController struct{}

// @Summary Get notification by ID
// @Description Get notification by ID
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param notification_id path int true "Notification ID"
// @Success 200 {object} response.ResponseData{data=model.NotificationOutput} "Notification retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notification not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/get-by-id/{notification_id} [get]
func (c *cNotificationController) GetNotificationByID(ctx *gin.Context) {
	var params vo.NotificationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.NotificationService().GetNotificationById(ctx, int64(params.NotificationID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get notifications by user ID (to)
// @Description Get notifications by user ID (to)
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} response.ResponseData{data=[]model.NotificationOutput} "Notifications retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notifications not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/get-by-user-id-to/{user_id} [get]
func (c *cNotificationController) GetNotificationsByUserIDTo(ctx *gin.Context) {
	var params vo.UserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.NotificationService().GetNotificationsByUserIDTo(ctx, int64(params.UserID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get notifications by user ID (from)
// @Description Get notifications by user ID (from)
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} response.ResponseData{data=model.NotificationOutput} "Notification retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notification not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/get-by-user-id-from/{user_id} [get]
func (c *cNotificationController) GetNotificationsByUserIDFrom(ctx *gin.Context) {
	var params vo.UserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.NotificationService().GetNotificationsByUserIDFrom(ctx, int64(params.UserID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Get notifications with user ID (to) as null
// @Description Get notifications with user ID (to) as null
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Success 200 {object} response.ResponseData{data=[]model.NotificationOutput} "Notifications retrieved successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notifications not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/get-by-user-id-to-null [get]
func (c *cNotificationController) GetNotificationsByUserIDToNull(ctx *gin.Context) {
	out, err := service.NotificationService().GetNotificationsFromUserIDToIsNull(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Create a new notification
// @Description Create a new notification
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param input body vo.CreateNotificationRequest true "Notification Input"
// @Success 201 {object} response.ResponseData{data=model.NotificationOutput} "Notification created successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/create [post]
func (c *cNotificationController) CreateNotification(ctx *gin.Context) {
	var input vo.CreateNotificationRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.NotificationService().CreateNotification(ctx, model.NotificationInput{
		From:    input.From,
		To:      input.To,
		Content: input.Content,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, out)
}

// @Summary Update Notification By ID
// @Description Update Notification By ID
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param input body vo.UpdateNotificationRequest true "Notification Input"
// @Success 200 {object} response.ResponseData{data=model.NotificationOutput} "Notification updated successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notification not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/update-status [put]
func (c *cNotificationController) UpdateNotification(ctx *gin.Context) {
	var input vo.UpdateNotificationRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.NotificationService().UpdateNotification(ctx, model.UpdateNotificationInput{
		UserID:         input.UserID,
		NotificationID: input.NotificationID,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, out)
}

// @Summary Delete notification by ID
// @Description Delete notification by ID
// @Tags Notification
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Notification ID"
// @Success 204 {object} response.ResponseData "Notification deleted successfully"
// @Failure 400 {object} response.ErrorResponseData "Bad Request"
// @Failure 404 {object} response.ErrorResponseData "Notification not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /notification/delete/{id} [delete]
func (c *cNotificationController) DeleteNotification(ctx *gin.Context) {
	var params vo.NotificationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if err := service.NotificationService().DeleteNotificationById(ctx, int64(params.NotificationID)); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, nil)
}
