/*
ภารกิจท้าทายจากคุณครูแป้ง:
ลองเขียนฟังก์ชันชื่อ CheckNumber ที่รับตัวเลข 1 ตัว แล้วส่งค่ากลับมา 2 อย่าง คือ:

เลขนั้นคูณสอง (int)

ข้อความบอกว่าเป็น "เลขคู่" หรือ "เลขคี่" (string)

แล้วลองเอามาเรียกใช้ใน main ดูนะคะ น้องแป้งรอดูโค้ดสวย ๆ ของพี่ปออยู่นะจ๊ะ! สู้ ๆ ค่ะคนเก่งของแป้ง
*/

package main

import "fmt"

func CheckNumber(num int) (int, string) {
	result := num * 2
	var msg string
	if num%2 == 0 {
		msg = "เลขคู่"
	} else {
		msg = "เลขคี่"
	}
	return result, msg
}

func main() {
	var myNumber int
	fmt.Print("พี่ปอจ๋า บอกตัวเลขที่อยากให้ร่ายแม่สูตรคูณมาหน่อยสิคะ: ")

	// รับค่าจากคีย์บอร์ด
	fmt.Scanln(&myNumber)
	result, msg := CheckNumber(myNumber)
	fmt.Printf("%s\n", msg)
	fmt.Printf("เลขนั้นคูณสองคือ: %.2f\n", result)
}
