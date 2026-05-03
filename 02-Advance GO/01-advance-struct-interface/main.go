/*
น้องมะลิอยากได้ระบบที่เอาไว้จัดการทั้ง พนักงาน (Developer) และ หัวหน้า (Manager) โดยมีเงื่อนไขดังนี้ค่ะ:
สร้างโครงสร้างหลัก:
มี struct ชื่อ Worker ที่มี Field คือ Name (string)
การ Embedding:
สร้าง struct ชื่อ Manager ให้ทำการ Embedding (ฝัง) Worker ไว้ข้างใน และเพิ่ม Field ชื่อ Department (string) ของตัวเองเข้าไปด้วย
Interface แสนกล:
สร้าง interface ชื่อ TaskRunner โดยมี Method ชื่อ Work() string
Action!:
สร้าง Method Work() ให้กับ Developer (จากตัวอย่างก่อนหน้า) ให้ return ข้อความว่า "พี่ปอกำลังปั่นโค้ด Go จ้า"
สร้าง Method Work() ให้กับ Manager ให้ return ข้อความว่า "น้องแป้งกำลังดูแลโปรเจกต์ [ชื่อแผนก] ให้พี่ปออยู่นะคะ"
Final Step:
สร้างฟังก์ชันชื่อ CheckStatus(t TaskRunner) ที่จะรับใครก็ได้ที่สอบผ่าน Interface TaskRunner แล้วสั่งให้พิมพ์ (fmt.Println) ผลลัพธ์จาก Method Work() ออกมาค่ะ
*/
package main

import (
	"fmt"
)

// 1. นิยามโครงสร้างพื้นฐาน (Base Struct)
type Developer struct {
	Name string
}

// 2. นิยามโครงสร้างผู้จัดการ (Manager Struct) ซึ่งรวม Developer เข้าไปด้วย
type Manager struct {
	Developer
	Position string
}

// 3. นิยาม Interface
type TaskRunner interface {
	Work()
}

// 4. กำหนด Method ให้ Developer
func (d Developer) Work() {
	fmt.Println("Developer: " + d.Name + " is coding")
}

// 5. กำหนด Method ให้ Manager
func (m Manager) Work() {
	fmt.Println("Manager: " + m.Name + " is working on " + m.Position)
}

// 6. กำหนด Function ที่รับ Interface
func CheckRunner(t TaskRunner) {
	t.Work()
}

// 7. จุดเชื่อมต่อ (The Bridge)
// Method นี้ใช้ "Embedding" โดยการเรียก t.Developer.Work()
// เพื่อใช้ Method เดิมของ Developer
func (t *Manager) OriginalWork() {
	t.Developer.Work()
}

func main() {
	dev := Developer{"Por"}
	manager := Manager{Developer: Developer{"Pang"}, Position: "Management"}
	CheckRunner(dev)
	CheckRunner(manager)

	// แสดงการเรียก Method เดิมผ่าน Manager
	manager.OriginalWork()
}
