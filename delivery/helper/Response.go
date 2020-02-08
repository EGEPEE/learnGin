package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondJSON(w *gin.Context, res *http.ResponseWriter, payload interface{}) {
	fmt.Println("status: ", res)

}
