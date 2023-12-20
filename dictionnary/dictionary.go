package dictionnary

import (
	"fmt"
)

type Crud struct {
	m map[string]string
}

func New() Crud{
	c := Crud{m: make(map[string]string)}
	return c
}

func (c *Crud) Add(mot string, definition string) {
    if c.m == nil {
        c.m = make(map[string]string)
    }
    c.m[mot] = definition
}

func (c Crud) Get(mot string){
	_, prs := c.m[mot]
	if prs == false {
		fmt.Println("Mot inexistant")
	}else {
		fmt.Println("Mot trouve")
		fmt.Println(c.m[mot])
	}
}

func (c Crud) List(){
	fmt.Println("Map list:", c.m)
}

func (c *Crud) Remove(mot string){
	delete(c.m, mot)
	fmt.Println(mot + " suprimme avec succes")
}
