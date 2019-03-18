package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var envValue map[string]string

type config interface {
	initSettingValue()

	init(engine *gin.Engine)
}

var configs = []config{
	WebConfig{},
}

func Bootstrap(engine *gin.Engine) (error error) {
	e := readEnv(&envValue)
	if e != nil {
		return e
	}
	for _, v := range configs {
		v.initSettingValue()
		v.init(engine)
	}
	return nil
}

func readEnv(env *map[string]string) (error error) {
	tempEnv := map[string]string{}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	envPath := dir + filepath.FromSlash("/.env")
	_, e := os.Stat(envPath)
	if e != nil {
		return e
	}
	file, _ := os.Open(envPath)
	buf := bufio.NewReader(file)
	rows := 0
	for {
		rows++
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return err
			}
		}
		if line == "" {
			continue
		}
		split := strings.Split(line, "=")
		if len(split) == 2 {
			tempEnv[split[0]] = split[1]
		} else {
			fmt.Printf("error env in line %d", rows)
		}
	}
	env = &tempEnv
	return nil
}

func Env(key string, defaultValue string) string {
	value := envValue[key]
	if value != "" {
		return value
	}
	return defaultValue
}
