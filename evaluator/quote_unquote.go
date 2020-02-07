package evaluator

import (
	"github.com/bamchoh/monkey_lang/ast"
	"github.com/bamchoh/monkey_lang/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
