/*
น้องมะลิอยากให้พี่ปอช่วยสอนเรื่อง Goroutine และ Channel แบบง่ายๆ หน่อยค่ะ โดยใช้สถานการณ์ดังนี้
การทำงานพร้อมกัน (Concurrency):
ให้สร้างฟังก์ชัน ProcessData(data int) ที่จะทำการ "จำลอง" การทำงานที่ใช้เวลา (เช่น พิมพ์ตัวเลขออกมาทีละตัวพร้อมหน่วงเวลา 1 วินาที)
Main หลัก:
ให้ฟังก์ชัน main สั่งให้ ProcessData(1) ถึง ProcessData(5) ทำงานไปพร้อมๆ กัน
สิ่งสำคัญที่ต้องโชว์:
ให้เห็นว่าตัวเลข 1, 2, 3, 4, 5 ถูกพิมพ์ออกมาแทบจะพร้อมกัน ไม่เรียงลำดับเวลาเป๊ะๆ (เพราะรันคนละคิวพร้อมๆ กัน) ไม่ใช่รอให้ตัวที่ 1 จบก่อนแล้วค่อยทำตัวที่ 2 นะคะ
*/
package main

import (
	"fmt"
	"time"
)

func ProcessData(seq int, c chan string) {
	for i := 1; i <= seq; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(1 * time.Second) // จำลองการทำงานที่ใช้เวลา 1 วินาที
	}
	c <- fmt.Sprintf("Done processing %d", seq)
}

func main() {
	channelData := make(chan string)

	go ProcessData(1, channelData)
	go ProcessData(5, channelData)

	result1 := <-channelData
	result2 := <-channelData

	fmt.Println()

	fmt.Println(result1)
	fmt.Println(result2)
}
