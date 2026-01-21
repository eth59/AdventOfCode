package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

//go:embed templates/*.template
var templateFS embed.FS

type TemplateData struct {
	Day, Year int
}

func main() {
	// on charge les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  Attention : Impossible de charger le fichier .env, utilisation de l'environnement existant")
	}

	// valeurs par défaut
	year := time.Now().Year()
	month := time.Now().Month().String()
	day := time.Now().Day()

	// vérification qu'on est en décembre si aucun arg
	if len(os.Args) == 1 && month != "December" {
		fmt.Println("❌ Erreur : Nous ne sommes pas en décembre.")
		os.Exit(1)
	}
	
	// lecture du jour (arg 1)
	if len(os.Args) > 1 {
		d, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("❌ Erreur : Le jour '%s' n'est pas un nombre valide.\n", os.Args[1])
			os.Exit(1)
		}
		if d < 1 || d > 25 {
			fmt.Println("❌ Erreur : Le jour doit être entre 1 et 25.")
			os.Exit(1)
		}
		day = d
	}

	// lecture de l'année (arg 2)
	if len(os.Args) > 2 {
		y, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("❌ Erreur : L'année '%s' n'est pas un nombre valide.\n", os.Args[2])
			os.Exit(1)
		}
		if y < 2015 || y > year {
			fmt.Println("❌ Erreur : Année invalide pour l'Advent of Code.")
			os.Exit(1)
		}
		year = y
	}

	// à partir de 2025, seulement 12 jours
	if year >= 2025 && day > 12 {
		fmt.Println("❌ Erreur : A partir de 2025, il n'y a plus que 12 jours.")
		os.Exit(1)
	}

	fmt.Printf("⌛ Préparation du Jour %d de l'Année %d...\n", day, year)

	// récupération du cookie de session
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		fmt.Println("❌ Erreur : La variable d'environnement AOC_SESSION n'est pas définie.")
		os.Exit(1)
	}

	// on crée le dossier du jour
	dirName := fmt.Sprintf("%d/day%d", year, day)
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Printf("❌ Erreur lors de la création du dossier %s : %v\n", dirName, err)
    	os.Exit(1)
	} else {
		fmt.Printf("📁 Dossier %s créé avec succès.\n", dirName)
	}
	
	// téléchargement de l'input
	inputURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, _ := http.NewRequest("GET", inputURL, nil)
	req.Header.Add("Cookie", "session=" + sessionCookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("⚠️  Erreur lors du téléchargement (Status: %d). Êtes-vous connecté à internet ?\n", resp.StatusCode)
	} else {
		// on sauvegarde l'input
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("⚠️  Erreur lors de la lecture de l'input : %v\n", err)
		} else {
			cleanInput := strings.TrimSpace(string(body))
			filePath := dirName + "/input.txt"
			err = os.WriteFile(filePath, []byte(cleanInput), 0644)
			if err != nil {
				fmt.Printf("⚠️  Erreur lors de la création du fichier input.txt : %v\n", err)
			} else {
				fmt.Printf("📥 Input téléchargée et sauvegardée dans %s.\n", filePath)
			}
		}
	}
	defer resp.Body.Close()

	// on récupére l'input test et on la sauvegarde
	testInput := scrapTestInput(year, day, sessionCookie)
	filePath := dirName + "/input_test.txt"
	err = os.WriteFile(filePath, []byte(testInput), 0644)
	if err != nil {
		fmt.Printf("⚠️  Erreur lors de l'écriture de l'input de test : %v\n", err)
	} else {
		fmt.Printf("🧪 Input de test sauvegardée dans %s.\n", filePath)
	}

	// génération des fichiers : *.go / go.mod / README.md
	templateMap := map[string]string{
		"go.mod.template": "go.mod",
		"main_test.go.template": "main_test.go",
		"main.go.template": "main.go",
		"part1.go.template": "part1.go",
		"part2.go.template": "part2.go",
		"README.md.template": "README.md",
	}
	data := TemplateData{day, year}
	for tmplName, fileName := range templateMap {
		target := fmt.Sprintf("%s/%s", dirName, fileName)
		generateFromTemplate(tmplName, target, data)
	}

	// ajout du nouveau module au workspace go.work
	cmd := exec.Command("go", "work", "use", "./"+dirName)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("⚠️  Impossible d'ajouter le dossier au go.work : %v\n", err)
		fmt.Printf("👉 Ajoute-le manuellement : go work use ./%s\n", dirName)
	} else {
		fmt.Println("🔧 Dossier ajouté au go.work avec succès.")
	}

	fmt.Printf("✅ Setup du Jour %d (%d) terminé dans %s !\n", day, year, dirName)
}

// génére des fichiers depuis le dossier templates
func generateFromTemplate(templateName, targetPath string, data TemplateData) {
	// vérification de l'existance du fichier pour pas overwrite
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("⚠️  Le fichier %s existe déjà, saut de la génération.\n", targetPath)
		return
	}

	// lecture du fichier template
	tmpl, err := template.ParseFS(templateFS, "templates/" + templateName)
	if err != nil {
		fmt.Printf("⚠️  Erreur dans le parsing du template %s : %v\n", templateName, err)
		return
	}

	file, _ := os.Create(targetPath)
	defer file.Close()

	tmpl.Execute(file, data)
	fmt.Printf("📝 Fichier %s généré avec succès.\n", targetPath)
}

// scrap l'input de test, devrait fonctionner pour les années >= 2020
func scrapTestInput(year, day int, sessionCookie string) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "session=" + sessionCookie)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("⚠️  Impossible de récuupérer la page de consigne pour l'input de test.")
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	html := string(body)

	// on cherche le code dans la première balise <pre><code>
	startTag := "<pre><code>"
	endTag := "</code></pre>"

	startIndex := strings.Index(html, startTag)
	if startIndex == -1 {
		fmt.Println("⚠️  Aucun input test trouvé sur la page.")
		return ""
	}

	// on avance l'index après la balise ouvrante
	startIndex += len(startTag)

	// on cherche la balise fermante à partir de là
	endIndex := strings.Index(html[startIndex:], endTag)
	if endIndex == -1 {
		fmt.Println("⚠️  Aucun input test trouvé sur la page.")
		return ""
	}

	testInput := html[startIndex:startIndex+endIndex]
	// on nettoie l'input des <em> au cas où il y en a
	testInput = strings.ReplaceAll(testInput, "<em>", "")
    testInput = strings.ReplaceAll(testInput, "</em>", "")

	return strings.TrimSpace(testInput)
}