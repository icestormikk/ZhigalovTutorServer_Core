package interfaces

type Configuration interface {
	Get(key string) (string, error)
	GetInt(key string) (int, error)
	GetFloat(key string) (float64, error)
	GetBoolean(key string) (bool, error)
}
