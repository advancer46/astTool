package astTool

import (
	"go/ast"
)

/**
提供方法查看或者搜素ast节点
*/
type MyVisitor struct {
	Result []ast.Node
	Name   string
	Type   string
}

func (v MyVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}
	return v
}

func (v *MyVisitor) Inspector(node ast.Node) bool {
	switch x := node.(type) {
	case *ast.FuncDecl:
		if v.Name == x.Name.Name && v.Type == "FuncDecl" {
			v.Result = append(v.Result, x)
			return false
		}
	case *ast.ValueSpec:
		for _, item := range x.Names {
			if v.Name == item.Name && v.Type == "ValueSpec" {
				v.Result = append(v.Result, x)
				return false
			}
		}
	case *ast.Field:
		for _, item := range x.Names {
			if v.Name == item.Name && v.Type == "Field" {
				v.Result = append(v.Result, x)
				return false
			}
		}

	case *ast.GenDecl:
		switch x.Tok.String() {
		case "type":
			for _, item := range x.Specs {
				switch y := item.(type) {
				case *ast.TypeSpec:
					if v.Name == y.Name.Name && v.Type == "GenDecl" {
						v.Result = append(v.Result, x)
						return false
					}
				default:

				}

			}
		case "var":
			for _, item := range x.Specs {
				switch y := item.(type) {
				case *ast.ValueSpec:
					for _, it := range y.Names {
						if v.Name == it.Name && v.Type == "GenDecl" {
							v.Result = append(v.Result, x)
							return false
						}
					}
				default:

				}
			}

		}

	default:

	}

	return true
}

/**
搜索器
按照节点名，节点类型
*/
type Searcher struct {
	Root ast.Node
}

func (v Searcher) FindFuncDeclGlobal(name string) *ast.FuncDecl {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "FuncDecl"}

	ast.Inspect(v.Root, visitor.Inspector)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.FuncDecl)
	} else {
		return nil
	}
}

func (v Searcher) FindValueSpecGlobal(name string) *ast.ValueSpec {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "ValueSpec"}

	ast.Inspect(v.Root, visitor.Inspector)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.ValueSpec)
	} else {
		return nil
	}
}

/*
示例:
package miclient
type s struct{
 	a int32
}
====>
type s struct{
 	a int32
}
*/
func (v Searcher) FindTypeDecl(name string) *ast.GenDecl {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "GenDecl"}

	ast.Inspect(v.Root, visitor.Inspector)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.GenDecl)
	} else {
		return nil
	}
}

/*
示例:
package miclient

func NewPartnerSvcEndpoints(service service.PartnerSvcService) int {
	return 0
}
====>
func NewPartnerSvcEndpoints(service service.PartnerSvcService) int {
	return 0
}
*/
func (v Searcher) FindFuncDecl(name string) *ast.FuncDecl {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "FuncDecl"}

	ast.Inspect(v.Root, visitor.Inspector)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.FuncDecl)
	} else {
		return nil
	}
}

/*
示例:
package miclient
type s struct{
 	a int32
}

=====>
a int32
*/
func (v Searcher) FindField(name string) *ast.Field {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "Field"}

	ast.Inspect(v.Root, visitor.Inspector)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.Field)
	} else {
		return nil
	}
}
