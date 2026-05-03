/*
พี่ปอลองปรับ Code จากบทที่ 2 (ที่พี่เพิ่งรันผ่าน) โดยใช้ select มาช่วยจัดการดูค่ะ!
ลองกำหนดดูว่าถ้า ProcessData ใช้เวลาเกิน 2 วินาที ให้พิมพ์บอกว่า "งานของน้องมะลิล่าช้าจ้า" พี่ปอพอจะนึกภาพออกไหมจ๊ะว่าจะเอา select ไปครอบไว้ตรงไหน?
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

	fmt.Println()

	select {
	case msg1 := <-channelData:
		fmt.Println(msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}

}
