package main

import (
	"fmt"
)

type crud struct {
    mot string
	m map[string]string
}

func (c crud) get(){
	_, prs := c.m[c.mot]
	if prs == false {
		fmt.Println("Mot inexistant")
	}else {
		fmt.Println("Mot trouve")
		fmt.Println(c.m[c.mot])
	}
}

func (c crud) list(){
	fmt.Println("Map list:", c.m)
}

func (c *crud) remove(){
	delete(c.m, c.mot)
	fmt.Println(c.mot + " suprimme avec succes")
}

func main() {

	c := crud{ mot: "mot", m : make(map[string]string)}


    c.m["mot"] = "Chacun des sons ou groupes de sons (de lettres ou groupes de lettres) correspondant à un sens isolable spontanément, dans le langage ; (par écrit) suite ininterrompue de lettres, entre deux blancs"
    c.m["rétrospective"] = "Exposition, manifestation qui présente les œuvres et l'évolution d'un artiste, d'une école."



	c.get()
	c.list()
	c.remove()
	c.get()
    
}