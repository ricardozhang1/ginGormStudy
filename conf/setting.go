package conf

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	SingularTable   bool   `mapstructure:"SINGULAR_TABLE"`
	LogMode         bool   `mapstructure:"LOG_MODE"`
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"`
	ConnMaxLifetime int    `mapstructure:"CONN_MAX_LIFETIME"`

	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

	LogSavePath string `mapstructure:"LOG_SAVE_PATH"`
	LogSaveName string `mapstructure:"LOG_SAVE_NAME"`
	LogFileExt  string `mapstructure:"LOG_FILE_EXT"`
	TimeFormat  string `mapstructure:"TIME_FORMAT"`

	DefaultCallerDepth int `mapstructure:"DEFAULT_CALLER_DEPTH"`
	DefaultPrefix string `mapstructure:"DEFAULT_PREFIX"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
