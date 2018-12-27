package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var EnvValue map[string]string

type appConfig interface {
	initConfig() (error error)
}

var configs = []appConfig{

}

func Bootstrap() (error error) {
	env, e := readEnv()
	if e != nil {
		return e
	}
	EnvValue = env
	initSettingValue()
	for _, v := range configs {
		initConfigError := v.initConfig()
		if initConfigError != nil {
			return initConfigError
		}
	}
	return nil
}

func readEnv() (env map[string]string, error error) {
	env = map[string]string{}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	envPath := dir + filepath.FromSlash("/.env")
	_, e := os.Stat(envPath)
	if e != nil {
		return nil, e
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
				return nil, err
			}
		}
		if line == "" {
			continue
		}
		split := strings.Split(line, "=")
		if len(split) == 2 {
			env[split[0]] = split[1]
		} else {
			fmt.Printf("error env in line %d", rows)
		}
	}
	return env, nil
}

func Env(key string, defaultValue ...string) string {
	value := EnvValue[key]
	if value != "" {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}
