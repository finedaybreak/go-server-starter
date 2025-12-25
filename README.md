# Go Server Starter

[中文文档](./README.zh_CN.md)

A production-ready Go web server boilerplate/starter kit with clean architecture, built-in authentication, RBAC, and comprehensive tooling.

## ✨ Features

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
- **Database**: MySQL with [GORM](https://gorm.io/) ORM, auto-migration support
- **Cache**: [Redis](https://github.com/redis/go-redis) integration
- **Authentication**: JWT-based auth with multi-device token expiration support
- **Authorization**: Role-Based Access Control (RBAC)
- **Validation**: Request validation via [go-playground/validator](https://github.com/go-playground/validator)
- **Internationalization**: i18n support (English & Chinese)
- **Logging**: Structured logging with [Zap](https://github.com/uber-go/zap) + log rotation with [Lumberjack](https://github.com/natefinch/lumberjack)
- **Configuration**: Environment-based config management with [Viper](https://github.com/spf13/viper)
- **Async Tasks**: Background job processing with [Asynq](https://github.com/hibiken/asynq)
- **ID Generation**: Distributed ID generation with [Snowflake](https://github.com/bwmarrin/snowflake)
- **Graceful Shutdown**: Clean server shutdown handling
- **Clean Architecture**: Layered structure (Handler → Service → Repository)

## 📁 Project Structure

```
go-server-starter/
├── cmd/
│   └── server/          # Application entry point
├── configs/             # Configuration files
│   ├── config.yml       # Default config
│   ├── config.dev.yml   # Development config
│   └── config.test.yml  # Test config
├── internal/
│   ├── app/             # Application initialization
│   ├── config/          # Config struct definitions
│   ├── constant/        # Constants
│   ├── ctx/             # Custom context
│   ├── dto/             # Data Transfer Objects
│   ├── enum/            # Enumerations
│   ├── exception/       # Exception handling
│   ├── handler/         # HTTP handlers (controllers)
│   ├── i18n/            # Internationalization
│   ├── middleware/      # HTTP middlewares
│   ├── model/           # Database models
│   ├── repo/            # Repository layer (data access)
│   ├── router/          # Route definitions
│   ├── seed/            # Database seeders
│   └── service/         # Business logic layer
├── pkg/
│   ├── asyn_queue/      # Asynq client/server
│   ├── auth/            # Authorization utilities
│   ├── database/        # Database connection
│   ├── jwt/             # JWT utilities
│   ├── logger/          # Logger configuration
│   ├── redis/           # Redis client
│   ├── snowflake/       # Snowflake ID generator
│   ├── translator/      # Translator utilities
│   ├── utils/           # Common utilities
│   └── validator/       # Validation rules
└── logs/                # Log files
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

### Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/go-server-starter.git
cd go-server-starter
```

2. Install dependencies:

```bash
go mod download
```

3. Configure the application:

Copy and modify the config file for your environment:

```bash
cp configs/config.yml configs/config.dev.yml
```

Edit `configs/config.dev.yml` with your database and Redis settings.

4. Run the server:

```bash
# Development mode
go run cmd/server/server.go -mode=dev

# Production mode
go run cmd/server/server.go -mode=prod

# Test mode
go run cmd/server/server.go -mode=test
```

The server will start on `http://localhost:8080` by default.

## ⚙️ Configuration

Configuration is managed via YAML files. Key settings include:

```yaml
server:
  port: 8080
  readTimeout: 10s
  writeTimeout: 10s
  apiPrefix: "/api"

jwt:
  issuer: go-server-starter
  tokenSecret: your-secret-key
  tokenExpires:
    web: 24h
    mobile: 360h
    desktop: 360h

database:
  host: localhost
  port: 3306
  username: root
  password: your-password
  databaseName: your-db

redis:
  host: localhost
  port: 6379
  password: your-password
  db: 0
```

## 🔐 Authentication & Authorization

### JWT Authentication

The project uses JWT for stateless authentication with device-specific token expiration:

- **Web**: 24 hours
- **Mobile/Desktop**: 15 days
- **Chrome Extension**: 30 days
- **API**: 48 hours

### Role-Based Access Control

Built-in roles:
- `super_admin` - Super administrator
- `admin` - Administrator
- `user` - Regular user
- `user_vip` - VIP user
- `user_svip` - SVIP user
- `guest` - Guest user

Protect routes with role checks:

```go
// Any of the specified roles
router.GET("/admin", auth.RoleCheckAny(enum.RoleCodeAdmin, enum.RoleCodeSuperAdmin), handler)

// All specified roles required
router.GET("/super", auth.RoleCheckAll(enum.RoleCodeSuperAdmin), handler)
```

## 🌐 API Endpoints

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/api/hello` | Health check | No |
| POST | `/api/auth/login/mobile` | Login via mobile + code | No |
| POST | `/api/auth/login/email` | Login via email + code | No |
| GET | `/api/user/info` | Get current user info | Yes |
| PUT | `/api/user/info` | Update user info | Yes |
| GET | `/api/user/table` | Get users list (paginated) | Yes |

## 🌍 Internationalization

The API supports multiple languages via the `Accept-Language` header:

```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/api/hello

# Chinese
curl -H "Accept-Language: zh" http://localhost:8080/api/hello
```

## 📝 Logging

Logs are structured with Zap and automatically rotated:

- **Log levels**: debug, info, warn, error, fatal
- **Output**: Console (dev) + File (info.log, error.log)
- **Rotation**: Configurable max size, age, and backup count

## 🛠️ Development

### Generate Code

```bash
./generate.sh
```

### Hot Reload

For development with hot reload, use [Air](https://github.com/cosmtrek/air):

```bash
air
```

## 📄 License

This project is open-sourced software licensed under the [MIT license](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

