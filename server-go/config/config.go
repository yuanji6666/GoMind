package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/prakis/loadenv"
)

type Config struct{
	MainConfig
	MysqlConfig
	RedisConfig
	EmailConfig
	JwtConfig
	OpenAIConfig
	RabbitmqConfig
	RAGConfig
}

type MysqlConfig struct{
	MysqlHost		string
	MysqlPort		int
	MysqlUser		string
	MysqlPassword	string
	MysqlDBName		string
	MysqlCharset	string
}
type RedisConfig struct{
	RedisHost		string
	RedisPort		int
	RedisDB			int
	RedisPassword	string
}

type EmailConfig struct{
	Authcode 	string
	Email		string
}

type MainConfig struct{
	Host 			string
	Port 			string
	AppName 		string
}


type RedisKeyConfig struct {
	CaptchaPrefix string
}

var DefaultRedisKeyConfig = RedisKeyConfig{
	CaptchaPrefix: "captcha:%s",
}

type JwtConfig struct{
	ExpireDuration		int
	Issuer				string
	Subject				string
	Key 				string
}

type OpenAIConfig struct{
	ApiKey		string
	ModelName	string
	BaseUrl 	string
}

type RabbitmqConfig struct {
	RabbitmqPort     int    
	RabbitmqHost     string 
	RabbitmqUsername string 
	RabbitmqPassword string 
	RabbitmqVhost    string 
}

type RAGConfig struct {
	Host string
	Port string
}

var config *Config

func InitConfig(){
	config = new(Config)

	envPath := resolveEnvFilePath()
	if err := loadenv.Load(envPath); err != nil {
		log.Fatalf("Load .env error: %v (tried %s)", err.Error(), envPath)
	}

	config.MainConfig = MainConfig{
		Host:    getEnv("APP_HOST", ""),
		Port:    getEnv("APP_PORT", ""),
		AppName: getEnv("APP_NAME", ""),
	}

	config.MysqlConfig = MysqlConfig{
		MysqlHost:     getEnv("MYSQL_HOST", ""),
		MysqlPort:     getEnvAsInt("MYSQL_PORT", 0),
		MysqlUser:     getEnv("MYSQL_USER", ""),
		MysqlPassword: getEnv("MYSQL_PASSWORD", ""),
		MysqlDBName:   getEnv("MYSQL_DB_NAME", ""),
		MysqlCharset:  getEnv("MYSQL_CHARSET", ""),
	}

	config.RedisConfig = RedisConfig{
		RedisHost:     getEnv("REDIS_HOST", ""),
		RedisPort:     getEnvAsInt("REDIS_PORT", 0),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
	}

	config.EmailConfig = EmailConfig{
		Authcode: getEnv("EMAIL_AUTHCODE", ""),
		Email:    getEnv("EMAIL", ""),
	}

	config.JwtConfig = JwtConfig{
		ExpireDuration: getEnvAsInt("JWT_EXPIRE_DURATION", 0),
		Issuer:         getEnv("JWT_ISSUER", ""),
		Subject:        getEnv("JWT_SUBJECT", ""),
		Key:            getEnv("JWT_KEY", ""),
	}

	config.OpenAIConfig = OpenAIConfig{
		ApiKey:    getEnv("OPENAI_API_KEY", ""),
		ModelName: getEnv("OPENAI_MODEL_NAME", ""),
		BaseUrl:   getEnv("OPENAI_BASE_URL", ""),
	}

	config.RabbitmqConfig = RabbitmqConfig{
		RabbitmqHost:     getEnv("RABBITMQ_HOST", ""),
		RabbitmqPort:     getEnvAsInt("RABBITMQ_PORT", 0),
		RabbitmqUsername: getEnv("RABBITMQ_USERNAME", ""),
		RabbitmqPassword: getEnv("RABBITMQ_PASSWORD", ""),
		RabbitmqVhost:    getEnv("RABBITMQ_VHOST", ""),
	}

	config.RAGConfig = RAGConfig{
		Host: getEnv("RAG_HOST", ""),
		Port: getEnv("RAG_PORT", ""),
	}
}

func GetConfig() *Config{
	if config == nil {
		config = new(Config)
		InitConfig()
	}
	return config
}

// resolveEnvFilePath picks the .env file to load.
// Order: ENV_FILE if set; else first existing among .env, ../.env (repo root when cwd is server-go).
func resolveEnvFilePath() string {
	if p := os.Getenv("ENV_FILE"); p != "" {
		return p
	}
	candidates := []string{".env", filepath.Join("..", ".env")}
	for _, c := range candidates {
		if st, err := os.Stat(c); err == nil && !st.IsDir() {
			return c
		}
	}
	return ".env"
}

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr, ok := os.LookupEnv(key)
	if !ok || valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Invalid integer env %s=%q: %v", key, valueStr, err)
	}
	return value
}

