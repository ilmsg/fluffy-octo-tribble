# fluffy-octo-tribble

run

```bash
$ go run cmd/main.go
```

preview: http://localhost:7010/

---

| path | method | description |
| --- | --- | --- |
| /api/users/register	| post | สมัครสมาชิก |
| /api/users/login		| post | ล็อกอินเข้าระบบ |
| /api/users/profile	| get  | ดูโปรไพล์ (ใช้ token จากการล็อกอิน) |
| /api/users/profile	| patch | อัพเดตโปรไพล์ (ใช้ token จากการล็อกอิน) |

---

# Todo

- [N] validate data
- [N] structure data response

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