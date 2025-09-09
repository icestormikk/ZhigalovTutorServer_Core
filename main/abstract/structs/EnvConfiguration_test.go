package structs

import (
	"os"
	"reflect"
	"testing"
)

var values = &map[string]string{
	"INT_VAR":             "1",
	"INCORRECT_INT_VAR":   "1d",
	"FLOAT_VAR":           "1.0",
	"INCORRECT_FLOAT_VAR": "1.0f",
	"BOOL_VAR":            "true",
	"INCORRECT_BOOL_VAR":  "truth",
}

var notExistentVarName = "VAR"

func TestEnvConfiguration_Get(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Attempt to get a variable (exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INT_VAR"},
			want:    "1",
			wantErr: false,
		},
		{
			name:    "Attempt to get a variable (exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_INT_VAR"},
			want:    "1d",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &EnvConfiguration{
				values: tt.fields.values,
			}
			got, err := ec.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvConfiguration_GetBool(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key          string
		defaultValue *bool
	}

	defaultValue := false

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "Attempt to get a bool (exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: "BOOL_VAR"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Attempt to get a bool (exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_BOOL_VAR"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a bool (exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: "BOOL_VAR", defaultValue: &defaultValue},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Attempt to get a bool (exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_BOOL_VAR", defaultValue: &defaultValue},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a bool (no exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a bool (no exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Attempt to get a bool (no exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a bool (no exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &EnvConfiguration{
				values: tt.fields.values,
			}
			got, err := ec.GetBool(tt.args.key, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("GetBool() got = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestEnvConfiguration_GetFloat(t *testing.T) {
	type fields struct {
		values map[string]string
	}

	type args struct {
		key          string
		defaultValue *float64
	}

	defaultValue := 0.0

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "Attempt to get a float (exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: "FLOAT_VAR"},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "Attempt to get a float (exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_FLOAT_VAR"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a float (exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: "FLOAT_VAR", defaultValue: &defaultValue},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "Attempt to get a float (exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_FLOAT_VAR", defaultValue: &defaultValue},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a float (no exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a float (no exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "Attempt to get a float (no exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get a float (no exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    0.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &EnvConfiguration{
				values: tt.fields.values,
			}
			got, err := ec.GetFloat(tt.args.key, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("GetFloat() got = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestEnvConfiguration_GetInt(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key          string
		defaultValue *int
	}

	defaultValue := 0

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "Attempt to get an int (exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INT_VAR"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Attempt to get an int (exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_INT_VAR"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get an int (exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: "INT_VAR", defaultValue: &defaultValue},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Attempt to get an int (exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: "INCORRECT_INT_VAR", defaultValue: &defaultValue},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get an int (no exists, correct, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get an int (no exists, correct, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    0,
			wantErr: false,
		},
		{
			name:    "Attempt to get an int (no exists, incorrect, no default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Attempt to get an int (no exists, incorrect, with default)",
			fields:  fields{values: *values},
			args:    args{key: notExistentVarName, defaultValue: &defaultValue},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &EnvConfiguration{
				values: tt.fields.values,
			}
			got, err := ec.GetInt(tt.args.key, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("GetInt() got = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestNewEnvConfiguration(t *testing.T) {
	type args struct {
		envFiles []string
	}

	envFileName := ".env.test"
	defaultEnvFile, err := os.CreateTemp(".", envFileName)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(defaultEnvFile *os.File) {
		err := defaultEnvFile.Close()
		if err != nil {
			t.Errorf("Error closing temp file")
		}
		err = os.Remove(defaultEnvFile.Name())
		if err != nil {
			t.Errorf("Error removing temp file")
		}
	}(defaultEnvFile)

	_, err = defaultEnvFile.WriteString("INT_VAR=1\nFLOAT_VAR=2\n")
	if err != nil {
		t.Errorf("Error while writing to temp file: %v", err)
	}

	tests := []struct {
		name    string
		args    args
		want    *EnvConfiguration
		wantErr bool
	}{
		{
			name:    "EnvConfiguration with empty file list",
			args:    args{envFiles: []string{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "EnvConfiguration with one file",
			args:    args{envFiles: []string{defaultEnvFile.Name()}},
			want:    &EnvConfiguration{values: map[string]string{"INT_VAR": "1", "FLOAT_VAR": "2"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); tt.wantErr && r == nil {
					t.Errorf("NewEnvConfiguration() error expected, but not received")
				}
			}()

			if got := NewEnvConfiguration(tt.args.envFiles...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
