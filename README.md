# fluffy-octo-tribble

run

```bash
$ go run cmd/main.go
```

preview: http://localhost:7010/

---

| path | method | description |
| --- | --- | --- |
| /api/users/register		| post 		| สมัครสมาชิก |
| /api/users/login			| post 		| ล็อกอินเข้าระบบ |
| /api/users/profile		| get  		| ดูโปรไพล์ (ใช้ token จากการล็อกอิน) |
| /api/users/profile		| patch 	| อัพเดตโปรไพล์ (ใช้ token จากการล็อกอิน) |
| /api/projects 				| get 		| แสดงรายการโปรเจคทั้งหมด |
| /api/projects 				| post 		| สร้างโปรเจค |
| /api/projects/:id			| get 		| แสดงโปรเจค |
| /api/projects/:id			| path 		| อัพเดตโปรเจค |
| /api/projects/:id			| delete	| ลบโปรเจค |
| /api/tasks		 				| get 		| แสดงรายการงานทั้งหมด |
| /api/tasks						| post 		| สร้างงาน |
| /api/tasks/:id				| get 		| แสดงงาน |
| /api/tasks/:id				| delete	| ลบงาน |
| /api/tasks/:id				| path 		| อัพเดตงาน |
| /api/tasks/:id/status	| get 		| ดูสถานะงาน |
| /api/tasks/:id/status	| path 		| อัพเดตสถานะงาน |


---

# Todo

- [Y] สมัครสมาชิกได้
- [Y] ล็อกอินเข้าระบบได้
- [Y] ดูโปรไพล์ของตัวเองได้
- [Y] อัพเดตโปรไพล์ของตัวเองได้
- [N] เปลี่ยนรหัสผ่านได้
- [N] 

### register
`/api/users/register`
```bash
curl --request POST \
  --url http://localhost:7010/users/register \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "scott@gmail.com",
	"password": "scott",
	"confirm_password": "scott",
	"name": "Scott",
	"location": "Bangkok"
}'
```