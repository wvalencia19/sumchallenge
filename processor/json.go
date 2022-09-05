package processor

import "encoding/json"

func UnmarshallToInterface(payload []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(payload), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
