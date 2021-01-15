// 准备好了GraphQL在Go里面需要的东西之后，来看看如何跟Gin结合
// controller/graphql/graphql.go
package graphql

import (
	"demo/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc{
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
		GraphiQL: true,
	})
	
  // 只需要通过Gin简单封装即可
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}