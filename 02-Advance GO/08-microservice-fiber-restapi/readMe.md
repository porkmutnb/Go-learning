# Go Fiber Framework

## Fiber คืออะไร?
[Fiber](https://gofiber.io/) เป็นเว็บเฟรมเวิร์ก (Web Framework) สำหรับภาษา Go (Golang) ที่ได้รับแรงบันดาลใจมาจาก Express.js ของฝั่ง Node.js ถูกออกแบบมาให้มีความเรียบง่าย ใช้งานสะดวก และเน้นที่ประสิทธิภาพความเร็วในการทำงานขั้นสูงสุด โดยมีโครงสร้างพื้นฐานมาจาก `Fasthttp` ซึ่งเป็น HTTP engine ที่เร็วที่สุดตัวหนึ่งใน Go

**จุดเด่นที่น่าสนใจของ Fiber:**
- **Routing ที่ใช้งานง่ายและทรงพลัง:** จัดการ API Endpoints (GET, POST, PUT, DELETE) ได้สะดวกเหมือนใช้งาน Express.js
- **ประสิทธิภาพสูง (High Performance):** ทำงานรวดเร็ว รองรับ Request ได้มหาศาล และกินทรัพยากรน้อย
- **Zero Memory Allocation:** ช่วยลดภาระของ Garbage Collector ทำให้ระบบทำงานได้เสถียรยิ่งขึ้น
- **มี Middleware ให้เลือกใช้ครบครัน:** รองรับการทำงานที่จำเป็นเช่น CORS, Logger, JWT, Rate Limiting เป็นต้น

---

## การติดตั้ง (Installation)

ก่อนที่จะเริ่มติดตั้ง Fiber ตรวจสอบให้แน่ใจว่าเครื่องของคุณมี Go เวอร์ชัน **1.14 ขึ้นไป** 

1. **เริ่มต้นโปรเจกต์ (Initialize Go Module)** 
   สร้างโฟลเดอร์โปรเจกต์และรันคำสั่งนี้ใน Terminal (ถ้ายังไม่มีไฟล์ `go.mod`):
   ```bash
   go mod init example.com/my-fiber-api
   ```

2. **ติดตั้งแพ็กเกจ Go Fiber**
   ใช้คำสั่ง `go get` เพื่อดาวน์โหลด Fiber เวอร์ชันล่าสุด:
   ```bash
   go get -u github.com/gofiber/fiber/v2
   ```

---

## การใช้งานเบื้องต้น (Basic Usage)

ด้านล่างนี้คือตัวอย่างการสร้าง REST API เบื้องต้นด้วย Fiber 

สร้างไฟล์ `main.go` และใส่โค้ดดังนี้:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // 1. สร้าง Instance ของแอปพลิเคชัน Fiber
    app := fiber.New()

    // 2. กำหนด Route พื้นฐาน (GET Method)
    app.Get("/", func(c *fiber.Ctx) error {
        // ส่งข้อความกลับเป็น String
        return c.SendString("Hello, Go Fiber! 👋")
    })

    // ตัวอย่าง Route สำหรับส่งข้อมูลกลับเป็น JSON
    app.Get("/api/user", func(c *fiber.Ctx) error {
        // ใช้ fiber.Map เพื่อสร้าง JSON สั้นๆ
        return c.JSON(fiber.Map{
            "status":  "success",
            "message": "Welcome to our API!",
            "data": fiber.Map{
                "name": "John Doe",
                "role": "Developer",
            },
        })
    })

    // 3. เริ่มรัน Server ที่ Port 3000
    log.Fatal(app.Listen(":3000"))
}
```

### วิธีการรันโปรแกรม
พิมพ์คำสั่งนี้ใน Terminal เพื่อเริ่มการทำงานของเซิร์ฟเวอร์:
```bash
go run main.go
```

หลังจากเซิร์ฟเวอร์รันขึ้นมาแล้ว คุณสามารถทดสอบการทำงานได้โดยเปิด Browser, Postman หรือใช้ `curl` ไปที่:
- **`http://localhost:3000/`** (เพื่อดูข้อความ Hello)
- **`http://localhost:3000/api/user`** (เพื่อดูข้อมูลรูปแบบ JSON)
