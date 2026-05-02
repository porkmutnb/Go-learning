/*
ลองสร้าง Slice เก็บ "ภาษาโปรแกรมที่พี่ปอเขียนเป็น" ดูหน่อยค่ะ (เช่น Java, Kotlin, Go) แล้วลองเพิ่มภาษาใหม่เข้าไป 1 ภาษา จากนั้นให้พ่นออกมาโชว์น้องแป้งหน่อยนะจ๊ะ

ลุยเลยค่ะพี่ปอ น้องแป้งสแตนด์บายตรวจการบ้านอยู่ตรงนี้เสมอค่ะ! สู้ ๆ นะคะพี่ชายคนเก่ง
*/

package main

import "fmt"

func main() {
	myLoves := []string{"Java", "Kotlin", "Go"}

	myLoves = append(myLoves, "Node JS")

	fmt.Println("ภาษาโปรแกรมที่พี่ปอเขียนเป็น:")

	for index, value := range myLoves {
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
