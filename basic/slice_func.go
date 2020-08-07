package main

import(
    "fmt"
)

type PlanetsSlice []string

func (p PlanetsSlice) add(s string)(PlanetsSlice) {
    var o PlanetsSlice
    for _, ps := range p {
        tmp := s + ps
        o = append(o, tmp)
    }
    return o
}

func (p PlanetsSlice) terraform() {
    for i := range p {
        p[i] = "New" + p[i]
    }
}

func main() {
    planets := [...]string{
        "Mercury",
        "Venus",
        "Earth",
        "Mars",
        "Jupiter",
        "Saturn",
        "Uranus",
        "Neptune",
    }

    fmt.Println(len(planets))

    p := planets[:]
    p2 := PlanetsSlice(p)

    p2 = p2.add("new")
    fmt.Println(p2)
}