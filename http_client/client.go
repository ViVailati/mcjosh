package http_client

import (
	"fmt"
	"net/http"

	"github.com/ViVailati/mcjosh/utils"
)

type MinecraftVersion struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Type        string `json:"type"`
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
}

type MinecraftManifest struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []MinecraftVersion `json:"versions"`
}

// Client is the HTTP client to use for fetching data from the APIs
type Client struct {
	*http.Client
}

// New creates a new instance of the client
func New() *Client {
	return &Client{http.DefaultClient}
}

// GetMinecraftVersions fetches the JSON file containing the versions of Minecraft
func (c *Client) GetMinecraftVersions() ([]MinecraftVersion, error) {
	fmt.Println("Fetching Minecraft versions...")
	resp, err := c.Get("https://piston-meta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manifest MinecraftManifest
	if err := utils.DecodeJSON(resp.Body, &manifest); err != nil {
		return nil, err
	}

	rv := utils.Filter(manifest.Versions, func(v MinecraftVersion) bool {
		return v.Type == "release"
	})

	return rv, nil
}
