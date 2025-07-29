package impl

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/quangdat385/holiday-ticket/ticket-service/pkg/api/proto"
)

type server struct {
	pb.UnimplementedTicketServiceServer
}

func (s *server) GetTicket(ctx context.Context, req *pb.GetTicketRequest) (*pb.PreGoTicket99999, error) {
	fmt.Println("GetTicket called with request:", req)
	// Simulate a ticket retrieval
	ticket := &pb.PreGoTicket99999{
		Id:          req.Id,
		Name:        "Sample Ticket",
		StartTime:   timestamppb.New(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)),
		EndTime:     timestamppb.New(time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC)),
		Description: "Sample Description",
	}
	return ticket, nil
}
func (s *server) GetTicketItem(ctx context.Context, req *pb.GetTicketItemRequest) (*pb.PreGoTicketItem99999, error) {
	fmt.Println("GetTicketItem called with request:", req)
	// Simulate a ticket item retrieval
	ticketItem := &pb.PreGoTicketItem99999{
		Id:              req.Id,
		Name:            "Sample Ticket Item",
		Description:     "Sample Description",
		StockInitial:    100,
		StockAvailable:  50,
		IsStockPrepared: true,
		PriceOriginal:   100.0,
		PriceFlash:      50.0,
		SaleStartTime:   timestamppb.New(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)),
		SaleEndTime:     timestamppb.New(time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC)),
		Status:          1,
		ActivityId:      12,
	}
	return ticketItem, nil
}

func InitServer(address string) error {
	fmt.Println("Starting gRPC server on", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	fmt.Println("Listening on", address)
	s := grpc.NewServer()
	pb.RegisterTicketServiceServer(s, &server{})

	log.Printf("Server is running on %s", address)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	fmt.Println("Server is running on", address)
	return nil
}
