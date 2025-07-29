package manager

type ManagerRouterGroup struct {
	TicketRouter
	RouteSegmentRouter
	TicketItemRouter
	StationRouter
	TicketSegmentPriceRouter
	SeatRouter
	TrainRouter
	SeatReservationRouter
}
