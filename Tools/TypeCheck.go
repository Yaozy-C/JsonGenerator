package Tools

import "log"

func Check(name string, val interface{}) {
	switch val.(type) {

	case []interface{}:
		info, ok := val.([]interface{})
		if !ok {
			panic("conv err:" + name)
		}

		if len(info) < 0 {
			log.Fatal("the list is empty:" + name)
			return
		}

		next, ok := info[0].(map[string]interface{})
		if !ok {
			panic("next conv err")
		}

		for k, v := range next {
			typeCheck(name, k, v)
		}
	case map[string]interface{}:
		for k, v := range val.(map[string]interface{}) {
			typeCheck(name, k, v)
		}
	}
}

func typeCheck(name string, valName string, val interface{}) {
	JsonInfo.InitObjectList(name)
	switch val.(type) {
	case float64:
		JsonInfo.Append(name, valName, "float")
	case string:
		JsonInfo.Append(name, valName, "string")
	case bool:
		JsonInfo.Append(name, valName, "bool")
	case []interface{}:

		info, ok := val.([]interface{})
		if !ok {
			panic("conv err:" + valName)
		}
		if len(info) == 0 {
			log.Println("the list is empty:" + valName)
			return
		}

		next, ok := info[0].(map[string]interface{})
		if !ok {
			switch info[0].(type) {
			case float64:
				JsonInfo.AppendChild(name, valName, "[]interface", "double", false)
			case string:
				JsonInfo.AppendChild(name, valName, "[]interface", "string", false)
			case bool:
				JsonInfo.AppendChild(name, valName, "[]interface", "bool", false)
			case []interface{}:
				JsonInfo.AppendChild(name, valName, "[]interface", "[]interface", true)
				child, ok := info[0].([]interface{})
				if !ok {
					panic("conv err: " + name + "     " + valName)
				}
				typeCheck(name, valName, child)
			}
		} else {
			JsonInfo.AppendChild(name, valName, "[]interface", "interface", false)
			for k, v := range next {
				typeCheck(valName, k, v)
			}
		}

	case map[string]interface{}:
		JsonInfo.Append(name, valName, "interface")
		for k, v := range val.(map[string]interface{}) {
			typeCheck(valName, k, v)
		}
	}
}
