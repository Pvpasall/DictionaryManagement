package main

import (
	"fmt"
)

func get(m map[string]string, mot string){
	_, prs := m[mot]
	if prs == false {
		fmt.Println("Mot inexistant")
	}else {
		fmt.Println(m[mot])
	}
}

func list(m map[string]string){
	fmt.Println("Map list:", m)
}

func remove(m map[string]string, mot string){
	delete(m, mot)
	fmt.Println(mot + " suprimme avec succes")
}

func main() {

    m := make(map[string]string)

    m["mot"] = "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs"
    m["rétrospective"] = "Exposition, manifestation qui présente les œuvres et l'évolution d'un artiste, d'une école."



	get(m, "mot")
	list(m)
	remove(m, "mot")
	get(m, "mot")
    
}