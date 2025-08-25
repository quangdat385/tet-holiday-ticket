package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
)

var InformationRouter = new(cInformationController)

type cInformationController struct {
}

// @Summary Get information by user ID
// @Description Retrieve information by user ID
// @Tags Information
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Success 200 {object} response.ResponseData{data=model.InformationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Information not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /information/get-by-user-id/{user_id} [get]
func (c *cInformationController) GetInformationByUserID(ctx *gin.Context) {
	var params vo.InformationUserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	fmt.Println("GetInformationByUserID called with params:", params.UserID)
	out, err := service.InformationService().GetInformationByUserID(ctx.Request.Context(), params.UserID)
	if err != nil {
		// Handle error (e.g., log it, return an error response)
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary Update information by user ID
// @Description Update information by user ID
// @Tags Information
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param request body vo.CreateInformationRequest true "Request body"
// @Success 200 {object} response.ResponseData{data=model.InformationOutput} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Information not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /information/update-by-user-id/{user_id} [put]
func (c *cInformationController) UpdateInformationByUserID(ctx *gin.Context) {
	var params vo.InformationUserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var input vo.CreateInformationRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.InformationService().UpdateInformationByUserID(ctx.Request.Context(), model.InfomationInput{
		UserID: params.UserID,
		Status: input.Status,
		Value:  input.Value,
		Type:   input.Type,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, out)
}

// @Summary Create information by user ID
// @Description Create information by user ID
// @Tags Information
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param request body vo.CreateInformationRequest true "Request body"
// @Success 200 {object} response.ResponseData{data=model.InformationOutput} "Success
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /information/create [post]
func (c *cInformationController) InsertInformationByUserID(ctx *gin.Context) {
	var input vo.CreateInformationRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.InformationService().InsertInformationByUserID(ctx.Request.Context(), model.InfomationInput{
		UserID: input.UserID,
		Status: input.Status,
		Value:  input.Value,
		Type:   input.Type,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, out)
}

// @Summary Delete information by ID
// @Description Delete information by ID
// @Tags Information
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param id path int true "Information ID"
// @Success 200 {object} response.ResponseData{data=string} "Success"
// @Failure 400 {object} response.ErrorResponseData "Invalid parameters"
// @Failure 404 {object} response.ErrorResponseData "Information not found"
// @Failure 500 {object} response.ErrorResponseData "Internal server error"
// @Router /information/delete/{id} [delete]
func (c *cInformationController) DeleteInformationByID(ctx *gin.Context) {
	var params vo.InformationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	out, err := service.InformationService().DeleteInformationByID(ctx.Request.Context(), params.InformationID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, out)
}
