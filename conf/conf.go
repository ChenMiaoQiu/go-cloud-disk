package conf

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	GinMode          string
	ServerProt       string
	JwtKey           string
	AdminUserName    string
	AdminPassword    string
	EmailAddr        string
	EmailSecretKey   string
	EmailSMTPServer  string
	FrontWeb         string
	LogLevel         string
	MysqlDSN         string
	BucketName       string
	BucketSecretID   string
	BucketSecretKey  string
	CloudDiskVersion string
	RedisAddr        string
	RedisPassword    string
	RedisDB          string
	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	// get env
	godotenv.Load()
	initEnv()
}

func initEnv() {
	GinMode = os.Getenv("GIN_MODE")
	ServerProt = os.Getenv("SERVER_PROT")
	JwtKey = os.Getenv("JWT_KEY")
	AdminUserName = os.Getenv("ADMIN_USER_NAME")
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	EmailAddr = os.Getenv("EMAIL_ADDR")
	EmailSecretKey = os.Getenv("EMAIL_SECRET_KEY")
	EmailSMTPServer = os.Getenv("EMAIL_SMTP_SERVER")
	FrontWeb = os.Getenv("FRONT_WEB")
	LogLevel = os.Getenv("LOG_LEVEL")
	MysqlDSN = os.Getenv("MYSQL_DSN")
	BucketName = os.Getenv("BUCKET_NAME")
	BucketSecretID = os.Getenv("BUCKET_SECRET_ID")
	BucketSecretKey = os.Getenv("BUCKET_SECRET_KEY")
	CloudDiskVersion = os.Getenv("CLOUD_DISK_VERSION")
	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDB = os.Getenv("REDIS_DB")
	RabbitMQ = os.Getenv("RABBITMQ")
	RabbitMQUser = os.Getenv("RABBITMQ_USER")
	RabbitMQPassword = os.Getenv("RABBITMQ_PASSWORD")
	RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	RabbitMQPort = os.Getenv("RABBITMQ_PORT")
}
