/*
ภารกิจจากน้องแป้ง:
ลองสร้าง Map เก็บ "Spec คอมพิวเตอร์ของพี่ปอ" ดูหน่อยค่ะ เช่น:

"CPU": "Ryzen 7 9800X3D"

"GPU": "RTX 5060 Ti"

"Monitor": "Samsung Viewfinity s6"

แล้วลองสั่งให้มัน Println ค่า GPU ออกมาโชว์น้องแป้งหน่อยนะจ๊ะ รอบนี้น้องแป้งขอเช็คความจำพี่ปอหน่อยว่าจำ Spec คอมตัวเองแม่นรึเปล่า! (หยอก ๆ นะคะ)

ลุยเลยค่ะพี่ปอ น้องแป้งรอตรวจการบ้านด้วยรอยยิ้มเหมือนเดิมจ้า!
*/

package main

import "fmt"

func main() {
	myPC := make(map[string]string)

	myPC["CPU"] = "Intel Core i7-13700K"
	myPC["RAM"] = "16GB"
	myPC["Storage"] = "1TB SSD"
	myPC["GPU"] = "NVIDIA GeForce RTX 4070 Ti"
	myPC["Monitor"] = "27-inch 144Hz QHD"
	myPC["mainboard"] = "MSI B760 GAMING PLUS WIFI DDR5"
	myPC["Cooler"] = "AIO 240mm"

	fmt.Println("สเปคคอมพิวเตอร์ของฉัน:")
	fmt.Printf("CPU: %s\n", myPC["CPU"])
	fmt.Printf("RAM: %s\n", myPC["RAM"])
	fmt.Printf("Storage: %s\n", myPC["Storage"])
	fmt.Printf("GPU: %s\n", myPC["GPU"])
	fmt.Printf("Monitor: %s\n", myPC["Monitor"])
	fmt.Printf("Mainboard: %s\n", myPC["mainboard"])
	fmt.Printf("Cooler: %s\n", myPC["Cooler"])

	delete(myPC, "Cooler")
	fmt.Printf("Cooler: %s\n", myPC["Cooler"])
}
