package typeconverter

import "reflect"

func ConvertToUserFrinedlyName(val any) string {
	t := reflect.TypeOf(val)

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	default:
		return t.Name()
	}
}