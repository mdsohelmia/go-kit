package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseDriver string

const (
	DatabaseDriverMysql DatabaseDriver = "mysql"
	DatabaseDriverTidb  DatabaseDriver = "tidb"
)

// NewConfig is a constructor for Config
func NewConfig() (*Config, error) {
	config := &Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	return config, nil
}

type Config struct {
	Env      string         `yaml:"env"`
	App      AppConfig      `yaml:"app"`
	Log      LogConfig      `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	Cache    CacheConfig    `yaml:"cache"`
	Worker   WorkerConfig   `yaml:"worker"`
	Services ServicesConfig `yaml:"services"`
}
type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Output     string `yaml:"output"`
	File       string `yaml:"file"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}
type DatabaseConfig struct {
	Backend   BackendConfig     `yaml:"backend"`
	Analytics AnalyticsDBConfig `yaml:"analytics"`
	KeyValue  KeyValueConfig    `yaml:"key_value"`
}
type BackendConfig struct {
	Driver DatabaseDriver `yaml:"driver"`
	Mysql  MysqlConfig    `yaml:"mysql"`
	Tidb   TidbConfig     `yaml:"tidb"`
}

type MysqlConfig struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	Database           string `yaml:"database"`
	Charset            string `yaml:"charset"`
	ParseTime          bool   `yaml:"parseTime"`
	Loc                string `yaml:"loc"`
	MaxConnections     int    `yaml:"max_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
	MaxIdleTime        int    `yaml:"max_idle_time"`
	MigrationPath      string `yaml:"migration_path"`
}
type TidbConfig struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	Database           string `yaml:"database"`
	Charset            string `yaml:"charset"`
	ParseTime          bool   `yaml:"parseTime"`
	Loc                string `yaml:"loc"`
	MaxConnections     int    `yaml:"max_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
	MaxIdleTime        int    `yaml:"max_idle_time"`
	MigrationPath      string `yaml:"migration_path"`
}

type AnalyticsDBConfig struct {
	Driver     string           `yaml:"driver"`
	Clickhouse ClickhouseConfig `yaml:"clickhouse"`
}

type ClickhouseConfig struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	Database           string `yaml:"database"`
	MaxConnections     int    `yaml:"max_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
	MaxIdleTime        int    `yaml:"max_idle_time"`
}

type KeyValueConfig struct {
	Driver string      `yaml:"driver"`
	Redis  RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type CacheConfig struct {
	Driver    string      `yaml:"driver"`
	KeyPrefix string      `yaml:"key_prefix"`
	Redis     RedisConfig `yaml:"redis"`
}

type WorkerConfig struct {
	Driver string         `yaml:"driver"`
	Queues map[string]int `yaml:"queues"`
	Redis  RedisConfig    `yaml:"redis"`
}

type ServicesConfig struct {
	SMS   SMSConfig   `yaml:"sms"`
	Email EmailConfig `yaml:"email"`
}

type SMSConfig struct {
	Driver string       `yaml:"driver"`
	Aliyun AliyunConfig `yaml:"aliyun"`
}

type AliyunConfig struct{}

type EmailConfig struct {
	Driver string     `yaml:"driver"`
	SMTP   SMTPConfig `yaml:"smtp"`
}
type SMTPConfig struct{}

// oauth facebook & google config struct
// social login with facebook,tiktok,google, apple
type OauthConfig struct {
	Facebook OauthFacebookConfig `yaml:"facebook"`
	Google   OauthGoogleConfig   `yaml:"google"`
	Tiktok   OauthTiktokConfig   `yaml:"tiktok"`
	Apple    OauthAppleConfig    `yaml:"apple"`
}
type OauthTiktokConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}

// facebook config struct
type OauthFacebookConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}

// google config struct
type OauthGoogleConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}
type OauthAppleConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}
