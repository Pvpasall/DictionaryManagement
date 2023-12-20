package main

import (
	"dictionnaire/dictionnary"
)

func main() {

	dict := dictionnary.New("dictionnary.json")

	dict.Add("mot", "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs")
	dict.Add("mot2", "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs")
	dict.Add("rétrospective", "Exposition, manifestation qui présente les œuvres et l'évolution d'un artiste, d'une école.")

	dict.Get("mot")
	dict.List()
	dict.Remove("mot")
	dict.Get("mot")

}
