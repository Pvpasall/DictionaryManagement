package httphandler

import (
	"dictionnaire/dictionnary"
	"encoding/json"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func List(w http.ResponseWriter, req *http.Request) {
	dict := dictionnary.New("./dictionnary/dictionnary.json")
	entries := dict.List()
	jsonData, _ := json.Marshal(entries)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Get(w http.ResponseWriter, req *http.Request) {
	dict := dictionnary.New("./dictionnary/dictionnary.json")
	mot := req.URL.Query().Get("mot")
	if mot != "" {
		definition := dict.Get(mot)
		if definition != "introuvable" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "La définition de %s est : %s", mot, definition)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Le mot %s n'a pas été trouvé", mot)
		}
	}
}

func Add(w http.ResponseWriter, req *http.Request) {
	dict := dictionnary.New("./dictionnary/dictionnary.json")
	mot := req.URL.Query().Get("mot")
	definition := req.URL.Query().Get("definition")

	if mot == "" || definition == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erreur lors de l'ajout du mot: %s. Reesayer !!", mot)
	} else {
		w.WriteHeader(http.StatusOK)
		dict.AddAsync(mot, definition)
		fmt.Fprintf(w, "Le mot: %s et sa definition sont ajoutes avec succes", mot)
	}
}

func Remove(w http.ResponseWriter, req *http.Request) {
	dict := dictionnary.New("./dictionnary/dictionnary.json")
	mot := req.URL.Query().Get("mot")

	if mot != "" {
		if dict.Remove(mot) == true {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Le mot: %s et sa definition ont ete supprime avec succes", mot)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Erreur lors de la suppression du mot: %s. Reesayer !!", mot)
		}
	}
}
