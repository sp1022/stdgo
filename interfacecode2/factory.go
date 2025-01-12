package interfacecode2

func Factory(name string) IDuck {
	switch name {
	case "duck":
		return &Duck{Color: "write", Age: 5}
	case "goose":
		return &Goose{Color: "blacke"}
	default:
		panic("no this animal")
	}
}
