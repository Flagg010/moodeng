package main
import (
    "fmt"
    "os"
    "path/filepath"
)

func findFiles(root string, extensions []string) []string {
    var filesFound []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            for _, ext := range extensions {
                if filepath.Ext(path) == ext {
                    filesFound = append(filesFound, path)
                }
            }
        }
        return nil
    })
    if err != nil {
        fmt.Println("Error:", err)
    }
    return filesFound
}

func main() {
    files := findFiles("C:\\Users\\", []string{".docx", ".pdf", ".xlsx"})
    for _, file := range files {
        fmt.Println("Found file:", file)
    }
}
