package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/quangdat385/holiday-ticket/order-service/cmd/swag/docs" // docs
	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/initialize"
	"github.com/quangdat385/holiday-ticket/order-service/utils/crypto"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Holiday Ticket Order Service
// @version         1.0.0
// @description     This is a sample server celler server.

// @contact.name   TEAM DATNGUYEN
// @contact.email  datnguyen03011985@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /ticket-order/api/v1/order
// @schemes   http, https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @type http
// @scheme bearer
// @bearerFormat JWT

// @security BearerAuth
func main() {
	r := initialize.Run()
	r.GET("/ping", func(c *gin.Context) {
		hashKey := crypto.GenerateHash("1", "123456")
		c.JSON(200, gin.H{
			"message": hashKey,
		})
	})
	r.GET("/ticket-order/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Starting order service...", global.Config.Server.Port)
	port := fmt.Sprintf(":%d", global.Config.Server.Port)
	r.Run(port) // listen and serve on
	// r.GET("/convert", func(c *gin.Context) {

	// 	imageFiles := []string{
	// 		"./cmd/server/img1.jpg",
	// 		"./cmd/server/img2.jpg",
	// 	}

	// 	pdf := gofpdf.New("P", "mm", "A4", "")
	// 	for _, img := range imageFiles {
	// 		pdf.AddPage()
	// 		pdf.ImageOptions(
	// 			img,
	// 			10, 10, 190, 0, // x, y, width, height (0 = auto)
	// 			false,
	// 			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	// 			0, "",
	// 		)
	// 	}
	// 	out := fmt.Sprintf("./cmd/server/%d", time.Now().Unix())
	// 	err := pdf.OutputFileAndClose(out + ".pdf")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"message": "success",
	// 	})
	// })
	// r.GET("/convert-pro", func(c *gin.Context) {
	// 	imageFiles := []string{
	// 		"./cmd/server/img1.jpg",
	// 		"./cmd/server/img2.jpg",
	// 	}

	// 	pdf := gofpdf.New("P", "mm", "A4", "")
	// 	for i, imgPath := range imageFiles {
	// 		imgBytes, err := os.ReadFile(imgPath)
	// 		if err != nil {
	// 			c.JSON(500, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		imgName := fmt.Sprintf("img%d", i)
	// 		pdf.AddPage()
	// 		pdf.RegisterImageOptionsReader(
	// 			imgName,
	// 			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	// 			bytes.NewReader(imgBytes),
	// 		)
	// 		pdf.ImageOptions(
	// 			imgName,
	// 			10, 10, 190, 0,
	// 			false,
	// 			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	// 			0, "",
	// 		)
	// 	}
	// 	out := fmt.Sprintf("./cmd/server/%d.pdf", time.Now().Unix())
	// 	err := pdf.OutputFileAndClose(out)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"message": "ok",
	// 	})
	// })
	// r.POST("/convert-upload", func(c *gin.Context) {
	// 	form, err := c.MultipartForm()
	// 	if err != nil {
	// 		c.JSON(400, gin.H{"error": "Invalid form"})
	// 		return
	// 	}
	// 	files := form.File["images"] // "images" là tên trường input

	// 	pdf := gofpdf.New("P", "mm", "A4", "")
	// 	for i, fileHeader := range files {
	// 		file, err := fileHeader.Open()
	// 		if err != nil {
	// 			c.JSON(500, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		defer file.Close()
	// 		imgBytes, err := io.ReadAll(file)
	// 		if err != nil {
	// 			c.JSON(500, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		imgName := fmt.Sprintf("img%d", i)
	// 		pdf.AddPage()
	// 		pdf.RegisterImageOptionsReader(
	// 			imgName,
	// 			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	// 			bytes.NewReader(imgBytes),
	// 		)
	// 		pdf.ImageOptions(
	// 			imgName,
	// 			10, 10, 190, 0,
	// 			false,
	// 			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	// 			0, "",
	// 		)
	// 	}
	// 	out := fmt.Sprintf("./cmd/server/%d.pdf", time.Now().Unix())
	// 	err = pdf.OutputFileAndClose(out)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{"message": "ok"})
	// })
	// port := fmt.Sprintf(":%d", global.Config.Server.Port)
	// r.Run(port)
}

//  website: [
//       {
//         url: 'https://api.datnguyendev.com/ticket-user/api/v1/docs',
//         service: 'User service'
//       },
//       {
//         url: 'https://api.datnguyendev.com/ticket-service/api/v1/swagger/index.html',
//         service: 'Ticket service  '
//       },
//       {
//         url: 'https://api.datnguyendev.com/ticket-order/api/v1/swagger/index.html',
//         service: 'Order service'
//       },
//       {
//         url: 'https://api.datnguyendev.com/ticket-payment/api/v1/swagger/index.html',
//         service: 'Payment service'
//       },
//       {
//         url: 'https://api.datnguyendev.com/ticket-communication/api/v1/swagger/index.html',
//         service: 'Communication service'
//       }
//     ],
