package impl

import (
	"context"
	"strconv"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

type sTicketSegmentPrice struct {
	r *database.Queries
}

func NewTicketSegmentPriceImpl(r *database.Queries) *sTicketSegmentPrice {
	return &sTicketSegmentPrice{
		r: r,
	}
}
func (s *sTicketSegmentPrice) GetTicketSegmentPriceById(context context.Context, id int64) (out model.TicketSegmentPriceOutPut, err error) {
	ticketSegmentPrice, err := s.r.GetTicketSegmentPriceById(context, id)
	if err != nil {
		return out, err
	}
	if ticketSegmentPrice == (database.GetTicketSegmentPriceByIdRow{}) {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	out = mapper.ToTicketSegmentPriceDTO(ticketSegmentPrice)
	return out, nil
}
func (s *sTicketSegmentPrice) GetTicketSegmentPriceByRouteSegmentId(context context.Context, routeSegmentId int64) (out model.TicketSegmentPriceOutPut, err error) {
	ticketSegmentPrice, err := s.r.GetTicketSegmentPricesByRouteSegmentId(context, routeSegmentId)
	if err != nil {
		return out, err
	}
	if ticketSegmentPrice == (database.GetTicketSegmentPricesByRouteSegmentIdRow{}) {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	out = mapper.ToTicketSegmentPriceDTO(database.GetTicketSegmentPriceByIdRow(ticketSegmentPrice))
	return out, nil
}
func (s *sTicketSegmentPrice) GetAllTicketSegmentPriceFromSegmentIDToToSegmentID(context context.Context, in model.TicketSegmentPriceListInPut) (out []model.TicketSegmentPriceOutPut, err error) {
	ticketSegmentPrice, err := s.r.GetTicketSegmentPricesFromSegmentIDToToSegmentID(context, database.GetTicketSegmentPricesFromSegmentIDToToSegmentIDParams{
		RouteSegmentID:   in.FromSegmentID,
		RouteSegmentID_2: in.ToSegmentID,
	})
	if err != nil {
		return out, err
	}
	if len(ticketSegmentPrice) == 0 {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	for _, ticketSegmentPriceRow := range ticketSegmentPrice {
		out = append(out, mapper.ToTicketSegmentPriceDTO(database.GetTicketSegmentPriceByIdRow(ticketSegmentPriceRow)))
	}
	return out, nil
}
func (s *sTicketSegmentPrice) GetAllTicketSegmentPrice(context context.Context, in model.GetAllTicketSegmentPriceInPut) (out []model.TicketSegmentPriceOutPut, err error) {
	ticketSegmentPrice, err := s.r.GetAllTicketSegmentPrices(context, database.GetAllTicketSegmentPricesParams{
		Limit:  int32(in.Limit),
		Offset: int32((in.Page - 1) * in.Limit),
	})
	if err != nil {
		return out, err
	}
	if len(ticketSegmentPrice) == 0 {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	for _, ticketSegmentPriceRow := range ticketSegmentPrice {
		out = append(out, mapper.ToTicketSegmentPriceDTO(database.GetTicketSegmentPriceByIdRow(ticketSegmentPriceRow)))
	}
	return out, nil
}
func (s *sTicketSegmentPrice) CreateTicketSegmentPrice(context context.Context, ticketSegmentPrice model.TicketSegmentPriceCreateInPut) (out model.TicketSegmentPriceOutPut, err error) {
	result, err := s.r.InsertTicketSegmentPrice(context, database.InsertTicketSegmentPriceParams{
		TicketItemID:   ticketSegmentPrice.TicketItemID,
		RouteSegmentID: ticketSegmentPrice.RouteSegmentID,
		Price:          strconv.FormatFloat(ticketSegmentPrice.Price, 'f', -1, 64),
	})
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	ticketSegmentPriceRow, err := s.r.GetTicketSegmentPriceById(context, id)
	if err != nil {
		return out, err
	}
	if ticketSegmentPriceRow == (database.GetTicketSegmentPriceByIdRow{}) {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	out = mapper.ToTicketSegmentPriceDTO(ticketSegmentPriceRow)
	return out, nil
}
func (s *sTicketSegmentPrice) UpdateTicketSegmentPrice(context context.Context, in model.TicketSegmentPriceUpdateInPut) (out model.TicketSegmentPriceOutPut, err error) {
	ticketSegmentPriceRow, err := s.r.GetTicketSegmentPriceById(context, in.ID)
	if err != nil {
		return out, err
	}
	if ticketSegmentPriceRow == (database.GetTicketSegmentPriceByIdRow{}) {
		return out, response.ErrTicketSegmentPriceNotFoundErr
	}
	if in.TicketItemID != 0 {
		ticketSegmentPriceRow.TicketItemID = in.TicketItemID
	}
	if in.RouteSegmentID != 0 {
		ticketSegmentPriceRow.RouteSegmentID = in.RouteSegmentID
	}
	if in.Price != 0 {
		ticketSegmentPriceRow.Price = strconv.FormatFloat(in.Price, 'f', -1, 64)
	}
	result, err := s.r.UpdateTicketSegmentPrice(context, database.UpdateTicketSegmentPriceParams{
		ID:             in.ID,
		TicketItemID:   ticketSegmentPriceRow.TicketItemID,
		RouteSegmentID: ticketSegmentPriceRow.RouteSegmentID,
		Price:          ticketSegmentPriceRow.Price,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	out = mapper.ToTicketSegmentPriceDTO(ticketSegmentPriceRow)
	return out, nil
}
func (s *sTicketSegmentPrice) DeleteTicketSegmentPrice(context context.Context, id int64) (err error) {
	ticketSegmentPriceRow, err := s.r.GetTicketSegmentPriceById(context, id)
	if err != nil {
		return err
	}
	if ticketSegmentPriceRow == (database.GetTicketSegmentPriceByIdRow{}) {
		return response.ErrTicketSegmentPriceNotFoundErr
	}
	_, err = s.r.DeleteTicketSegmentPrice(context, id)
	if err != nil {
		return err
	}
	return nil
}
