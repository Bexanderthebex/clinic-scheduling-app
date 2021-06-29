package elasticsearch

type ElasticSearchResponse struct {
	Response map[string]interface{}
}

func (esr ElasticSearchResponse) TotalHits() int64 {
	return int64(esr.Response["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
}

func (esr ElasticSearchResponse) First() interface{} {
	return esr.Response["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"]
}
