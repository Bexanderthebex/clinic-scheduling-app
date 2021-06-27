package repository

// For Document cache
type DocumentCache interface {
	Find(map[string]interface{}) map[string]interface{}
}

// For Key value type cache
type KeyValueCache interface {
	Get(string) string
}
