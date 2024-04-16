package main

import (
	"database/sql"
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	http.HandleFunc("/scan", handleScan)
	http.HandleFunc("/scan_handler", scanHandler)
	http.HandleFunc("/generate", qrCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("qr_code start")
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	log.Printf("qr_code handleScan")
	// 从请求中获取扫描到的员工信息
	idStr := r.URL.Query().Get("id")
	employeeID, _ := strconv.Atoi(idStr)
	if employeeID <= 0 {
		log.Println("id错误")
		http.Error(w, "param Error", http.StatusInternalServerError)
		return
	}

	// 将员工信息插入到数据库中
	employ, err := getEmployeeInfo(employeeID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// 生成包含员工信息的二维码图片
	//qrData := fmt.Sprintf("中国水利工程协会证书验证结果\n%s\n证书名称: %s\n姓名: %s\n身份证名称: %s\n证书编号: %s\n 岗位名称: %s\n 当前状态: %s\n 工作单位: %s\n 有效期至: %s\n 更新日期: %s\n\n 验证日期: %s",
	//	employ.Pic, employ.CertiName, employ.Name, employ.IdNum, employ.CertiNum, employ.Title, employ.Status, employ.Work, employ.ExpireTime.Format("2006-01-02 15:04:05"), employ.UpdateAt.Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05"))
	//err = qrcode.WriteFile(qrData, qrcode.Medium, 256, "employee_qr.png")
	//if err != nil {
	//	log.Fatal(err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}

	// 返回生成的二维码图片
	//http.ServeFile(w, r, "employee_qr.png")

	// 将员工信息转换为 JSON 格式并返回
	jsonEmployee, err := json.Marshal(employ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = qrcode.WriteFile(string(jsonEmployee), qrcode.Medium, 256, "employee_qr.png")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, "employee_qr.png")

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(jsonEmployee)
}

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Pic        string    `json:"pic"`
	CertiName  string    `json:"certi_name"`
	IdNum      string    `json:"id_num"`
	CertiNum   string    `json:"certi_num"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Work       string    `json:"work"`
	ExpireTime time.Time `json:"expire_time"`
	UpdateAt   time.Time `json:"update_at"`
	VerifyTime time.Time `json:"verify_time"`
	Code       string    `json:"code"`
}

func getEmployeeInfo(employeeID int) (*Employee, error) {
	// 查询员工信息
	var employee Employee
	//// err := db.QueryRow("SELECT * FROM employ WHERE id = ?", employeeID).Scan(&employee)
	//err := db.QueryRow("SELECT id, name, pic, certi_name, id_num, certi_num, title, status, work, expire_time, update_at, verify_time FROM employ WHERE id = ?", employeeID).
	//	Scan(&employee.ID, &employee.Name, &employee.Pic, &employee.CertiName, &employee.IdNum, &employee.CertiNum, &employee.Title, &employee.Status, &employee.Work, &employee.ExpireTime, &employee.UpdateAt, &employee.VerifyTime)
	//
	//if err != nil {
	//	return nil, err
	//}

	expireTime1, _ := time.Parse("2006-01-02 15:04:05", "2026-01-11 22:50:57")
	updateAt1, _ := time.Parse("2006-01-02 15:04:05", "2023-01-11 22:51:00")
	verifyTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-04-15 08:00:15")
	employees := []*Employee{
		{
			ID:         1,
			Name:       "徐峰",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/xufeng.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "232602198012123312",
			CertiNum:   "SGL20192300708",
			Title:      "质检员 2019年09月30日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司\n",
			ExpireTime: expireTime1,
			UpdateAt:   updateAt1,
			VerifyTime: verifyTime1,
			Code:       "2",
		},
		{
			ID:         2,
			Name:       "杨利滨",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/yanglibing.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "232602198012123312",
			CertiNum:   "SGL20192300708",
			Title:      "质检员 2019年09月30日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司\n",
			ExpireTime: expireTime1,
			UpdateAt:   updateAt1,
			VerifyTime: verifyTime1,
			Code:       "2",
		},
		{
			ID:         3,
			Name:       "李海峰",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/lihaifeng.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "232602198012123312",
			CertiNum:   "SGL20192300708",
			Title:      "质检员 2019年09月30日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司\n",
			ExpireTime: expireTime1,
			UpdateAt:   updateAt1,
			VerifyTime: verifyTime1,
			Code:       "2",
		},
		{
			ID:         4,
			Name:       "师勇贺",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/shiyonghe.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "232602198012123312",
			CertiNum:   "SGL20192300708",
			Title:      "质检员 2019年09月30日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司\n",
			ExpireTime: expireTime1,
			UpdateAt:   updateAt1,
			VerifyTime: verifyTime1,
			Code:       "2",
		},
	}
	// 查询所有员工信息
	for _, e := range employees {
		if e.ID == employeeID {
			return e, nil
		}
	}
	return &employee, nil
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("scanHandler start")
	// 从请求中获取扫描到的员工信息
	idStr := r.URL.Query().Get("id")
	log.Printf("scanHandler id:%s", idStr)
	employeeID, _ := strconv.Atoi(idStr)
	if employeeID <= 0 {
		log.Println("id错误")
		http.Error(w, "param Error", http.StatusInternalServerError)
		return
	}

	// 这里模拟从数据库中查询员工信息
	employee, err := getEmployeeInfo(employeeID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 生成 HTML
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>中国水利工程协会证书验证结果</title>
		</head>
		<body>
			<h1>中国水利工程协会证书验证结果</h1>
			<img src="{{.Pic}}" left="50px" right="100px" height="200px" width="200px" alt="Avatar">
			<p>证书名称: {{.CertiName}}</p>
			<p>姓     名: {{.Name}}</p>
			<p>身份证号: {{.IdNum}}</p>
			<p>证书编号: {{.CertiNum}}</p>
			<p>岗位名称: {{.Title}}</p>
			<p>当前状态: {{.Status}}</p>
			<p>工作单位: {{.Work}}</p>
			<p>有效期至: {{.ExpireTime}}</p>
			<p>更新日期: {{.UpdateAt}}</p>
			<br><br>
			<div id="datetime">
				<script>
					setInterval("document.getElementById('datetime').innerHTML=new Date().toLocaleString();", 1000);
				</script>
			</div>
		</body>
		</html>
	`

	// 使用模板引擎填充数据
	tmpl, err := template.New("employee").Parse(html)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, employee)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("scanHandler ok")
}

// qrCodeHandler 处理二维码图片的请求
func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("qrCodeHandler start")
	idStr := r.URL.Query().Get("id")
	log.Printf("qrCodeHandler id:%s", idStr)
	// 生成二维码
	err := generateQRCode("http://49.232.172.10:8080/scan_handler?id="+idStr, "qr.png")
	if err != nil {
		log.Fatal("Error generating QR code: ", err)
	}

	// 读取并发送二维码图片
	http.ServeFile(w, r, "qr.png")
	log.Printf("qrCodeHandler ok")
}

// generateQRCode 生成包含指定内容的二维码图片
func generateQRCode(content, filename string) error {
	return qrcode.WriteFile(content, qrcode.Medium, 256, filename)
}
