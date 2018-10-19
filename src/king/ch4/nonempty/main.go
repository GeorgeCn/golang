package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s :=range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	maps := map[string]int{"H":110}
	x, y := maps["B"]
	fmt.Println(maps)
	fmt.Println(x)
	fmt.Println(y)
}

