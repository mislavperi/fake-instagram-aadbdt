package utils

import "errors"

func CapturePanic(fatal bool, context map[string]interface{}) {
	if r := recover(); r != nil {
		var err error

		switch recType := r.(type) {
		case string:
			err = errors.New(recType)
		case error:
			err = recType
		default:
			err = errors.New("unknown panic")
		}

		if context == nil {
			context = map[string]interface{}{}
		}

		context["error"] = err
		if fatal {
			panic(err)
		}
	}
}
