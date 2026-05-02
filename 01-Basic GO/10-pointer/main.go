/*
ภารกิจ "คนรัก Pointer" จากน้องแป้ง:
ลองสร้าง Method ชื่อ DoubleVersion ให้กับ Struct Mod โดยใช้ Pointer Receiver นะคะ (ใช้ (m *Mod))
เพื่อให้มันไปคูณตัวเลข Version ของม็อดตัวจริงให้เพิ่มขึ้นเป็น 2 เท่าค่ะ
รอดูผลลัพธ์จาก Terminal ของพี่ปอนะจ๊ะ รอบนี้น้องแป้งมั่นใจว่าพี่ปอจะคุม "ที่อยู่" ของข้อมูลได้แม่นยำแน่นอน สู้ ๆ ค่ะพี่ปอ!
*/

package main

import "fmt"

type Mod struct {
	name    string
	version int
}

func (m *Mod) DoubleVersion() {
	m.version = m.version * 2
}

func main() {
	quantummobstacker := Mod{
		name:    "QuantumMobStacker",
		version: 1,
	}
	fmt.Printf("%s Version old %d\n", quantummobstacker.name, quantummobstacker.version)
	quantummobstacker.DoubleVersion()
	fmt.Printf("%s Version new %d\n", quantummobstacker.name, quantummobstacker.version)
}

/*
จุดที่อยากเตือนเพื่อความเข้าใจที่เป๊ะขึ้น:
Function vs Method:
	ที่พี่ปอเขียนมาคือ Function ที่รับ Parameter เป็น Pointer ค่ะ ซึ่งใช้งานได้ปกติเลย
	แต่ถ้าอยากเขียนแบบ Method (เพื่อให้เรียกใช้แบบ quantummobstacker.DoubleVersion()) พี่ปอต้องย้าย (m *Mod) ไปไว้หน้าชื่อฟังก์ชันแบบนี้ค่ะ:
	GO:
		func (m *Mod) DoubleVersion() {
			m.version = m.version * 2
		}
	วิธีนี้จะทำให้โครงสร้างโค้ดของพี่ปอดูเป็นระเบียบและเรียกใช้ง่ายเหมือนตอนเขียน Java เลย
Naming Convention:
	ในภาษา Go ถ้าเราตั้งชื่อฟิลด์ใน Struct เป็นตัวพิมพ์เล็ก (เช่น name, version) มันจะถือว่าเป็น Private (ใช้ได้แค่ใน Package เดียวกัน) ค่ะ
	ถ้าวันหลังพี่ปออยากให้ม็อดนี้ถูกเรียกใช้จากไฟล์อื่นได้ อย่าลืมเปลี่ยนเป็นตัวพิมพ์ใหญ่อย่าง Name และ Version นะจ๊ะ
Syntax about Pointer
	func (m *Type) Method() {
		// ... ใช้ &m
	}
	func (m Type) Method() {
		// ... ใช้ m
	}
*/
