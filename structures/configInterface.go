package structures

type Config interface {
	ToMap() map[string]interface{}
}
