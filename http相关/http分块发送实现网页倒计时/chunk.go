package main

import (
	"fmt"
	"net/http"
	"time"
)

var updatecount = `
<script>
document.getElementById("msg").innerHTML="%d";
</script>
`

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("<h1 id='msg'></h1>"))
		//分块发送
		rw.Header().Set("Transfer-Encoding", "chunked")

		for i := 10; i > -1; i-- {
			rw.Write([]byte(fmt.Sprintf(updatecount, i)))
			time.Sleep(time.Second * 1)
			rw.(http.Flusher).Flush()
		}

	})
	http.ListenAndServe(":8080", nil)
}
