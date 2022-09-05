package processor

import (
	"reflect"
)

func AddPayloadNumbers(payload interface{}) (total int) {
	payloadType := reflect.TypeOf(payload)

	switch payloadType.Kind() {
	case reflect.Map:
		for _, j := range payload.(map[string]interface{}) {
			total += AddPayloadNumbers(j)
		}
	case reflect.Array, reflect.Slice:
		for _, i := range payload.([]interface{}) {
			total += AddPayloadNumbers(i)
		}
	case reflect.Float64:
		total += int(payload.(float64))
	case reflect.Int:
		total += payload.(int)
	}
	return total
}
