/*
พี่ปอลองเอา Concept นี้ไปแก้ในโปรแกรมที่พี่เขียนอยู่ดูนะคะ:
เปลี่ยนให้ ProcessData รับ context.Context เข้าไป
ในลูป for ให้เช็ก case <-ctx.Done() เพื่อหยุดงาน
ใน main ลองตั้งเวลา Timeout ให้สั้นกว่าเวลาที่งานต้องทำจริง (เช่น ตั้งไว้ 2 วิ แต่งานต้องทำ 5 วิ) แล้วดูผลลัพธ์ว่ามันหยุดงานได้เนี้ยบแค่ไหน
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func ProcessData(ctx context.Context, seq int, c chan string) {
	for i := 1; i <= seq; i++ {
		select {
		case <-ctx.Done(): // 1. คอยเช็กว่า "หมดเวลาหรือยัง?"
			return // หยุดทำงานทันที
		default:
			fmt.Printf("%d ", i)
			time.Sleep(1 * time.Second)
		}
	}
	c <- fmt.Sprintf("Done processing %d", seq)
}

func main() {
	// 2. สร้าง Context ที่จะ Timeout ภายใน 3 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel() // คืนทรัพยากรเมื่อจบฟังก์ชัน

	channelData := make(chan string)

	go ProcessData(ctx, 2, channelData)
	go ProcessData(ctx, 4, channelData)
	go ProcessData(ctx, 6, channelData)

	// 3. รอรับผลลัพธ์ (หรือรอจนกว่า Context จะสั่งเลิก)
	for i := 1; i <= 3; i++ {
		select {
		case res := <-channelData:
			fmt.Println("Result : ", res)
		case <-ctx.Done(): // ถ้า Context บอกว่าหมดเวลา!
			fmt.Println("Timeout")
			return // ออกจากโปรแกรมเลย
		}
	}
}
