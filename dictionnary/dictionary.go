package dictionnary

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Dico struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

type CRUD struct {
	filePath string
}

func New(filePath string) CRUD {
	return CRUD{filePath: filePath}
}

func (c CRUD) Add(mot, definition string) string {
	entry := Dico{Mot: mot, Definition: definition}
	entries := c.readDico()
	entries = append(entries, entry)
	c.writeDico(entries)

	return mot + " et sadefinition ajouter avec succes !"
}

func (c CRUD) Get(mot string) string {
	entries := c.readDico()
	for _, entry := range entries {
		if entry.Mot == mot {
			return entry.Definition
		}
	}
	return mot + " introuvable"
}

func (c CRUD) List() []Dico {
	return c.readDico()
}

func (c CRUD) Remove(mot string) string {
	entries := c.readDico()
	var updatedDico []Dico
	for _, entry := range entries {
		if entry.Mot != mot {
			updatedDico = append(updatedDico, entry)
		}
	}
	c.writeDico(updatedDico)

	return mot + " et sa definition supprime avec succes"
}

func (c CRUD) readDico() []Dico {
	file, _ := os.Open(c.filePath)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var entries []Dico
	_ = json.Unmarshal(byteValue, &entries)
	return entries
}

func (c CRUD) writeDico(entries []Dico) {
	file, _ := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	byteValue, _ := json.Marshal(entries)
	_, _ = file.Write(byteValue)
}
