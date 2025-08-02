# go-side-project

本專案是一個以 Go 語言開發的 RESTful API 與 LINE Bot 整合範例，支援用戶註冊、登入、密碼修改、用戶管理、商品管理、分類管理，以及 LINE Bot 待辦事項（Todo）管理。資料庫採用 PostgreSQL，並支援 JWT 驗證與 ngrok 反向代理。

---

## 主要功能

### 1. 用戶管理 API
- **註冊/登入/刪除用戶**：支援基本的用戶註冊、登入（JWT）、刪除等功能。
- **密碼修改**：已登入用戶可修改密碼。
- **用戶查詢**：支援分頁、關鍵字搜尋。

### 2. 商品管理 API
- **新增商品**：支援建立新商品，包含名稱、描述、價格、庫存、分類等資訊。
- **商品分類**：支援多層級分類管理，可建立父子分類關係。
- **資料驗證**：自動驗證必填欄位與外鍵關聯。

### 3. 分類管理 API
- **新增分類**：支援建立商品分類，可設定父分類建立層級結構。
- **分類關聯**：支援一對多關係，一個分類可包含多個商品。

### 4. LINE Bot 待辦事項
- **Webhook**：接收 LINE 訊息事件，根據指令操作待辦事項。
- **指令支援**：
  - `/todo 內容` 新增待辦
  - `/done id` 完成待辦
  - `/edit id 新內容` 編輯待辦（已完成不可編輯）
  - `/show` 顯示今日所有待辦與完成事項
  - `/help` 回傳 Flex Message 格式的指令教學
- **所有訊息回覆皆透過 LINE Message API 主動推送**

### 5. 架構與技術
- **Gin**：HTTP 路由與中介層
- **GORM**：ORM 操作 PostgreSQL
- **JWT**：用戶登入驗證
- **ngrok**：本地開發時自動取得公開網址
- **LINE Bot SDK**：串接 LINE Messaging API
- **Docker Compose**：一鍵啟動 PostgreSQL

---

## 目錄結構簡介

- `/controllers`：API 與 LINE webhook 控制器
- `/models`：資料庫模型（含 enum 型態）
- `/repositories`：資料存取層（用戶、商品、分類、日誌）
- `/utils`：工具（JWT、LINE、回應格式、日誌等）
- `/middleware`：Gin 中介層（Logger、Auth）
- `/routes`：路由設定
- `/database`：資料庫連線與自動遷移
- `/docker-compose.yml`：PostgreSQL 容器設定

---

## API 端點

### 認證相關
- `POST /auth/register` - 用戶註冊
- `POST /auth/login` - 用戶登入

### 需要認證的 API
- `GET /api/users` - 取得用戶列表
- `GET /api/users/:id` - 取得指定用戶
- `DELETE /api/users/:id` - 刪除用戶
- `POST /api/change-password` - 修改密碼
- `POST /api/products` - 新增商品
- `POST /api/categories` - 新增分類

### LINE Bot
- `POST /line/webhook` - LINE Bot Webhook

---

## 快速啟動

1. **設定 .env**  
   參考 `.env` 檔案，設定資料庫連線資訊與 JWT 密鑰。

2. **啟動資料庫（PostgreSQL）**
   ```sh
   docker-compose up -d
    ```

3. **啟動後端服務**
    ```sh
    go run main.go
    ```

4. **啟動 ngrok**
    ```sh
    ngrok http 8080
    ```

5. **設定 LINE Bot Webhook URL 以及 channel secret 和 message api token**
   將 ngrok 提供的 URL 設定到 LINE Developers Console 的 Webhook URL。
   前往 [LINE Developers Console](https://developers.line.biz/console/) 設定你的 LINE Bot。
   將相關資訊填入 `.env` 檔案中。
   確保 channel secret 和 message api token 正確無誤。