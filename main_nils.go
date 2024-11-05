package main

import (
	"os/exec"
	"time"
)

func openWebPage(url string) {
	// Spécifier le chemin d'accès à Microsoft Edge
	cmd := exec.Command("cmd", "/C", "start", "msedge", url)
	err := cmd.Start()
	if err != nil {
		println("Erreur lors de l'ouverture de la page :", err)
	}
}

func main() {
	url := "https://www.zoodio.live/th/live-streaming"
	for {
		openWebPage(url)
		time.Sleep(10 * time.Second) // Ajustez le délai selon vos besoins
	}
}
