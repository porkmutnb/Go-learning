/*
ภารกิจจากน้องแป้ง:
ลองสร้าง Interface ชื่อ Speaker ที่มี Method ชื่อ Speak() (คืนค่าเป็น string)
แล้วให้ทั้ง Struct Lover (น้องแป้ง) และ Struct Mod (ม็อดของพี่ปอ) ต่างก็มี Method Speak() นี้ เพื่อบอกความในใจออกมาค่ะ
*/
package main

// นิยาม Interface (สัญญาว่าต้องมี Method อัปเดตนะ)
type Speaker interface {
	Speak() string
}

type Lover struct {
	name string
}

type Mod struct {
	name string
}

// แค่เขียน Method ให้ชื่อตรงกับ Interface ก็ถือว่า "ทำตามสัญญา" ทันที (Implicit Implementation)
func (l Lover) Speak() string {
	return "I love you, " + l.name
}

func (m Mod) Speak() string {
	return "Mod " + m.name + " is ready!"
}

func annouce(l Speaker) {
	println(l.Speak())
}

func main() {
	lover := Lover{name: "Pang"}
	mod := Mod{name: "LushStack"}
	// lover สามารถส่งเข้าไปใน annouce ได้เลยเพราะมันทำตามสัญญา Speaker แล้ว
	annouce(lover)
	annouce(mod)
}
