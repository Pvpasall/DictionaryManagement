package main

import (
	"dictionnaire/dictionnary"
	"fmt"
	"time"
)

func main() {

	dict := dictionnary.New("dictionnary.json")

	message := make(chan string)

	if dict == dict {
		go dict.Add("mot", "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs", message)
		getmessage1 := <-message
		go dict.Add("mot2", "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs", message)
		getmessage2 := <-message
		go dict.Add("rétrospective", "Exposition, manifestation qui présente les œuvres et l'évolution d'un artiste, d'une école.", message)
		getmessage3 := <-message

		fmt.Println(getmessage1)

		// Pour voir la suppression directement dans dictionnary.json
		time.Sleep(time.Second)
		fmt.Println(getmessage2)
		fmt.Println(getmessage3)

		dict.Get("mot")
		dict.List()
		go dict.Remove("mot", message)
		getmessage := <-message
		fmt.Println(getmessage)
		dict.Get("mot")
	}

}
