package config

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/panjf2000/ants/v2"

	"github.com/go-redis/redis/v8"
	"github.com/haitao-sun03/donation/backend-end/user/model"
	"github.com/haitao-sun03/donation/backend-end/user/routinepool"
	logging "github.com/haitao-sun03/logging/config"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configs struct {
	DB    DBConfig
	Redis RedisConfig
	Jwt   JwtConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type JwtConfig struct {
	SecretKey string
	WhiteList []string
}

var Config Configs
var RedisClient *redis.Client
var validate *validator.Validate // 定义全局 validator 实例

func Init() {
	// 设置viper读取配置文件
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filename))
	configPath := filepath.Join(projectRoot, "config")

	viper.AutomaticEnv()
	env := viper.GetString("ENV")
	// 配置文件所在的路径
	switch env {
	case "dev":
		viper.AddConfigPath(filepath.Join(configPath, "dev"))
	case "test":
		viper.AddConfigPath(filepath.Join(configPath, "test"))
	default:
		// k8s容器内挂载configMap路径
		viper.AddConfigPath("/etc/app/config")
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	fmt.Println("Loaded config file:", viper.ConfigFileUsed())
	// 解析配置到结构体
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	validate = validator.New()

	// 解析出日志配置
	var loggingConfig logging.LoggingConfig
	if err := viper.UnmarshalKey("logging", &loggingConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	if err := validate.Struct(Config); err != nil {
		panic(fmt.Errorf("config validation failed: %v", err))
	}

	err := InitDatabase()
	if err != nil {
		panic(fmt.Errorf("unable to init database, %v", err))
	}
	InitRedis()
	//将日志配置传递给日志模块并初始化
	logging.InitLogging(loggingConfig)

	InitRoutinePool(1000)
}

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func InitDatabase() error {
	var initErr error
	dbOnce.Do(func() {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
			Config.DB.User,
			Config.DB.Password,
			Config.DB.Host,
			Config.DB.Port,
			Config.DB.DBName,
			// config.DB.SSLMode,
		)

		fmt.Println("dsn :", dsn)

		// 配置 GORM Logger
		// newLogger := logger.New(
		// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // 使用标准输出作为日志写入器
		// 	logger.Config{
		// 		LogLevel:      logger.Info,             // 设置日志级别为 Info
		// 		SlowThreshold: 1000 * time.Millisecond, // 慢 SQL 阈值
		// 		Colorful:      true,                    // 启用彩色输出
		// 	},
		// )

		dbInstance, initErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// Logger: newLogger,
		})

		if initErr != nil {
			fmt.Printf("连接失败: %v\n", initErr)
			return
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			fmt.Printf("获取连接池失败: %v\n", err)
			initErr = err
			return
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

	})
	// 自动创建表
	AutoMigrate()
	return initErr
}

func GetDB() *gorm.DB {
	if err := InitDatabase(); err != nil {
		panic(err) // 或更优雅的错误处理
	}
	return dbInstance
}

func AutoMigrate() {
	dbInstance.AutoMigrate(&model.UserModel{})
	dbInstance.AutoMigrate(&model.UserActivityRoleModel{})
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Address,
		Password: Config.Redis.Password, // 没有密码可以为空字符串
		DB:       Config.Redis.DB,       // 使用默认DB 0
	})

	// 测试连接
	ctx := context.Background()
	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connection successful:", pong)

}

var RoutinePool *ants.Pool

func InitRoutinePool(cap int) {
	RoutinePool = routinepool.NewRoutinePool(cap)
}

func TunePoolCap(cap int) {
	RoutinePool.Tune(cap)
}
