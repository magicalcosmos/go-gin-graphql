// 从schema开始
// schema/schema.go
package schema

// 引入包graphql-go
import (
	"github.com/graphql-go/graphql"
)

// 定义跟查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello": &queryHello, // queryHello 参考schema/hello.go
	},
})

// 定义Schema用于http handler处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, // 需要通过GraphQL更新数据，可以定义Mutation
})