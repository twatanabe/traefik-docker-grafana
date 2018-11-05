package main

import (
    "net/http"
    "fmt"
    "log"
    "os"
    "bufio"
    "encoding/base64"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {  
        
        fmt.Fprint(writer, "<p>Hello from Golang")
        
        imgFile, err := os.Open("go.png")
        if err != nil {
        	log.Fatal(err)
        }
        defer imgFile.Close()


        fInfo, _ := imgFile.Stat()
        var size int64 = fInfo.Size()
        buf := make([]byte, size)

        fReader := bufio.NewReader(imgFile)
        fReader.Read(buf)
        imgBase64Str := base64.StdEncoding.EncodeToString(buf)

		fmt.Fprintf(writer, "<img src='data:image/jpg; base64, " + imgBase64Str + "' alt='go'>")
		fmt.Fprintf(writer,"</p>")
    })
    log.Fatal(http.ListenAndServe(":3000", nil))
}