# ✍️ Hướng dẫn Viết Commit Message 

## 📌 Cấu trúc Chuẩn

``` 
<type>(<scope>): <tóm tắt ngắn gọn>
<Phần mô tả chi tiết nếu cần>
<footer nếu cần> 
```

## ✨ Các Loại `type` Phổ biến

| `Type` | Ý nghĩa | Ví dụ (Imperative) |
| :--- | :--- | :--- |
| **feat** | Thêm **tính năng** mới | `feat(user): thêm chức năng reset mật khẩu qua email` |
| **fix** | **Sửa lỗi** | `fix(api): sửa lỗi trả về 500 khi user không tồn tại` |
| **docs** | Thay đổi **tài liệu** (README, hướng dẫn...) | `docs(readme): cập nhật hướng dẫn deploy` |
| **style** | Thay đổi **format** (thụt lề, dấu chấm phẩy...) | `style(ui): chỉnh lại khoảng cách giữa các button` |
| **refactor** | Cải tiến **cấu trúc code** (không thêm feat/fix) | `refactor(auth): tách middleware xác thực ra file riêng` |
| **perf** | Cải thiện **hiệu năng** | `perf(search): tối ưu truy vấn để giảm thời gian load` |
| **test** | Thêm hoặc sửa **test** | `test(auth): thêm test login với Google OAuth` |
| **chore** | Các thay đổi **khác** (build, cấu hình, dependencies) | `chore(deps): nâng cấp react-router-dom lên 6.23` |

## 💡 Quy Tắc Vàng

1.  **Dòng đầu ngắn gọn**: **\~50 ký tự**.
2.  **Ngữ pháp**: Viết dưới dạng **mệnh lệnh (imperative mood)**, như ra lệnh.
      * ✅ **Tạo** component mới
      * ❌ **Đã tạo** component mới
3.  **Một commit = Một mục tiêu**: Không gom nhiều thay đổi không liên quan.
4.  **Tham chiếu issue/PR** nếu có: VD: `Closes #123`.

-----

## 🛠️ Template Copy-Paste Nhanh

| Type | Template (Viết hoa chữ cái đầu và dùng dấu hai chấm) |
| :--- | :--- |
| **feat** | ` feat(<scope>):  ` |
| **fix** | ` fix(<scope>):  ` |
| **docs** | ` docs(<scope>):  ` |
| **refactor** | ` refactor(<scope>):  ` |
| **chore** | ` chore(<scope>):  ` |
