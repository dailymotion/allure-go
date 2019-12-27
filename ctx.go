package allure

import "log"

func getFromCtx(key string) interface{} {
	value, ok := ctxMgr.GetValue(key)
	if !ok {
		log.Printf("could not extract object by key \"%s\"\n", key)
	}

	return value
}

func manipulateOnObjectFromCtx(key string, action func(object interface{})) {
	if object, ok := ctxMgr.GetValue(key); ok {
		action(object)
	} else {
		log.Printf("could not extract object by key \"%s\"\n", key)
	}
}
