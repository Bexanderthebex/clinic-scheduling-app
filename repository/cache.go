package repository

// For Document cache
type DocumentCache interface {
	Find(map[string]interface{}) map[string]interface{}
	IndexExists(string) (bool, error)
	CreateIndex(string) (map[string]interface{}, error)
}

// For Key value type cache
type KeyValueCache interface {
	Get(string) string
}
