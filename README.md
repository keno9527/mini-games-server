# mini-games-server

基于 Go + Gin 的小游戏后端服务，提供游戏列表、用户管理、游戏记录与统计等 RESTful API。数据通过本地 JSON 文件持久化。

## 技术栈

- Go 1.21
- [Gin](https://github.com/gin-gonic/gin) Web 框架
- [gin-contrib/cors](https://github.com/gin-contrib/cors) 跨域支持
- [google/uuid](https://github.com/google/uuid) UUID 生成
- JSON 文件存储（`data/users.json`、`data/records.json`）

## 目录结构

```
mini-games-server/
├── main.go              // 入口，路由注册与中间件配置
├── handlers/            // HTTP 处理器
│   ├── games.go         // 游戏相关接口
│   ├── users.go         // 用户相关接口
│   └── records.go       // 游戏记录与统计
├── models/              // 数据模型定义
│   └── models.go
├── storage/             // 数据持久化（JSON 文件）
│   └── storage.go
├── logfile/             // 日志文件初始化
│   └── logfile.go
├── data/                // 本地数据文件
│   ├── users.json
│   └── records.json
├── go.mod
└── go.sum
```

## 快速开始

### 环境要求

- Go 1.21 及以上

### 安装依赖

```bash
go mod download
```

### 启动服务

```bash
go run main.go
```

服务默认运行在 `http://localhost:8080`，日志会写入 `log/` 目录。

## CORS 配置

默认允许以下前端地址访问：

- `http://localhost:5173`
- `http://localhost:3000`

如需修改，请编辑 `main.go` 中的 `cors.Config`。

## API 接口

所有接口前缀为 `/api`。

### 游戏

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/games` | 获取游戏列表 |
| GET | `/api/games/:id` | 获取指定游戏详情 |

### 用户

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/users` | 获取用户列表 |
| POST | `/api/users` | 创建用户 |
| DELETE | `/api/users/:id` | 删除用户 |
| GET | `/api/users/:id/stats` | 获取用户游戏统计 |

### 游戏记录

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/users/:id/records` | 获取某用户的游戏记录 |
| POST | `/api/users/:id/records` | 新增某用户的一条游戏记录 |

## 数据模型

### Game

```go
type Game struct {
    ID           string
    Name         string
    Description  string
    CoverImage   string
    Tags         []string
    Difficulties []string // 简单、中等、复杂
}
```

### User

```go
type User struct {
    ID        string
    Name      string
    Avatar    string
    CreatedAt time.Time
}
```

### GameRecord

```go
type GameRecord struct {
    ID       string
    UserID   string
    GameID   string
    Score    int
    Duration int       // 秒
    PlayedAt time.Time
    Result   string    // "win" | "lose" | "complete"
}
```

### UserStats

聚合输出用户的总局数、总时长、总分及按游戏维度的统计信息。

## 日志

- 所有 API 请求在完成后会输出一条 `[API] METHOD PATH STATUS COST` 日志。
- 日志通过 `logfile.Init("log")` 写入到项目根目录下的 `log/` 文件夹。

## 许可证

本项目仅供学习与示例用途。
