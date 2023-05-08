package limiter

import (
	"bytes"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	for i := 0; i < 200; i++ {
		tag := i
		if i%20 == 0 {
			time.Sleep(time.Second * 2)
			// time.Sleep(time.Second * 1)
		}
		go func() {
			client := &http.Client{Timeout: 100 * time.Second}
			resp, err := client.Get("http://127.0.0.1:8080/api/user/login/role")
			if err != nil {
				log.Println("Err:", tag, err)
			} else {
				var buffer [512]byte
				result := bytes.NewBuffer(nil)
				n, _ := resp.Body.Read(buffer[0:])
				result.Write(buffer[0:n])
				log.Println(result)
			}
			defer client.CloseIdleConnections()
		}()
	}
}
