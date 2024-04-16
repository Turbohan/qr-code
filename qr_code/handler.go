package qr_code

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"qrcode/model"
	"strconv"

	"github.com/skip2/go-qrcode"
)

func HandleScan(w http.ResponseWriter, r *http.Request) {
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
	employ, err := model.GetEmployeeInfo(employeeID)
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

func ScanHandler(w http.ResponseWriter, r *http.Request) {
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
	employee, err := model.GetEmployeeInfo(employeeID)
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
			<h2 style="text-align: center;">中国水利工程协会证书验证结果</h2>
			<div style="text-align: center;">
				<img src="{{.Pic}}" height="150px" width="100px" alt="Avatar">
			</div>
			<p>证书名称: {{.CertiName}}</p>
			<p>姓&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;名: {{.Name}}</p>
			<p>身份证号: {{.IdNum}}</p>
			<p>证书编号: {{.CertiNum}}</p>
			<p>岗位名称: {{.Title}}</p>
			<p>当前状态: {{.Status}}</p>
			<p>工作单位: {{.Work}}</p>
			<p>有效期至: %s</p>
			<p>更新日期: %s</p>
			<script>
				function updateDate() {
					var currentDate = new Date();
					var formattedDate = currentDate.getFullYear() + "-" + (currentDate.getMonth() + 1) + "-" + currentDate.getDate() + " " + currentDate.toLocaleTimeString();
					document.getElementById("currentDate").innerText = formattedDate;
				}
				window.onload = function() {
					updateDate();
					setInterval(updateDate, 1000); // 每秒更新一次日期
				};
			</script>
			<br>
			<p>验证时间: <span id="currentDate"></span></p>
			</br>
		</body>
		</html>
	`
	html = fmt.Sprintf(html, employee.ExpireTime.Format("2006年1月2日"), employee.UpdateAt.Format("2006年1月2日"))

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

// QrCodeHandler 处理二维码图片的请求
func QrCodeHandler(w http.ResponseWriter, r *http.Request) {
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
