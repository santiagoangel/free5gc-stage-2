package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"free5gc/lib/http2_util"
	. "free5gc/lib/openapi/models"
	"free5gc/lib/path_util"
	"log"
	"net/http"
)

var (
	NRF_LOG_PATH = path_util.Gofree5gcPath("free5gc/src/nrf/Management/sslkeylog.log")
	NRF_PEM_PATH = path_util.Gofree5gcPath("free5gc/support/TLS/nrf.pem")
	NRF_KEY_PATH = path_util.Gofree5gcPath("free5gc/support/TLS/nrf.key")
)

func main() {
	router := gin.Default()

	router.POST("", func(c *gin.Context) {
		/*buf, err := c.GetRawData()
		if err != nil {
			t.Errorf(err.Error())
		}
		// Remove NL line feed, new line character
		//requestBody = string(buf[:len(b uf)-1])*/
		var ND NotificationData

		if err := c.ShouldBindJSON(&ND); err != nil {
			log.Panic(err.Error())
		}
		fmt.Println(ND)
		c.JSON(http.StatusNoContent, gin.H{})
	})

	srv, err := http2_util.NewServer(":30678", NRF_LOG_PATH, router)
	if err != nil {
		log.Panic(err.Error())
	}

	err2 := srv.ListenAndServeTLS(NRF_PEM_PATH, NRF_KEY_PATH)
	if err2 != nil && err2 != http.ErrServerClosed {
		log.Panic(err2.Error())
	}

}
