package JsonTest

import (
	"JsonGenerator/Tools"
	"encoding/json"
	"log"
	"testing"
)

func TestListNestList(t *testing.T) {
	jsonStr := "{\n  \"Age\": 20,\n  \"Test\": [\n    [\n      [\n        11\n      ],\n      [\n        12\n      ]\n    ],\n    [\n      [\n        21\n      ],\n      [\n        22\n      ]\n    ],\n    [\n      [\n        31\n      ],\n      [\n        32\n      ]\n    ],\n    [\n      [\n        41\n      ],\n      [\n        42\n      ]\n    ],\n    [\n      [\n        51\n      ],\n      [\n        52\n      ]\n    ],\n    [\n      [\n        61\n      ],\n      [\n        62\n      ]\n    ]\n  ],\n  \"Score\": {\n    \"Name\": \"math\"\n  },\n  \"Name\": \"test1\",\n  \"Status\": true,\n  \"PhoneList\": [\n    {\n      \"Name\": \"apple\",\n      \"Number\": \"123213123213\"\n    },\n    {\n      \"Name\": \"hw\",\n      \"Number\": \"4564645645\"\n    }\n  ]\n}"
	var info map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &info)

	if err != nil {
		log.Fatalf("Unmarshal err: %s", err)
		return
	}

	Tools.JsonInfo.Init()
	Tools.Check("Info", info)

	log.Println("----------------------------------")

	Tools.JsonInfo.Gen()
}

func TestNormalJson(t *testing.T) {
	jsonStr := "{\"StudentID\":\"100\",\"Name\":\"tmc\",\"Hometown\":\"usa\"}"
	var info map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &info)

	if err != nil {
		log.Fatalf("Unmarshal err: %s", err)
		return
	}

	Tools.JsonInfo.Init()
	Tools.Check("Student", info)

	log.Println("----------------------------------")

	Tools.JsonInfo.Gen()
}

func TestListJson(t *testing.T) {
	jsonStr := "[{\n\t\t\"StudentID\": \"100\",\n\t\t\"Name\": \"aaa\",\n\t\t\"Hometown\": \"china\"\n\t},\n\t{\n\t\t\"StudentID\": \"101\",\n\t\t\"Name\": \"bbb\",\n\t\t\"Hometown\": \"us\"\n\t},\n\t{\n\t\t\"StudentID\": \"102\",\n\t\t\"Name\": \"ccc\",\n\t\t\"Hometown\": \"england\"\n\t}\n]"
	var info []interface{}

	err := json.Unmarshal([]byte(jsonStr), &info)

	if err != nil {
		log.Fatalf("Unmarshal err: %s", err)
		return
	}

	Tools.JsonInfo.Init()
	Tools.Check("Student", info)

	log.Println("----------------------------------")

	Tools.JsonInfo.Gen()
}

