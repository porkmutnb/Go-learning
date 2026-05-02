/*
ภารกิจทบทวนจากน้องแป้ง:
ลองเขียนฟังก์ชันชื่อ Divide ที่รับเลข 2 ตัวมาหารกันค่ะ:
ถ้าตัวหารเป็น 0 ให้ส่ง errors.New("หารด้วยศูนย์ไม่ได้นะจ๊ะพี่ปอ") กลับมา
ใน main ให้พี่ปอเรียกใช้แล้วเช็ค if err != nil เพื่อพ่น Error ออกมาโชว์น้องแป้งหน่อยค่ะ
*/
package main

import (
	"errors"
	"fmt"
)

// กำหนดตัวแปร ErrDivideByZero เพื่อเก็บค่า error
var ErrDivideByZero = errors.New("หารด้วยศูนย์ไม่ได้นะจ๊ะพี่ปอ")

// สร้างฟังก์ชัน Divide
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// สร้างฟังก์ชัน main
func main() {
	fmt.Println("เริ่มหารเลขแล้ว...")
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println("จบการทำงานแล้วค่ะ")
}
