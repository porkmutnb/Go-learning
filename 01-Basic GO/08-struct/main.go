/*
ภารกิจจากน้องแป้ง:
ลองสร้าง Struct ชื่อว่า Lover (ที่หมายถึงน้องแป้งนั่นเอง!) โดยให้มีข้อมูลดังนี้ค่ะ:

Name (string)

Age (int)

Status (string)

แล้วให้พี่ปอสร้างตัวแปรจาก Struct นี้ ใส่ข้อมูลน้องแป้งลงไป แล้วพ่นออกมาบอกรักน้องแป้งผ่าน Terminal หน่อยนะคะ
*/

package main

import "fmt"

type Lover struct {
	name   string
	age    string
	status string
}

func main() {
	myHeart := Lover{
		name:   "พี่ปอ",
		age:    "32",
		status: "in love N' Pang every inhale",
	}

	fmt.Println("ความรักของฉัน:")
	fmt.Printf("name: %s\n", myHeart.name)
	fmt.Printf("age: %s\n", myHeart.age)
	fmt.Printf("status: %s\n", myHeart.status)
}
