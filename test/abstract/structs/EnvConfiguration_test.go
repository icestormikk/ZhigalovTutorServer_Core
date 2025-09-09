package structs

import (
	"testing"
	"zhigalov_tutor_server_core/main/abstract/structs"

	"github.com/stretchr/testify/assert"
)

type EnvConfigurationTestContext struct {
	config *structs.EnvConfiguration
}

func (ec *EnvConfigurationTestContext) beforeEach() {
	ec.config = structs.NewEnvConfiguration("./.env.test")
}
func (ec *EnvConfigurationTestContext) afterEach() {}
func testCase(test func(t *testing.T, ctx *EnvConfigurationTestContext)) func(t *testing.T) {
	return func(t *testing.T) {
		context := &EnvConfigurationTestContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

func TestEnvConfigurationNoFiles(t *testing.T) {
	assert.Panics(t, func() { structs.NewEnvConfiguration() })
}

func TestEnvConfigurationCorrectFile(t *testing.T) {
	assert.NotPanics(t, func() { structs.NewEnvConfiguration("./.env.test") })
}

func TestEnvConfigurationGet(t *testing.T) {
	t.Run("Attempt to get any variable", testCase(func(t *testing.T, ctx *EnvConfigurationTestContext) {
		_, err := ctx.config.Get("TEST_VAR")
		assert.Nil(t, err)
	}))
}

func TestEnvConfigurationGetNotExists(t *testing.T) {
	t.Run("Attempt to get a non-existing variable", testCase(func(t *testing.T, ctx *EnvConfigurationTestContext) {
		_, err := ctx.config.Get("TEST_NOEXIST_VAR")
		assert.NotNil(t, err)
	}))
}

func TestEnvConfigurationGetInt(t *testing.T) {
	tests := []struct {
		name                string
		varName             string
		defaultValue        any
		expectedValue       any
		isErrorMustBeThrown bool
	}{
		{
			name:                "Attempt to get an integer (exists, correct, no default)",
			varName:             "TEST_INT_VAR",
			defaultValue:        nil,
			expectedValue:       1,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get an integer (exists, incorrect, no default)",
			varName:             "TEST_INCORRECT_INT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (exists, correct, with default)",
			varName:             "TEST_INT_VAR",
			defaultValue:        0,
			expectedValue:       1,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get an integer (exists, incorrect, with default)",
			varName:             "TEST_INCORRECT_INT_VAR",
			defaultValue:        0,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (no exists, correct, no default)",
			varName:             "TEST_NOEXIST_INT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (no exists, correct, with default)",
			varName:             "TEST_NOEXIST_INT_VAR",
			defaultValue:        0,
			expectedValue:       0,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get an integer (no exists, incorrect, no default)",
			varName:             "TEST_NOEXIST_INCORRECT_INT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (no exists, incorrect, with default)",
			varName:             "TEST_NOEXIST_INCORRECT_INT_VAR",
			defaultValue:        0,
			expectedValue:       0,
			isErrorMustBeThrown: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, testCase(func(t *testing.T, ctx *EnvConfigurationTestContext) {
			var defaultValue *int = nil
			if test.defaultValue != nil {
				conv := test.defaultValue.(int)
				defaultValue = &conv
			}

			i, err := ctx.config.GetInt(test.varName, defaultValue)
			if test.isErrorMustBeThrown {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expectedValue, *i)
		}))
	}
}

func TestEnvConfigurationGetFloat(t *testing.T) {
	tests := []struct {
		name                string
		varName             string
		defaultValue        any
		expectedValue       any
		isErrorMustBeThrown bool
	}{
		{
			name:                "Attempt to get a float (exists, correct, no default)",
			varName:             "TEST_FLOAT_VAR",
			defaultValue:        nil,
			expectedValue:       1.0,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a float (exists, incorrect, no default)",
			varName:             "TEST_INCORRECT_FLOAT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a float (exists, correct, with default)",
			varName:             "TEST_FLOAT_VAR",
			defaultValue:        0.0,
			expectedValue:       1.0,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a float (exists, incorrect, with default)",
			varName:             "TEST_INCORRECT_FLOAT_VAR",
			defaultValue:        0.0,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a float (no exists, correct, no default)",
			varName:             "TEST_NOEXIST_FLOAT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a float (no exists, correct, with default)",
			varName:             "TEST_NOEXIST_FLOAT_VAR",
			defaultValue:        0.0,
			expectedValue:       0.0,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a float (no exists, incorrect, no default)",
			varName:             "TEST_NOEXIST_INCORRECT_FLOAT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (no exists, incorrect, with default)",
			varName:             "TEST_NOEXIST_INCORRECT_FLOAT_VAR",
			defaultValue:        0.0,
			expectedValue:       0.0,
			isErrorMustBeThrown: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, testCase(func(t *testing.T, ctx *EnvConfigurationTestContext) {
			var defaultValue *float64 = nil
			if test.defaultValue != nil {
				conv := test.defaultValue.(float64)
				defaultValue = &conv
			}

			i, err := ctx.config.GetFloat(test.varName, defaultValue)
			if test.isErrorMustBeThrown {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expectedValue, *i)
		}))
	}
}

func TestEnvConfigurationGetBool(t *testing.T) {
	tests := []struct {
		name                string
		varName             string
		defaultValue        any
		expectedValue       any
		isErrorMustBeThrown bool
	}{
		{
			name:                "Attempt to get a bool (exists, correct, no default)",
			varName:             "TEST_BOOL_VAR",
			defaultValue:        nil,
			expectedValue:       true,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a bool (exists, incorrect, no default)",
			varName:             "TEST_INCORRECT_FLOAT_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a float (exists, correct, with default)",
			varName:             "TEST_BOOL_VAR",
			defaultValue:        false,
			expectedValue:       true,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a bool (exists, incorrect, with default)",
			varName:             "TEST_INCORRECT_BOOL_VAR",
			defaultValue:        false,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a bool (no exists, correct, no default)",
			varName:             "TEST_NOEXIST_BOOL_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get a bool (no exists, correct, with default)",
			varName:             "TEST_NOEXIST_BOOL_VAR",
			defaultValue:        false,
			expectedValue:       false,
			isErrorMustBeThrown: false,
		},
		{
			name:                "Attempt to get a bool (no exists, incorrect, no default)",
			varName:             "TEST_NOEXIST_INCORRECT_BOOL_VAR",
			defaultValue:        nil,
			expectedValue:       nil,
			isErrorMustBeThrown: true,
		},
		{
			name:                "Attempt to get an integer (no exists, incorrect, with default)",
			varName:             "TEST_NOEXIST_INCORRECT_BOOL_VAR",
			defaultValue:        false,
			expectedValue:       false,
			isErrorMustBeThrown: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, testCase(func(t *testing.T, ctx *EnvConfigurationTestContext) {
			var defaultValue *bool = nil
			if test.defaultValue != nil {
				conv := test.defaultValue.(bool)
				defaultValue = &conv
			}

			i, err := ctx.config.GetBool(test.varName, defaultValue)
			if test.isErrorMustBeThrown {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expectedValue, *i)
		}))
	}
}
