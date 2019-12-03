package allure

import "log"

func getFromCtx(key string) interface{} {
	value, ok := ctxMgr.GetValue(key)
	if !ok {
		log.Fatalf("could not extract object by key \"%s\"\n", key)
	}

	return value
}

func manipulateOnObjectFromCtx(key string, action func(object interface{})) {
	if object, ok := ctxMgr.GetValue(key); ok {
		action(object)
	} else {
		log.Fatalf("could not extract object by key \"%s\"\n", key)
	}
}
