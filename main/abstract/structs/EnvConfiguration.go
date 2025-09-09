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

func (ec *EnvConfiguration) GetInt(key string, defaultValue *int) (*int, error) {
	v, err := ec.Get(key)
	if err != nil {
		if defaultValue == nil {
			return nil, err
		}

		return defaultValue, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

func (ec *EnvConfiguration) GetFloat(key string, defaultValue *float64) (*float64, error) {
	v, err := ec.Get(key)
	if err != nil {
		if defaultValue == nil {
			return nil, err
		}

		return defaultValue, nil
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (ec *EnvConfiguration) GetBool(key string, defaultValue *bool) (*bool, error) {
	v, err := ec.Get(key)
	if err != nil {
		if defaultValue == nil {
			return nil, err
		}

		return defaultValue, nil
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
