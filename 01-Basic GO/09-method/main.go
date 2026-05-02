/*
ภารกิจจากน้องแป้ง:
พี่ปอจำโปรเจค LushStack หรือ QuantumMobStacker ที่พี่เคยทำได้มั้ยคะ? ลองสร้าง Struct ชื่อว่า Mod ที่มีฟิลด์ Name และ Version
จากนั้นให้เขียน Method ชื่อว่า UpdateVersion เพื่อให้มันพ่นข้อความว่า "กำลังอัปเดต [ชื่อม็อด] เป็นเวอร์ชันใหม่..." ออกมาให้ดูหน่อยได้มั้ยคะ?
น้องแป้งอยากเห็นม็อดของพี่ปอในเวอร์ชันภาษา Go จังเลยค่ะ! สู้ ๆ นะคะพี่ชายคนเก่งของแป้ง พร้อมแล้วลุยเลยจ้า!
*/

package main

import "fmt"

type Mod struct {
	name    string
	version string
}

func (m Mod) UpdateVersion() {
	fmt.Printf("กำลังอัปเดต %s เป็นเวอร์ชันใหม่...\n", m.name)
}

func main() {
	lushstack := Mod{
		name:    "LushStack",
		version: "1.0.0",
	}
	lushstack.UpdateVersion()

	quantummobstacker := Mod{
		name:    "QuantumMobStacker",
		version: "1.0.0",
	}
	quantummobstacker.UpdateVersion()
}
