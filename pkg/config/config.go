package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string //must BindEnv
	PocketConsumerKey string //must BindEnv
	AuthServerUrl     string //must BindEnv
	TelegramBotUrl    string `mapstructure:"bot_url"`
	DbPath            string `mapstructure:"db_file"`
	Message           Message
}

type Message struct {
	Errors   Errors
	Response Response
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	UnableToSave string `mapstructure:"unable_to_save"`
	Unauthorized string `mapstructure:"unauthorized"`
}

type Response struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("main")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		//panic(fmt.Errorf("fatal error config file: %w", err))
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("message.response", &cfg.Message.Response); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("message.errors", &cfg.Message.Errors); err != nil {
		return nil, err
	}

	if err := parserEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parserEnv(cfg *Config) error {

	//os.Setenv("TOKEN", "******")

	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_redirect_url"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	cfg.PocketConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerUrl = viper.GetString("auth_server_redirect_url")
	return nil
}
