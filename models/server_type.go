package models

type ServerType struct {
	Name string
	URL  string
}

type serverTypes []ServerType

var ServerTypes = serverTypes{
	{Name: "Vanilla", URL: "https://piston-data.mojang.com/v2/objects/"},
}
