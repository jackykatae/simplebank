# SimpleBank

โปรเจกต์นี้เป็นตัวอย่างการสร้าง RESTful API ด้วยภาษา Go โดยใช้ [sqlc](https://sqlc.dev/) สำหรับการเชื่อมต่อและจัดการกับฐานข้อมูล PostgreSQL

## คุณสมบัติ

- สร้างและจัดการบัญชีธนาคาร
- ฝาก ถอน โอนเงินระหว่างบัญชี
- ใช้ sqlc สร้างโค้ด Go จากไฟล์ SQL
- เชื่อมต่อฐานข้อมูล PostgreSQL

## การติดตั้ง

1. ติดตั้ง Go และ PostgreSQL
2. ติดตั้ง sqlc ด้วยคำสั่ง:
    ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```
3. สร้างไฟล์ `.env` สำหรับข้อมูลการเชื่อมต่อฐานข้อมูล

## การใช้งาน

1. สร้างฐานข้อมูล PostgreSQL
2. สร้างโค้ด Go จาก SQL ด้วยคำสั่ง:
    ```bash
    sqlc generate
    ```
3. รันเซิร์ฟเวอร์:
    ```bash
    go run main.go
    ```

## โครงสร้างโปรเจกต์

- `db/` : ไฟล์ SQL และโค้ดที่ sqlc สร้างขึ้น
- `api/` : โค้ดสำหรับ REST API
- `main.go` : จุดเริ่มต้นของโปรเจกต์

## เอกสารเพิ่มเติม

- [sqlc Documentation](https://docs.sqlc.dev/)
- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
