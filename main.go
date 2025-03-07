package main

import (
	"fmt"
	"go-grom/controller"
	"net/http"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ฟังก์ชันที่ใช้ในการจัดการคำขอ HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // ส่งข้อความตอบกลับไปยังผู้ใช้
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")

	dialactor := mysql.Open(dsn)

	_, err = gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}

	// กำหนด HTTP route ที่เรียกใช้ handler
	http.HandleFunc("/", handler) // เมื่อเข้าถึง URL "/" จะเรียก handler

	// เริ่ม HTTP server
	go controller.StartServer() // หากคุณต้องการรัน server ของ controller ใน background

	// เริ่ม HTTP server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
