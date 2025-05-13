package models

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
