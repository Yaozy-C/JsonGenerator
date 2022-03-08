package Tools

import (
	"log"
	"os"
	"strings"
)

func writeToFile(name string, msg string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
	err = f.Close()
	if err != nil {
		log.Printf("close file err:%s", err)
		return
	}
}

func replaceMax(name string) string {
	if len(name) == 0 {
		panic("name is empty")
	}
	if name[0] >= 'a' {
		return strings.Replace(name, string(name[0]), string(name[0]-32), 1)
	}
	return name
}

func replaceMin(name string) string {
	if len(name) == 0 {
		panic("name is empty")
	}
	if name[0] < 'a' {
		return strings.Replace(name, string(name[0]), string(name[0]+32), 1)
	}
	return name
}

func replaceAll(name string) string {
	if len(name) == 0 {
		panic("name is empty")
	}

	str := ""
	for _, i2 := range name {
		if i2 >= 'a' {
			str += string(i2 - 32)
		} else {
			str += string(i2)
		}
	}

	return str
}

func GenerateHead(name string, tType string, childType string) string {

	switch tType {
	case "string":
		return "#include <string>\n"
	case "[]interface":
		head := "#include <vector>\n"
		if childType == "[]interface" || childType == "interface" {
			head += "#include \"" + replaceMax(name) + ".h\"\n"
		}
		return head
	case "interface":
		return "#include \"" + replaceMax(name) + ".h\"\n"
	}

	return ""
}

func GenerateObject(name string) string {
	objectName := replaceMax(name)

	objectName += " : public Base::DataPacket "
	return objectName
}

func GenerateProperty(name string, tType string, childType string, count int) string {

	switch tType {
	case "string":
		return "	std::string " + replaceMin(name) + ";\n\n"
	case "bool":
		return "	bool " + replaceMin(name) + ";\n\n"
	case "float":
		return "	double " + replaceMin(name) + ";\n\n"
	case "[]interface":
		info := ""
		if childType == "[]interface" || childType == "interface" {
			info = "std::vector<" + replaceMax(name) + "> "
		} else {

			info = "std::vector<" + childType + "> "
		}

		for i := 1; i < count; i++ {
			info = "std::vector<" + info + ">"
		}

		return "	" + info + replaceMin(name) + ";\n\n"
	case "interface":
		return "	" + replaceMax(name) + " " + replaceMin(name) + ";\n\n"
	}

	return ""
}

func GenerateFunction(name string) (string, string) {
	decode := "	Base::GetJsonValue(reader, \"" + name + "\", " + replaceMin(name) + ");\n"
	encode := "	Base::AddJsonValue(writer, \"" + name + "\", " + replaceMin(name) + ");\n"

	return decode, encode
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Println(err)
		return false
	}
	return true
}

func Generate(name string, infoList []*Object) {

	if !isExist("./code") {
		err := os.MkdirAll("./code", 0777)
		if err != nil {
			panic(err)
		}
	}

	fileName := "./code/" + replaceMax(name) + ".h"

	head := "#ifndef " + replaceAll(name) + "_H\n#define " + replaceAll(name) + "_H\n\n#include \"DataPacket.h\"\n"

	property := "public:\n\n"
	object := "class " + GenerateObject(name) + "{\n"
	encode := "void " + replaceMax(name) + "::EncodeJson(cJSON *writer) {\n"
	decode := "void " + replaceMax(name) + "::DecodeJson(cJSON *reader) {\n"
	for _, info := range infoList {
		res := GenerateHead(info.name, info.tType, info.childType)
		if !strings.Contains(head, res) {
			head += GenerateHead(info.name, info.tType, info.childType)
		}

		property += GenerateProperty(info.name, info.tType, info.childType, info.count)
		d, e := GenerateFunction(info.name)
		encode += e
		decode += d
	}

	encode += "}\n\n"
	decode += "}\n\n"

	object += property
	object += "	void DecodeJson(cJSON *reader) override;\n\n	void EncodeJson(cJSON *writer) override;\n};\n\n"

	code := head + "\n" + object

	code += encode
	code += decode

	code += "#endif //" + replaceAll(name) + "_H"
	log.Println("\n" + fileName + "\n" + code)

	writeToFile(fileName, code)
	return
}
