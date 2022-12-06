package config

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Printf("%s key is not set", key)
		panic(fmt.Sprintf("%s key is not set", key))
	}
}

func getStringArray(key string) []string {
	panicIfNotPresent(key)
	strValues := strings.Split(viper.GetString(key), ",")
	for i, str := range strValues {
		strValues[i] = strings.TrimSpace(str)
	}
	return strValues
}

func panicIfNotPresent(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func getIntOrDefault(key string, defaultValue int) int {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return getIntValue(key)
}

func getIntValue(key string) int {
	panicIfNotPresent(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}
