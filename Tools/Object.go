package Tools

import "log"

type Object struct {
	name      string
	tType     string
	childType string
	count     int
}

type DataTree struct {
	objectTree map[string][]*Object
}

func (tree *DataTree) InitObjectList(name string) {
	_, ok := tree.objectTree[name]
	if !ok {
		tree.objectTree[name] = make([]*Object, 0)
	}

}

func (tree *DataTree) Append(name string, valName string, tType string) {

	for _, object := range tree.objectTree[name] {
		if object.name == valName {
			object.name = valName
			object.tType = tType
			return
		}
	}

	object := &Object{
		name:  valName,
		tType: tType,
		count: 0,
	}
	tree.objectTree[name] = append(tree.objectTree[name], object)
}

func (tree *DataTree) AppendChild(name string, valName string, tType string, childType string, isList bool) {
	find := false
	for _, object := range tree.objectTree[name] {
		if object.name == valName {
			find = true
			object.name = valName
			object.childType = childType
			object.tType = tType
			object.count++
		}
	}

	count := 0
	if isList {
		count = 1
	}

	if !find {

		object := &Object{
			name:      valName,
			tType:     tType,
			childType: childType,
			count:     count,
		}
		tree.objectTree[name] = append(tree.objectTree[name], object)
	}
}

func (tree *DataTree) Init() {
	tree.objectTree = make(map[string][]*Object)
}

func (tree *DataTree) Print() {
	for s, objects := range tree.objectTree {
		log.Println("base name:", s)

		for _, object := range objects {
			log.Printf("name:%s   type:%s    child type:%s\n", object.name, object.tType, object.childType)
		}

		log.Println("----------------------------------")
	}
}

func (tree *DataTree) Gen() {
	for s, objects := range tree.objectTree {
		Generate(s, objects)
	}
}

var JsonInfo = DataTree{}
