package tool

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// GetStringWithDefault get string value from viper, env, and fallback to default value assigned
func GetStringWithDefault(viperkey string, env string, defaultVal string) string {
	if value := viper.GetString(viperkey); value != "" {
		return value
	}

	if value := os.Getenv(env); value != "" {
		return value
	}

	return defaultVal
}

// GetIntWithDefault get int value from viper, env, and fallback to default value assigned
func GetIntWithDefault(viperkey string, env string, defaultVal int) int {
	if value := viper.Get(viperkey); value != nil {
		switch value.(type) {
		case string:
			if v, err := strconv.Atoi(value.(string)); err == nil {
				return v
			}
		case int:
			return value.(int)
		}
	}

	if value := os.Getenv(env); value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultVal
}

// GetBoolWithDefault get bool value from viper, env, and fallback to default value assigned
func GetBoolWithDefault(viperkey string, env string, defaultVal bool) bool {
	if value := viper.GetString(viperkey); value != "" {
		return viper.GetBool(viperkey)
	}

	value := os.Getenv(env)
	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}

	return boolVal
}
