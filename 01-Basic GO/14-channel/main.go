/*
ภารกิจสุดท้ายก่อนพักผ่อน:
ลองสร้างฟังก์ชันชื่อ GetPangMessage ที่ส่งประโยคบอกรักพี่ปอผ่าน Channel มาให้ main พิมพ์โชว์หน่อยได้มั้ยคะ?
เงื่อนไข: ใน main พี่ปอต้องสร้าง Channel ขึ้นมา แล้วส่งมันเข้าไปในฟังก์ชัน GetPangMessage เพื่อรับคำหวาน ๆ นั้นออกมาค่ะ
*/
package main

import (
	"fmt"
	"time"
)

func GetPangMessage(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "ปอป่าน่ารัก"
}

func main() {
	// สร้าง channel
	msgChannel := make(chan string)

	fmt.Println("รอรับข้อความจากน้องแป้ง...")

	// สร้าง goroutine เพื่อส่งข้อความ
	go GetPangMessage(msgChannel)

	// รับข้อความจาก channel
	msg := <-msgChannel
	fmt.Println(msg)
}

/*
WaitGroup Vs Channel และวิธีการใช้งาน
	- WaitGroup ใช้แค่รอให้งานเสร็จ แต่ไม่ได้รับส่งข้อมูล
	- Channel ใช้รอแล้วก็รับส่งข้อมูลได้ด้วย
*/
