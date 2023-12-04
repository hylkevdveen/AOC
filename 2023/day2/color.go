package main

type Color struct {
	Name string
	Max  int
}

var (
	Red   = &Color{Name: "red", Max: 12}
	Green = &Color{Name: "green", Max: 13}
	Blue  = &Color{Name: "blue", Max: 14}
)

var Colors = []*Color{Red, Green, Blue}

func getByName(name string) *Color {
	for _, color := range Colors {
		if color.Name == name {
			return color
		}
	}
	return nil
}
