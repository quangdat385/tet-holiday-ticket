package initialize

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
	routers "github.com/quangdat385/holiday-ticket/ticket-service/internal/router"
	pb "github.com/quangdat385/holiday-ticket/ticket-service/pkg/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())
	// r.Use(middleware.CSRFMiddleware())
	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager

	MainGroup := r.Group("/ticket-service/api/v1/ticket")
	{
		MainGroup.GET("/checkStatus")
		MainGroup.GET("/hello", func(c *gin.Context) {
			conn, err := grpc.NewClient("localhost:50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			defer conn.Close()
			fmt.Println("Connected to gRPC server", conn.Target())

			client := pb.NewUserServiceGrpcClient(conn)
			resp, err := client.FindOne(context.Background(), &pb.UserById{UserId: 1})
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{
				"code":    200,
				"message": "ok",
				"data":    resp,
			})
		})
		{
			userRouter.TicketRouter.InitTicketRouter(MainGroup)
			userRouter.TicketItemRouter.InitTicketItemRouter(MainGroup)
			userRouter.RouteSegmentRouter.InitRouteSegmentRouter(MainGroup)
			userRouter.StationRouter.InitStationRouter(MainGroup)
			userRouter.TicketSegmentPriceRouter.InitTicketSegmentPriceRouter(MainGroup)
			userRouter.SeatReservationRouter.InitSeatReservationRouter(MainGroup)
			userRouter.TrainRouter.InitTrainRouter(MainGroup)
			userRouter.SeatRouter.InitSeatRouter(MainGroup)
		}
		{
			managerRouter.TicketRouter.InitTicketRouter(MainGroup)
			managerRouter.RouteSegmentRouter.InitRouteSegmentRouter(MainGroup)
			managerRouter.TicketItemRouter.InitTicketItemRouter(MainGroup)
			managerRouter.StationRouter.InitStationRouter(MainGroup)
			managerRouter.TicketSegmentPriceRouter.InitTicketSegmentPriceRouter(MainGroup)
			managerRouter.SeatReservationRouter.InitSeatReservationRouter(MainGroup)
			managerRouter.TrainRouter.InitTrainRouter(MainGroup)
			managerRouter.SeatRouter.InitSeatRouter(MainGroup)
		}
	}
	return r
}
