package config

import "os"

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

//ENVIRONMENT
const ENVIRONMENT string = LOCAL

// MAP LIST ENVIRONMENT LOCAL AND DEVELOPMENT
var env = map[string]map[string]string{
	"local": {
		"MS_PORT":         "11000",
		"MYSQL_HOST":   "localhost",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "",
		"MYSQL_SCHEMA": "majoo",
		"SECRET_KEY": "MAJOO",
		"APP_NAME": "MAJOO",
	},
	"development": {
		"MS_PORT":      "",
		"MYSQL_HOST":   "",
		"MYSQL_PORT":   "",
		"MYSQL_USER":   "",
		"MYSQL_PASS":   "",
		"MYSQL_SCHEMA": "",
		"SECRET_KEY": "MAJOO",
		"APP_NAME": "MAJOO",
	},
}

var CONFIG = env[ENVIRONMENT]

// CHECK ENVIRONMENT AND GET ENVIRONMENT
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// INITIALIZE CONFIGURATION
func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}