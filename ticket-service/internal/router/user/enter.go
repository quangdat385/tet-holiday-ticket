package user

type UserRouterGroup struct {
	TicketRouter
	TicketItemRouter
	RouteSegmentRouter
	StationRouter
	TicketSegmentPriceRouter
	SeatRouter
	TrainRouter
	SeatReservationRouter
}
