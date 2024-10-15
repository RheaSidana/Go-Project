package initializer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filePath = "./.env"

func LoadEnvVariables() error {
	file, err := openEnvFile()
	if err != nil {
		return err
	}
	defer file.Close()

	envMap, err := readEnvFile(file)
	if err != nil {
		return err
	}

	err = setEnvVariables(envMap)
	if err != nil {
		return err
	}

	return nil
}

func setEnvVariables(envMap map[string]string) error {
	for key, value := range envMap {
		err := setEnvVariable(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func splitKeyValue(line string) (key, value string, err error) {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("error parsing line: %s", line)
	}

	key = strings.TrimSpace(parts[0])
	value = strings.TrimSpace(parts[1])
	return
}

func setEnvVariable(key, value string) error {
	if err := os.Setenv(key, value); err != nil {
		return fmt.Errorf("error setting env variable: %v", err)
	}
	return nil
}

func readEnvFile(file *os.File) (map[string]string, error) {
	scanner := bufio.NewScanner(file)
	envMap := make(map[string]string)
	for scanner.Scan() {
		// fmt.Println("File reading!")
		line := scanner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, err := splitKeyValue(line)
		if err != nil {
			return nil, err
		}
		// fmt.Println("key and value extracted successfully!")

		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %v", err)
	}
	return envMap, nil
}

func openEnvFile() (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening .env file: %v", err)
	}

	return file, nil
}
