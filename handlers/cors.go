package handlers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Req.Method == "OPTIONS" {
			c.Abort(200)
		}
	}
}

//
// func AuthRequired() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		res, err := jwt.ParseFromRequest(c.Req, func(t *jwt.Token) ([]byte, error) {
// 			return []byte("someRandomSigningKey"), nil
// 		})
// 		if err != nil {
// 			// c.JSON(200,http.StatusUnauthorized, "Access denied.")
// 			writeError(c.Writer, "Unauthorized", 401, nil)
// 			c.Fail(401, err)
// 			return
// 		}
//
// 		c.Set("user", res.Claims)
// 		log.Println(res)
// 	}
// }
