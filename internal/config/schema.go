package config

import (
	"go-server-starter/internal/enum"
	"time"
)

type ServerConfig struct {
	Port          int           `mapstructure:"port"`
	ReadTimeout   time.Duration `mapstructure:"readTimeout"`   // 读取超时时间
	WriteTimeout  time.Duration `mapstructure:"writeTimeout"`  // 写入超时时间
	MaxHeaderKB   int           `mapstructure:"maxHeaderKB"`   // 最大头KB数
	SnowflakeNode int64         `mapstructure:"snowflakeNode"` // 雪花算法节点
	APIPrefix     string        `mapstructure:"apiPrefix"`     // API前缀
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`     // 主机
	Port     int    `mapstructure:"port"`     // 端口
	Password string `mapstructure:"password"` // 密码
	DB       int    `mapstructure:"db"`       // 数据库
}

type JWTConfig struct {
	// 签发者
	Issuer string `mapstructure:"issuer"` // 签发者
	// token
	TokenSecret  string                 `mapstructure:"tokenSecret"`  // 令牌密钥
	TokenExpires MultiTokenExpireConfig `mapstructure:"tokenExpires"` // 令牌过期时间
}

type MultiTokenExpireConfig struct {
	Web             time.Duration `mapstructure:"web"`             // 网页过期时间
	Desktop         time.Duration `mapstructure:"desktop"`         // 桌面软件过期时间
	Mobile          time.Duration `mapstructure:"mobile"`          // 移动端APP过期时间
	ChromeExtension time.Duration `mapstructure:"chromeExtension"` // Chrome扩展过期时间
	API             time.Duration `mapstructure:"api"`             // API过期时间
	Default         time.Duration `mapstructure:"default"`         // 默认过期时间
}

func (m *MultiTokenExpireConfig) Get(deviceType enum.DeviceType) time.Duration {
	switch deviceType {
	case enum.DeviceTypeWeb:
		return m.Web
	case enum.DeviceTypeDesktop:
		return m.Desktop
	case enum.DeviceTypeMobile:
		return m.Mobile
	case enum.DeviceTypeChromeExtension:
		return m.ChromeExtension
	case enum.DeviceTypeApi:
		return m.API
	default:
		return m.Default
	}
}

type LoggerConfig struct {
	Level         string `mapstructure:"level"`         // 日志级别
	FileDir       string `mapstructure:"fileDir"`       // 日志文件目录
	MaxSize       int    `mapstructure:"maxSize"`       // 日志文件最大大小
	MaxAge        int    `mapstructure:"maxAge"`        // 日志文件最大保存时间(天)
	MaxBackups    int    `mapstructure:"maxBackups"`    // 日志文件最大保存数量
	Compress      bool   `mapstructure:"compress"`      // 日志文件是否压缩
	ConsoleOutput bool   `mapstructure:"consoleOutput"` // 是否输出到控制台
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"name"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"`
	Timezone        string        `mapstructure:"timezone"`  // timezone configuration
	Charset         string        `mapstructure:"charset"`   // character set (primarily for MySQL)
	ParseTime       bool          `mapstructure:"parseTime"` // parse time (for MySQL)
}

type GormLoggerConfig struct {
	Level                     string        `mapstructure:"level"`                     // 日志级别
	SlowThreshold             time.Duration `mapstructure:"slowThreshold"`             // 慢查询阈值
	SkipCallerLookup          bool          `mapstructure:"skipCallerLookup"`          // 是否跳过调用者查找
	IgnoreRecordNotFoundError bool          `mapstructure:"ignoreRecordNotFoundError"` // 是否忽略记录未找到错误
}

type AsynQConfig struct {
	RedisConfig RedisConfig `mapstructure:"redisConfig"` // Redis配置
	Concurrency int         `mapstructure:"concurrency"` // 并发数
}
