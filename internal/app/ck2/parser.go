package ck2

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type CK2Parser struct {
	Filepath  string  `json:"filepath"`
	Namespace string  `json:"namespace"`
	Level     int     `json:"level"`
	Data      []*Node `json:"data"`
	Scope     *Node   `json:"-"`
}

type NodeType string

const (
	Entity   NodeType = "Entity"
	Block    NodeType = "Block"
	Property NodeType = "Property"
)

type Node struct {
	// element is either a Block or a Property
	Type   NodeType `json:"type"`
	Parent *Node    `json:"-"`
	Level  int      `json:"level"`
	// not nil if type is Property
	Key   string `json:"key"`
	Value string `json:"value"`
	// not nil if type is Block
	Data []*Node `json:"data,omitempty"`
}

func NewParser(file_path string) *CK2Parser {
	return &CK2Parser{
		Filepath:  file_path,
		Namespace: "",
		Level:     0,
		Data:      []*Node{},
		Scope:     nil,
	}
}

func (parser *CK2Parser) NewNode(node_type NodeType, key string, value string) *Node {
	return &Node{
		Type:  node_type,
		Level: parser.Level,
		Key:   key,
		Value: value,
		Data:  []*Node{},
	}
}

func (parser *CK2Parser) InsertNode(node_type NodeType, key string, value string) {
	switch node_type {
	case Entity:
		// enter into entity scope
		if parser.Scope == nil {
			parser.Data = append(parser.Data, &Node{
				Type:  Entity,
				Level: parser.Level,
				Key:   key,
				Value: value,
				Data:  []*Node{},
			})
			parser.Scope = parser.Data[len(parser.Data)-1]
			parser.Level++
		}
	case Block:
		fmt.Println("ENTER into scope of:", strconv.Quote(key))
		new_node := parser.NewNode(Block, key, value)
		parser.Scope.Data = append(parser.Scope.Data, new_node)
		new_node.Parent = parser.Scope
		fmt.Println("PARENT:", new_node.Parent.Key)

		if !strings.Contains(value, "}") {
			parser.Scope = new_node
			parser.Level++
		}
		fmt.Println("scope:", parser.Scope)
	case Property:
		new_property := parser.NewNode(Property, key, value)
		parser.Scope.Data = append(parser.Scope.Data, new_property)
		fmt.Println("property:", strconv.Quote(key), "value:", strconv.Quote(value))
	}
}

func (parser *CK2Parser) ParseLine(line []byte) []byte {
	if bytes.Contains(line, []byte("=")) {
		kv := bytes.SplitN(line, []byte(" = "), 2)
		key := string(kv[0])
		value := string(kv[1])
		fmt.Println("key:", strconv.Quote(key), "value:", strconv.Quote(value))

		// namespace
		if parser.Scope == nil && parser.Namespace == "" {
			parser.Namespace = value
		} else {
			if value[0] == byte('{') {
				// enter into entity scope
				if parser.Scope == nil {
					parser.InsertNode(Entity, key, value)
				} else {
					// enter another scope
					parser.InsertNode(Block, key, value)
				}
			} else {
				parser.InsertNode(Property, key, value)
			}
		}
	}

	if len(line) > 0 && line[0] == byte('}') {
		fmt.Println("END of scope:", parser.Scope)
		parser.Scope = parser.Scope.Parent
		fmt.Println("NEXT SCOPE:", parser.Scope)
		parser.Level--

		// if parser.Scope == parser.PrevScope {
		// 	parser.Scope = nil
		// 	parser.Level--
		// } else {
		// 	parser.Scope = parser.Scope.Parent
		// 	parser.Level--
		// }
	}

	fmt.Print("\n")

	return line
}
