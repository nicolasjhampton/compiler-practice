package generator

import (
	"fmt"
	"strings"
	"github.com/fatih/structs"
	p "compiler-practice/parser"
)


// type Generator struct {
// 	Tree *p.DefNode
// }


func Generate(tree p.Node) string {
	name := structs.Name(tree)
	fmt.Println( name );
	// treeMap := structs.Map(tree)
	//uncId := structs.Map(treeMap["name"])

	// funcName := tree.(p.DefNode);
	// fmt.Println( funcName );
	// defname := funcName.Name
	// fmt.Println( defname );
	// str := defname.(p.IdentifierNode) 
	// fmt.Println(str.Name)
	switch name {
	case "DefNode":
		def := tree.(p.DefNode)
		name := def.Name.(p.IdentifierNode).Name
		args := def.ArgNames
		body := def.Body
		fmt.Println(body.(p.CallNode).Name)
		return fmt.Sprintf("function %s(%s) { return %s };",
			name,
			joinArgs(args),
			Generate(body))
	case "CallNode":
		call := tree.(p.CallNode)
		name := call.Name.(p.IdentifierNode).Name
		args := call.ArgExprs
		return fmt.Sprintf("%s(%s)",
			name,
			joinArgs(args))
	default:
		panic(fmt.Sprintf("Unexpected node type: %s", name))
	}
}

func joinArgs(args []p.Node) string {
	list := []string{}
	for _, arg := range args {
		list = append(list, arg.(p.IdentifierNode).Name)
	}
	return strings.Join(list, ",")
}