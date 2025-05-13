package http_client

import (
	"fmt"
	"net/http"

	"github.com/ViVailati/mcjosh/models"
	"github.com/ViVailati/mcjosh/utils"
)

// Client is the HTTP client to use for fetching data from the APIs
type Client struct {
	*http.Client
}

// New creates a new instance of the client
func New() *Client {
	return &Client{http.DefaultClient}
}

// GetMinecraftVersion fetches the JSON file containing the versions of Minecraft
func (c *Client) GetMinecraftVersion() error {
	fmt.Println("Fetching Minecraft versions...")
	resp, err := c.Get("https://piston-meta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var manifest models.MinecraftManifest
	if err := utils.DecodeJSON(resp.Body, &manifest); err != nil {
		return err
	}

	fmt.Println(manifest)
	return nil
}
