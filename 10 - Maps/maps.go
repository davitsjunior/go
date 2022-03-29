package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Lavínia",
		"sobrenome": "Tschmerizja",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Davi",
			"ultimo":   "Junior",
		},
		"curso": {
			"nome":   "Proc Dados",
			"campus": "Unopar",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}
	fmt.Println(usuario2)

}
