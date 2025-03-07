package main

import (
	"fmt"
	"go-grom/controller"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// รูปแบบการใส่
	// dsn := "user:pwd@tcp(ip:port)/dbname?collation=utf8mb4_unicode_ci&parseTime=true"

	// ใช้ไฟล์ config.yaml เก็บค่า config ต่างๆ เช่นค่า dsn
	// dsn := "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark?collation=utf8mb4_unicode_ci&parseTime=true"
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")

	dialactor := mysql.Open(dsn)

	// _ แปลว่า มีตัวแปลอยู่นะแต่ยังไม่ให้เอาไปเช็ค หรือ ใช้ เพื่อไม่ให้ err
	// db, err := gorm.Open(dialactor)
	_, err = gorm.Open(dialactor)
	// db,err := gorm.Open(dialactor)
	if err != nil {
		panic(err)

	}

	// StartServer web API

	controller.StartServer()
	// print("connnn ssssss")

	// countries := []model.Country{}
	// db.Find(&countries)

	// // fmt.Printf("%v", countries)

	// // gpt
	// landmarks := []model.Landmark{}
	// // db.Select("name").Find(&landmarks)
	// db.Find(&landmarks)
	// // for _, landmark := range landmarks {
	// // 	fmt.Println(landmark.Name)
	// // }

	// // fmt.Printf("%v", landmarks)

	// // Preload ใช้เรียก sql แบบเล็กๆ หลายๆ ก้อน ส่วน Joins เรียก ข้อมูล ใหญ่ๆมาครั้งเดียวเลย
	// // Joins ใช้ query น้อยกว่า  ซับซ้อนกว่า มักใช้การซ้อนข้อมูลเยอะๆ  โดยทั่วๆไปดูเหมือนจะเร็วกว่า
	// // db.Preload("CountryData").Find(&landmarks)
	// // db.Joins("CountryData").Find(&landmarks)

	// db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Find(&countries)
	// // fmt.Printf("%v", countries)

	// country := model.Country{}
	// db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").First(&country)
	// // fmt.Printf("%v", country)

	// db.Limit(3).Order("Name").Find(&countries)
	// db.Limit(3).Order("Name desc").Find(&countries)
	// fmt.Printf("%v", countries)

	// // Distinct  แปลว่าเอามาแค่ที่ ไม่ซ้ำกัน
	// db.Distinct("Idx").Find(&countries)
	// fmt.Println(countries)

	// type unique struct {
	// 	Idx int
	// }

	// uniques := []unique{}
	// db.Model(model.Country{}).Distinct("Idx").Find(&uniques)
	// fmt.Println(uniques)

	// country := model.Country{Name: "USA"}
	// result := db.Create(&country)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// if result.RowsAffected > 0 {
	// 	println("sssssssssssss")
	// }

	// countries := []model.Country{}
	// db.Find(&countries)
	// fmt.Printf("%v", countries)

}
