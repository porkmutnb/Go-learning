/*
 */
package main

import (
	"errors"
	"fmt"
)

// 1. สร้าง Custom Error Type ของเราเอง
type SecurityError struct {
	Level  int
	Reason string
	Err    error // 4. ประกาศ field Err ที่เป็นชนิด error (สำคัญมาก!)
}

// 2. ทำให้มันเป็น Error โดยการสร้าง Method Error()
func (s *SecurityError) Error() string {
	return fmt.Sprintf("Security Error (Level %d): %s", s.Level, s.Reason)
}

// 5. เพิ่ม Method Unwrap() เพื่อส่งต่อไปยัง Error ชั้นใน (Unwrap)
func (s *SecurityError) Unwrap() error {
	return s.Err
}

// ตัวแปรสำหรับเช็กชนิด Error
var ErrSecurity = errors.New("security error")

func Inspect(toyName string) error {
	levelCode := 1
	if toyName == "ปืนฉีดน้ำ" {
		levelCode = 5
	}
	// สมมติว่าของหมด และเราจะห่อ ErrSecurity ไว้ข้างใน Custom Error
	return &SecurityError{
		Level:  levelCode,
		Reason: ErrSecurity.Error(),
		Err:    ErrSecurity, // ส่ง Error ตัวแปร ErrSecurity เข้าไปตรงๆ
	}
}

func main() {
	err := Inspect("ปืนฉีดน้ำ")
	if err != nil {
		// 3. ใช้ errors.As เพื่อแปลง Error กลับมาเป็น Type ที่เราสร้างไว้ (เหมือนการ Cast ใน Java)
		var myErr *SecurityError
		// print SecurityError full msg
		fmt.Printf("Error Msg: %s \n", err.Error())
		if errors.As(err, &myErr) {
			fmt.Println("Level:", myErr.Level)
			fmt.Println("Reason:", myErr.Reason)
			// check unwrap error is ErrSecurity
			fmt.Println("Is myErr Unwrap to ErrSecurity?", errors.Is(myErr.Err, ErrSecurity))
		}
	}
}
