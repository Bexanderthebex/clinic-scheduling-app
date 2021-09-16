package elasticsearch

type ElasticSearchResponse struct {
	Response map[string]interface{}
}

func (esr ElasticSearchResponse) TotalHits() int64 {
	return int64(esr.Response["hits"].(map[string]interface{})["total"].(float64))
}

func (esr ElasticSearchResponse) First() interface{} {
	return esr.Response["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"]
}

func (esr ElasticSearchResponse) Paginate(pageSize int) []interface{} {
	results := esr.Response["hits"].(map[string]interface{})["hits"].([]interface{})

	maxLength := pageSize
	if len(results) < pageSize {
		maxLength = len(results)
	}

	response := make([]interface{}, 0)
	for i := 0; i < maxLength; i++ {
		response = append(response, esr.Response["hits"].(map[string]interface{})["hits"].([]interface{})[i].(map[string]interface{})["_source"])
	}

	return response
}
