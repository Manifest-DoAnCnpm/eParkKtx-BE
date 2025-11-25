
## Mục lục

[README](README.md) | [MESSAGE-RULE](MESSAGE-RULE.md) 

# eParkKTX

**eParkKTX** là một hệ thống quản lý bãi đỗ xe thông minh dành cho sinh viên và cư dân trong khu ký túc xá. Dự án giúp tối ưu hóa việc quản lý xe, tiết kiệm thời gian tìm chỗ đậu và hỗ trợ giám sát an ninh, thông qua giao diện web trực quan và dễ sử dụng.

## Cấu trúc dự án
```
[Request từ Client]
        ↓ dto
[Middleware → kiểm tra token, log,...]
        ↓ dto
[Router → match URL]
        ↓ dto 
[Controller → bind JSON, gọi service]
        ↓
[Service → xử lý nghiệp vụ]
        ↓
[Repository → truy vấn DB]
        ↓
[Response → trả về Client]

```

## Cách chạy
```
go mod init eparkktx
```

```
go get github.com/payOSHQ/payos-lib-golang
```

```
go run main.go
```

# API Endpoints

## 1. Student Management

### 1.1 Create New Student
- **Method**: `POST`
- **Endpoint**: `/api/students`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "userRequest": {
      "name": "Nguyễn Văn A",
      "password": "matkhau12345",
      "phoneNumber": "0123456789",
      "dob": "2000-01-01",
      "gender": "Nam"
    },
    "school": "Đại học Công nghệ",
    "room": "A101"
  }
  ```
- **Success Response (201)**:
  ```json
  {
    "success": true,
    "message": "Student created successfully"
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Invalid request data
  - `409 Conflict`: Username already exists
  - `500 Internal Server Error`: Failed to create student

### 1.2 Search Student by Name
- **Method**: `POST`
- **Endpoint**: `/api/students/search`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "name": "Nguyễn Văn A"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "success": true,
    "data": {
      "user_id": "ea3e6056-c844-4f2a-85c2-22928cf89fc2",
      "name": "Nguyễn Văn A",
      "phone_number": "0123456789",
      "gender": "Nam",
      "dob": "2000-01-01T00:00:00Z",
      "school": "Đại học Công nghệ",
      "room": "A101"
    }
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Invalid request data
  - `404 Not Found`: Student not found
  - `500 Internal Server Error`: Failed to get student information

<!-- ## 2. Payment

### 2.1 Create Payment
- **Method**: `POST`
- **Endpoint**: `/create-payment`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "amount": 50000,
    "description": "Thanh toán phí gửi xe tháng 10"
  }
  ```
 -->



## Tính năng chính

- **Quản lý sinh viên và xe**: Đăng ký, cập nhật thông tin sinh viên và phương tiện.
- **Đặt chỗ và theo dõi bãi xe**: Kiểm tra chỗ trống, đặt chỗ trước, quản lý lượt vào/ra.
- **Báo cáo và thống kê**: Thống kê số lượng xe, lượt ra/vào, thời gian sử dụng bãi xe.
- **Hỗ trợ nhiều vai trò**: Admin, nhân viên quản lý, và người dùng sinh viên.
- **Bảo mật và phân quyền**: Đảm bảo dữ liệu an toàn và phân quyền hợp lý.

## Mục tiêu

1. Tối ưu hóa quản lý bãi xe trong khu ký túc xá.
2. Giúp sinh viên tiết kiệm thời gian tìm bãi đậu.
3. Hỗ trợ quản lý an ninh và báo cáo trực quan.
4. Dễ dàng mở rộng cho nhiều khu ký túc xá và loại xe khác nhau.

## Công nghệ sử dụng

- **Backend**: Gin,...
- **Frontend**: ReactJS,...
- **Database**: SQLite
- **Khác**: REST API, Docker (tùy chọn triển khai), GitHub Actions (CI/CD)

