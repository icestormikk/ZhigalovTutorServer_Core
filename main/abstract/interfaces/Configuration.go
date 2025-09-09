package interfaces

type Configuration interface {
	Get(key string) (string, error)
	GetInt(key string, defaultValue *int) (*int, error)
	GetFloat(key string, defaultValue *float64) (*float64, error)
	GetBool(key string, defaultValue *bool) (*bool, error)
}
