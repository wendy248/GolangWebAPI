package main

import (
	"Github/GolangWebAPI/buku"
	"Github/GolangWebAPI/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//Connect to MySQL "dbbuku"
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/dbbuku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Error")
	}
	// fmt.Println("Database Connected")

	//Add table on MySQL "dbbuku"
	// db.AutoMigrate(buku.Buku{})
	// buku := buku.Buku{
	// 	ID:          1,
	// 	Title:       "Buku Matematika",
	// 	Description: "Belajar matematika itu menyenangkan",
	// 	Price:       10000,
	// 	Rating:      5,
	// 	CreatedAt:   time.Time{},
	// 	UpdatedAt:   time.Time{},
	// }

	// // input variable buku (struct) to MySQL
	// err = db.Create(&buku).Error
	// if err != nil {
	// 	fmt.Println("Error while inputing data")
	// }

	//db.First (retrieve data from MySQL) explanation : https://gorm.io/docs/query.html#Retrieving-a-single-object
	// err = db.First(&buku).Error
	// if err != nil {
	// 	fmt.Println("Error while getting data from MySQL")
	// }

	// fmt.Println("Judul :", buku.Title)
	// fmt.Println("Harga :", buku.Price)
	// fmt.Println("Deskripsi :", buku.Description)

	//Update data by ID
	var buku buku.Buku
	// err = db.Debug().Where("id = ?", 11).Find(&buku).Error 		//Check + Select data by ID
	// if err != nil {
	// 	fmt.Println("Error searching for data")
	// }
	// buku.Title = "perubahan untuk ID 11"

	db.Model(&buku).Where("id=?", 3).Update("title", "perubahan db.Model") //Update data MySQL using db.Model
	db.Save(&buku)
	if err != nil {
		fmt.Println("Error updating data")
	}

	// //Deleting data
	// db.Where("id = 2").Delete(&buku)

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.GET("/home", handler.HomeHandler)

	//localhost:8080/item?judul=filosofi-teras&harga=20000 (soal 1)
	r.GET("/item", handler.PriceCheck)

	//localhost:8080/id/2/product?name="tv"&price=1000 (soal 2)
	r.GET("/id/:number/product", handler.PriceCheckFull)

	//localhost:8080/customer/wendy/10/surabaya
	r.GET("/customer/:nama/:id/:kota", handler.UserInformation)

	//localhost:8080/stock?nama=baju&jumlah=100&tempat=surabaya
	r.GET("/stock", handler.ProductQuantity)

	//localhost:8080/QuantyStock/baju/bekas?harga=1000&jumlah=100
	r.GET("/QuantyStock/:name/:kondisi", handler.ProductCondition)

	//localhost:8080/biodata
	r.POST("/biodata", handler.BiodataHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
