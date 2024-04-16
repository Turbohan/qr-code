package model

import "time"

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

func GetEmployeeInfo(employeeID int) (*Employee, error) {
	// 查询员工信息
	var employee Employee
	//// err := db.QueryRow("SELECT * FROM employ WHERE id = ?", employeeID).Scan(&employee)
	//err := db.QueryRow("SELECT id, name, pic, certi_name, id_num, certi_num, title, status, work, expire_time, update_at, verify_time FROM employ WHERE id = ?", employeeID).
	//	Scan(&employee.ID, &employee.Name, &employee.Pic, &employee.CertiName, &employee.IdNum, &employee.CertiNum, &employee.Title, &employee.Status, &employee.Work, &employee.ExpireTime, &employee.UpdateAt, &employee.VerifyTime)
	//
	//if err != nil {
	//	return nil, err
	//}

	expireTime1, _ := time.Parse("2006-01-02", "2026-01-11")
	updateAt1, _ := time.Parse("2006-01-02 15:04:05", "2023-01-11")
	expireTime2, _ := time.Parse("2006-01-02", "2026-01-09")
	updateAt2, _ := time.Parse("2006-01-02 15:04:05", "2023-01-9")
	expireTime3, _ := time.Parse("2006-01-02", "2027-01-10")
	updateAt3, _ := time.Parse("2006-01-02 15:04:05", "2023-01-10")
	expireTime4, _ := time.Parse("2006-01-02", "2026-05-18")
	updateAt4, _ := time.Parse("2006-01-02 15:04:05", "2023-05-18")
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
			Work:       "黑龙江省华海建筑水利工程有限公司",
			ExpireTime: expireTime1,
			UpdateAt:   updateAt1,
			VerifyTime: time.Now(),
		},
		{
			ID:         2,
			Name:       "杨利滨",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/yanglibing.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "23260*********4910",
			CertiNum:   "SGL20192300696",
			Title:      "施工员 2019年09月27日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司",
			ExpireTime: expireTime2,
			UpdateAt:   updateAt2,
			VerifyTime: time.Now(),
		},
		{
			ID:         3,
			Name:       "李海峰",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/lihaifeng.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "23022*********2455",
			CertiNum:   "SGL20192300877",
			Title:      "质检员 2020年07月28日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司",
			ExpireTime: expireTime3,
			UpdateAt:   updateAt3,
			VerifyTime: time.Now(),
		},
		{
			ID:         4,
			Name:       "师勇贺",
			Pic:        "https://innowoa.ks3-cn-beijing.ksyuncs.com/test/code/shiyonghe.jpg",
			CertiName:  "水利水电工程施工现场管理人员培训合格证书",
			IdNum:      "23062*********0653",
			CertiNum:   "SGL20232300458",
			Title:      "施工员 2023年05月18日",
			Status:     "正常",
			Work:       "黑龙江省华海建筑水利工程有限公司",
			ExpireTime: expireTime4,
			UpdateAt:   updateAt4,
			VerifyTime: time.Now(),
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
