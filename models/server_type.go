package models

type ServerType struct {
	Name string
	URL  string
}

type serverTypes []ServerType

var ServerTypes = serverTypes{
	{Name: "Vanilla", URL: "https://piston-data.mojang.com/v2/objects/"},
	{Name: "Paper", URL: "https://papermc.io/api/v2/projects/paper"},
	{Name: "Spigot", URL: "https://hub.spigotmc.org/nexus/content/repositories/snapshots/org/spigotmc/spigot/"},
	{Name: "Purpur", URL: "https://purpur.pl3x.net/api/v2/"},
}
