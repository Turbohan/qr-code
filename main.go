package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"qrcode/qr_code"
)

var db *sql.DB

func main() {
	var err error
	// 连接到MySQL数据库
	//db, err = sql.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/qr_code?parseTime=true")
	db, err = sql.Open("mysql", "apptable123:Q)#2Mp9zu>@tcp(10.99.128.194:3306)/chatflow_admin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("qr_code db start")

	// 设置HTTP路由
	http.HandleFunc("/scan", qr_code.HandleScan)
	http.HandleFunc("/scan_handler", qr_code.ScanHandler)
	http.HandleFunc("/generate", qr_code.QrCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("qr_code start")
}
