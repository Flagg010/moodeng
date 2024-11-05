package main
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "io/ioutil"
    "log"
)

func encryptFile(filename string, key []byte) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }
    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return err
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

    return ioutil.WriteFile(filename+".enc", ciphertext, 0644)
}

func main() {
    key := []byte("examplekey1234567890abcdef1234") // Clé de 32 octets (256 bits)
    err := encryptFile("C:\\path\\to\\file.docx", key)
    if err != nil {
        log.Fatal("Encryption failed:", err)
    }
    log.Println("File encrypted successfully.")
}
