package structs

import (
	"errors"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfiguration struct {
	values map[string]string
}

func NewEnvConfiguration(envFiles ...string) *EnvConfiguration {
	var values map[string]string

	values, err := godotenv.Read(envFiles...)
	if err != nil {
		panic(err)
	}

	return &EnvConfiguration{values: values}
}

func (ec *EnvConfiguration) Get(key string) (string, error) {
	value, ok := ec.values[key]
	if !ok {
		return "", errors.New(key + " env variable is missing")
	}

	return value, nil
}

func (ec *EnvConfiguration) GetInt(key string) (int, error) {
	v, err := ec.Get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (ec *EnvConfiguration) GetFloat(key string) (float64, error) {
	v, err := ec.Get(key)
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (ec *EnvConfiguration) GetBoolean(key string) (bool, error) {
	v, err := ec.Get(key)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}

	return b, nil
}
