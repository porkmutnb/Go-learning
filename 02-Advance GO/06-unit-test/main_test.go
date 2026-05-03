/*
1. กฎเหล็ก 3 ข้อของการเขียน Test ใน Go
ชื่อไฟล์: ต้องลงท้ายด้วย _test.go เสมอ (เช่น main_test.go)
ชื่อฟังก์ชัน: ต้องขึ้นต้นด้วยคำว่า Test ตามด้วยชื่อฟังก์ชันที่ต้องการเทส (เช่น TestDivide)
Parameter: ต้องรับค่า t *testing.T เข้ามาเสมอ เพื่อเอาไว้สั่งแจ้งเตือนเวลาเทสไม่ผ่านค่ะ
2. วิธีการรัน
พี่ปอแค่พิมพ์คำสั่งใน Terminal ว่า: go test หรือถ้าอยากเห็นรายละเอียดแบบฉ่ำๆ ก็ go test -v จ้ะ
*/

package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSecurityError5(t *testing.T) {
	err := Inspect("ปืนฉีดน้ำ")
	expected := 5
	// 1. ดักจับ: ถ้าไม่มี Error ออกมาเลย ทั้งที่ส่ง "ปืนฉีดน้ำ" เข้าไป... แบบนี้คือพัง!
	if err == nil {
		t.Fatal("จริงๆ มันต้องมี Error ออกมานะจ๊ะพี่ปอ!")
	}
	// 2. ตรวจสอบไส้ใน: แกะเอา SecurityError ออกมาเช็ก Level
	var myErr *SecurityError
	if errors.As(err, &myErr) {
		// ถ้าแกะออกมาได้แล้ว แต่ Level ดันไม่ใช่อย่างที่หวัง... สั่งพังเลย!
		if myErr.Level != expected {
			t.Errorf("Level ผิดจ้า! อยากได้ %d แต่ได้ %d", expected, myErr.Level)
		}
	} else {
		// ถ้าแกะออกมาเป็น SecurityError ไม่ได้เลย... ก็พังเหมือนกัน!
		t.Error("Error ที่ได้มาไม่ใช่ชนิด SecurityError นะเนี่ย")
	}
}

func TestSecurityError1(t *testing.T) {
	err := Inspect("กระบอกฉีดน้ำ")
	expected := 1
	// 1. ดักจับ: ถ้าไม่มี Error ออกมาเลย ทั้งที่ส่ง "ปืนฉีดน้ำ" เข้าไป... แบบนี้คือพัง!
	if err == nil {
		t.Fatal("จริงๆ มันต้องมี Error ออกมานะจ๊ะพี่ปอ!")
	}
	// 2. ตรวจสอบไส้ใน: แกะเอา SecurityError ออกมาเช็ก Level
	var myErr *SecurityError
	if errors.As(err, &myErr) {
		// ถ้าแกะออกมาได้แล้ว แต่ Level ดันไม่ใช่อย่างที่หวัง... สั่งพังเลย!
		if myErr.Level != expected {
			t.Errorf("Level ผิดจ้า! อยากได้ %d แต่ได้ %d", expected, myErr.Level)
		}
	} else {
		// ถ้าแกะออกมาเป็น SecurityError ไม่ได้เลย... ก็พังเหมือนกัน!
		t.Error("Error ที่ได้มาไม่ใช่ชนิด SecurityError นะเนี่ย")
	}
}
