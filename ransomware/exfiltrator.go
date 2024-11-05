package main
import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func sendData(filePath string, serverURL string) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(data))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/octet-stream")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending data:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Data sent successfully, status code:", resp.StatusCode)
}

func main() {
    sendData("C:\\path\\to\\file.zip", "https://example.com/upload")
}
