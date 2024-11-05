# MooDeng Supremacy

**MooDeng Supremacy** est un programme en Go qui capture et envoie diverses informations depuis le bureau Windows vers un webhook Discord. Le programme inclut plusieurs fonctionnalités de surveillance et d'interactions sur le système de l'utilisateur.

## Fonctionnalités

Le programme MooDeng Supremacy exécute les actions suivantes lors de son exécution :

- Capture et sauvegarde le contenu du presse-papiers dans un fichier `clipboard.txt`.
- Envoie tous les fichiers présents sur le bureau, à condition qu'ils soient sous un certain seuil de taille, vers un webhook Discord.
- Capture et enregistre la liste des processus en cours dans `processlist.txt`.
- Extrait et sauvegarde les informations système dans `sysinfo.txt`.
- Ouvre une URL YouTube en boucle dans le navigateur par défaut de l'utilisateur.
- Affiche une série de pop-ups contenant le message "GET MOODENGED".

## Prérequis

Ce projet nécessite l'installation de Go et la configuration de votre environnement de développement. 

### Installation de Go

1. **Téléchargez Go** :
   - Allez sur [https://go.dev/dl/](https://go.dev/dl/) pour télécharger la dernière version de Go pour votre système d'exploitation.

2. **Installez Go** :
   - **Windows** : Exécutez le fichier `.msi` téléchargé et suivez les instructions.
   - **Linux** : Utilisez les commandes suivantes pour installer Go dans `/usr/local` :
     ```bash
     sudo tar -C /usr/local -xzf go1.xx.linux-amd64.tar.gz
     ```
   - **MacOS** : Installez avec Homebrew :
     ```bash
     brew install go
     ```

3. **Ajoutez Go à votre `PATH`** (si nécessaire) :
   - Sur Linux/MacOS, ajoutez `/usr/local/go/bin` à votre `PATH` dans votre fichier de profil (`.bashrc`, `.zshrc`).

4. **Vérifiez l'installation** :
   - Tapez dans un terminal :
     ```bash
     go version
     ```
   - Vous devriez voir la version de Go installée.

### Configuration du Module Go pour MooDeng Supremacy

1. **Initialisez le module Go** :
   - Dans le répertoire de votre projet, exécutez la commande suivante :
     ```bash
     go mod init github.com/votre-nom-utilisateur/moodeng-supremacy
     ```

## Compilation de MooDeng Supremacy

Après la configuration de votre environnement, compilez le programme en suivant les étapes suivantes :

1. **Compilez pour Windows** :
   - Depuis le répertoire racine de votre projet, exécutez la commande suivante pour créer un exécutable pour Windows :
     ```bash
     GOOS=windows GOARCH=amd64 go build -o moodeng_supremacy.exe main.go
     ```
   - Cette commande crée un fichier exécutable `moodeng_supremacy.exe` pour Windows 64 bits.

## Exécution de MooDeng Supremacy

Après compilation, vous pouvez exécuter le programme de deux manières :

- **Double-cliquez** sur `moodeng_supremacy.exe` dans l'explorateur Windows.
- Ou **utilisez la ligne de commande** pour l'exécuter :
  ```bash
  ./moodeng_supremacy.exe
