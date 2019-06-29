package middleware
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
func  Authmiddleware(c *gin.Context){
	fmt.Println("Hello my dear")
	token := c.GetHeader("Authorization")
	fmt.Println("token : ", token)

	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H {"error":http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}
	c.Next()
	fmt.Println("Goodbye to Romance !!!!")
}
