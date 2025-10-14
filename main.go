package eparkktx

import (
	"github.com/gin-gonic/gin"
	"eParkKtx/config"
)


func main(){
	server := gin.Default();

	config.ConnectDatabase()

	

}
