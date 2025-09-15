package impl

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	pb "github.com/quangdat385/holiday-ticket/ticket-service/pkg/api/proto"
)

type server struct {
	pb.UnimplementedTicketServiceServer
	r *database.Queries
}

func (s *server) GetTicket(ctx context.Context, req *pb.GetTicketRequest) (*pb.PreGoTicket99999, error) {
	fmt.Println("GetTicket called with request:", req)
	ticketDB, err := s.r.GetTicketById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// Simulate a ticket retrieval
	ticket := &pb.PreGoTicket99999{
		Id:          ticketDB.ID,
		Name:        ticketDB.Name,
		StartTime:   timestamppb.New(ticketDB.StartTime),
		EndTime:     timestamppb.New(ticketDB.EndTime),
		Description: ticketDB.Description.String,
	}
	return ticket, nil
}
func (s *server) GetTicketItem(ctx context.Context, req *pb.GetTicketItemRequest) (*pb.PreGoTicketItem99999, error) {
	fmt.Println("GetTicketItem called with request:", req)
	ticketItemDB, err := s.r.GetTicketItemById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// Simulate a ticket item retrieval
	ticketItem := &pb.PreGoTicketItem99999{
		Id:             ticketItemDB.ID,
		Name:           ticketItemDB.Name,
		Description:    ticketItemDB.Description.String,
		StockInitial:   ticketItemDB.StockInitial,
		StockAvailable: ticketItemDB.StockAvailable,
		PriceOriginal:  parseStringToFloat64(ticketItemDB.PriceOriginal),
		SaleStartTime:  timestamppb.New(ticketItemDB.SaleStartTime),
		SaleEndTime:    timestamppb.New(ticketItemDB.SaleEndTime),
		Status:         ticketItemDB.Status,
		ActivityId:     int32(ticketItemDB.ActivityID),
	}
	return ticketItem, nil
}
func parseStringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func InitServer(address string) error {
	fmt.Println("Starting gRPC server on", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	fmt.Println("Listening on", address)
	s := grpc.NewServer()
	r := database.New(global.Mdb)
	pb.RegisterTicketServiceServer(s, &server{
		r: r,
	})

	log.Printf("Server is running on %s", address)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	fmt.Println("Server is running on", address)
	return nil
}
