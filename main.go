package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ฟังก์ชันที่เป็น exported ซึ่งจะถูกใช้โดย Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// การอ่านไฟล์ config ด้วย viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")

	// การเชื่อมต่อฐานข้อมูล MySQL
	dialactor := mysql.Open(dsn)
	_, err = gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}

	// กำหนด HTTP route ที่ใช้ handler
	http.HandleFunc("/", Handler)

	// เริ่ม HTTP server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