func TestDifficultyJson(t *testing.T) {
	jsonStr := "{\n  \"Data\": {\n    \"article\": {\n      \"title\": \"??????\",\n      \"searchType\": 1,\n      \"data\": [{\n        \"article_id\": 10496,\n        \"title\": \"??????????????????baby????????????\",\n        \"content\": \"\",\n        \"author\": \"??????????????????\",\n        \"business_brand_name\": null,\n        \"reads\": 0,\n        \"is_top\": 0,\n        \"is_hot\": 0,\n        \"is_discuss\": 0,\n        \"is_red\": 0,\n        \"user_id\": 1020,\n        \"user_name\": \"??????\",\n        \"sort_no\": 99,\n        \"articletype_id\": 0,\n        \"articletype\": {\n          \"articletype_id\": 1002,\n          \"typename\": \"????????????\",\n          \"fid\": 0\n        },\n        \"image_id\": null,\n        \"image\": {\n          \"file_id\": 15648,\n          \"path\": \"http://xxxx.jpg\",\n          \"grade_code\": \"7f631ed8-0c80-4e91-8b4a-c\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"grade_code\": \"\",\n        \"imageList\": [],\n        \"keydes\": \"\",\n        \"photo\": null,\n        \"iscollect\": 0,\n        \"utime\": \"2018-09-19T16:00:08.447\",\n        \"likecount\": 0,\n        \"commentlist\": null,\n        \"statusName\": \"auditSuccess\",\n        \"contentType\": 1\n      }]\n    },\n    \"video\": {\n      \"title\": \"??????\",\n      \"searchType\": 2,\n      \"data\": [{\n        \"courses_id\": 1059,\n        \"title\": \"??????????????????\",\n        \"content\": \"????????????????????????0-1???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????\",\n        \"price\": 0,\n        \"author\": \"????????????\",\n        \"purchaseNotes\": \"???????????????????????????????????????\",\n        \"is_top\": 0,\n        \"is_hot\": 1,\n        \"is_red\": 0,\n        \"is_discuss\": 0,\n        \"is_concentration\": 1,\n        \"user\": {\n          \"userdtl\": null,\n          \"coachUser\": null,\n          \"businessUser\": null,\n          \"user_id\": 1146,\n          \"phone\": \"13648953657\",\n          \"name\": \"????????????\",\n          \"token\": \"3C76F191825EF5302BBC10609A903C77\",\n          \"is_enable\": 1,\n          \"ctime\": \"2018-09-03T20:04:54.153\",\n          \"utime\": \"2018-09-03T20:04:54.153\",\n          \"grade\": {\n            \"grade_id\": 1008,\n            \"grade_name\": \"??????????????????\",\n            \"sort_no\": 99,\n            \"ctime\": \"2018-07-04T10:59:22.067\",\n            \"utime\": \"2018-07-04T10:59:22.067\",\n            \"description\": null,\n            \"userRole\": [\n              \"coach-ordinary\"\n            ],\n            \"userType\": \"coachUser\",\n            \"alias\": \"coach-ordinary\",\n            \"perssions\": null,\n            \"isUsed\": false,\n            \"userCount\": 0\n          },\n          \"userType\": \"coachUser\",\n          \"isFollow\": 0,\n          \"image\": {\n            \"file_id\": 13293,\n            \"path\": \"http://xxxx.jpg\",\n            \"grade_code\": \"\",\n            \"sort_no\": 99,\n            \"filetype_name\": \"pictext\"\n          },\n          \"resume\": \"????????????????????????????????????????????????????????????\",\n          \"nikeName\": \"????????????\",\n          \"label\": \"??????????????????\"\n        },\n        \"image\": {\n          \"file_id\": 14899,\n          \"path\": \"http://xxxxx.png\",\n          \"grade_code\": \"6bddbf33-1c75-4fca-8195-9f41540fc8b7\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"imageList\": [],\n        \"sort_no\": 99,\n        \"ctime\": \"2018-09-10T16:34:34.983\",\n        \"utime\": \"2018-09-10T16:34:34.997\",\n        \"keydes\": \"\",\n        \"article_source\": \"??????\",\n        \"coursesType\": 2,\n        \"classify_id\": 0,\n        \"classify\": {\n          \"classify_id\": 1005,\n          \"fid\": 1002,\n          \"classifys\": null,\n          \"classifyName\": \"????????????\",\n          \"ctime\": \"2018-06-29T10:42:34\",\n          \"sort_no\": 1,\n          \"classifyType\": 2,\n          \"isUsed\": false\n        },\n        \"practicalPeople\": \"0-1???\",\n        \"totalNumber\": 0,\n        \"studyNumber\": 3,\n        \"vip_label\": 0,\n        \"integral\": 0,\n        \"discount\": 1,\n        \"consumptionType\": \"free\",\n        \"consumptionDetails\": [],\n        \"score\": 2.5,\n        \"isPurchase\": 0,\n        \"isScore\": 0,\n        \"comments\": null,\n        \"progressShow\": null,\n        \"isCollect\": 0,\n        \"statusName\": \"1\",\n        \"fcModuleType\": 0\n      }]\n    },\n    \"audio\": {\n      \"title\": \"??????\",\n      \"searchType\": 3,\n      \"data\": [{\n        \"courses_id\": 1068,\n        \"title\": \"????????????\",\n        \"content\": \"????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????\",\n        \"price\": 0,\n        \"author\": \"?????????????????????\",\n        \"purchaseNotes\": \"???????????????????????????????????????\",\n        \"is_top\": 0,\n        \"is_hot\": 1,\n        \"is_red\": 0,\n        \"is_discuss\": 0,\n        \"is_concentration\": 1,\n        \"user\": {\n          \"userdtl\": null,\n          \"coachUser\": null,\n          \"businessUser\": null,\n          \"user_id\": 1150,\n          \"phone\": \"13523478965\",\n          \"name\": \"?????????????????????\",\n          \"token\": \"4F97A2CABFE986B370066730142653CE\",\n          \"is_enable\": 1,\n          \"ctime\": \"2018-09-03T20:11:53.227\",\n          \"utime\": \"2018-09-03T20:11:53.227\",\n          \"grade\": {\n            \"grade_id\": 1008,\n            \"grade_name\": \"??????????????????\",\n            \"sort_no\": 99,\n            \"ctime\": \"2018-07-04T10:59:22.067\",\n            \"utime\": \"2018-07-04T10:59:22.067\",\n            \"description\": null,\n            \"userRole\": [\n              \"coach-ordinary\"\n            ],\n            \"userType\": \"coachUser\",\n            \"alias\": \"coach-ordinary\",\n            \"perssions\": null,\n            \"isUsed\": false,\n            \"userCount\": 0\n          },\n          \"userType\": \"coachUser\",\n          \"isFollow\": 0,\n          \"image\": {\n            \"file_id\": 13297,\n            \"path\": \"http://xxxxxx.jpg\",\n            \"grade_code\": \"\",\n            \"sort_no\": 99,\n            \"filetype_name\": \"pictext\"\n          },\n          \"resume\": \"????????????????????????????????????????????????????????????\",\n          \"nikeName\": \"?????????????????????\",\n          \"label\": \"????????????\"\n        },\n        \"image\": {\n          \"file_id\": 8830,\n          \"path\": \"http://xxxx.jpg\",\n          \"grade_code\": \"\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"imageList\": [],\n        \"sort_no\": 99,\n        \"ctime\": \"2018-08-21T23:39:20.967\",\n        \"utime\": \"2018-08-21T23:39:20.977\",\n        \"keydes\": \"\",\n        \"article_source\": \"??????\",\n        \"coursesType\": 3,\n        \"classify_id\": 0,\n        \"classify\": {\n          \"classify_id\": 1016,\n          \"fid\": 1000,\n          \"classifys\": null,\n          \"classifyName\": \"??????\",\n          \"ctime\": \"2018-08-07T00:00:00\",\n          \"sort_no\": 2,\n          \"classifyType\": 3,\n          \"isUsed\": false\n        },\n        \"practicalPeople\": \"0-6???\",\n        \"totalNumber\": 0,\n        \"studyNumber\": 1,\n        \"vip_label\": 0,\n        \"integral\": 0,\n        \"discount\": 1,\n        \"consumptionType\": \"free\",\n        \"consumptionDetails\": [],\n        \"score\": 0,\n        \"isPurchase\": 0,\n        \"isScore\": 0,\n        \"comments\": null,\n        \"progressShow\": null,\n        \"isCollect\": 0,\n        \"statusName\": \"1\",\n        \"fcModuleType\": 0\n      }]\n    },\n    \"activity\": {\n      \"title\": \"??????\",\n      \"searchType\": 4,\n      \"data\": [{\n        \"article_id\": 1053,\n        \"title\": \"XXXX.??????????????? | \\\"??????\\\",????????????????????????????????????????????????\",\n        \"content\": \"\",\n        \"author\": null,\n        \"business_brand_name\": null,\n        \"reads\": 0,\n        \"is_top\": 0,\n        \"is_hot\": 0,\n        \"is_discuss\": 0,\n        \"is_red\": 0,\n        \"user_id\": null,\n        \"user_name\": null,\n        \"sort_no\": 80,\n        \"articletype_id\": 0,\n        \"articletype\": null,\n        \"image_id\": null,\n        \"image\": {\n          \"file_id\": 10449,\n          \"path\": \"http://xxxx.jpg\",\n          \"grade_code\": \"\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"grade_code\": \"\",\n        \"imageList\": [],\n        \"keydes\": null,\n        \"photo\": null,\n        \"iscollect\": 0,\n        \"utime\": \"0001-01-01T00:00:00\",\n        \"likecount\": 0,\n        \"commentlist\": null,\n        \"statusName\": null,\n        \"contentType\": 0\n      }]\n    },\n    \"book\": {\n      \"title\": \"??????\",\n      \"searchType\": 5,\n      \"data\": [{\n        \"content_id\": 1582,\n        \"title\": \"???????????????\",\n        \"summarize\": \"???????????????????????????????????????????????????????????????????????????????????????????????????????????????\",\n        \"image\": {\n          \"file_id\": 10415,\n          \"path\": \"http://xxxx.png\",\n          \"grade_code\": \"\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"ageGroupStart\": 3,\n        \"ageGroupEnd\": 3\n      }]\n    },\n    \"toy\": {\n      \"title\": \"?????????\",\n      \"searchType\": 6,\n      \"data\": [{\n        \"content_id\": 1629,\n        \"title\": \"?????????\",\n        \"summarize\": \"???????????????????????????????????????????????????1???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????\",\n        \"image\": {\n          \"file_id\": 10805,\n          \"path\": \"http://xxxx.jpg\",\n          \"grade_code\": \"\",\n          \"sort_no\": 99,\n          \"filetype_name\": \"pictext\"\n        },\n        \"ageGroupStart\": 5,\n        \"ageGroupEnd\": 6\n      }]\n    }\n  }\n}"
	var info map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &info)

	if err != nil {
		log.Fatalf("Unmarshal err: %s", err)
		return
	}

	Tools.JsonInfo.Init()
	Tools.Check("DataModel", info)

	log.Println("----------------------------------")

	Tools.JsonInfo.Gen()
}
