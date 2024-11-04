package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func main() {
	UserOrAdmin()
}

func UserOrAdmin() {
	// Récupérer le nom de l'utilisateur
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'utilisateur :", err)
		return
	}

	// Afficher le nom de l'utilisateur
	fmt.Println("Nom de l'utilisateur :", currentUser.Username)

	// Vérifier si l'utilisateur est un admin
	isAdmin := false
	if runtime.GOOS == "windows" {
		// Sur Windows, utiliser 'net session' pour vérifier les privilèges admin
		cmd := exec.Command("net", "session")
		err := cmd.Run()
		if err == nil {
			isAdmin = true
		}
	} else {
		// Sur Unix/Linux, vérifier si l'utilisateur est dans le groupe 'sudo'
		cmd := exec.Command("groups", currentUser.Username)
		output, err := cmd.Output()
		if err == nil {
			if string(output) == "" || !strings.Contains(string(output), "sudo") {
				isAdmin = false
			} else {
				isAdmin = true
			}
		}
	}

	// Afficher si l'utilisateur est admin
	if isAdmin {
		fmt.Println("L'utilisateur a des privilèges d'administrateur.")
	} else {
		fmt.Println("L'utilisateur n'a pas de privilèges d'administrateur.")
	}
}
