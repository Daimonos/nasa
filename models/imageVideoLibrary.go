package models

type Data struct {
	Center      string   `json:"center"`
	DateCreated string   `json:"date_created"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	MediaType   string   `json:"media_type"`
	NasaID      string   `json:"nasa_id"`
	Title       string   `json:"title"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Render string `json:"render"`
}

type Item struct {
	Data  []Data `json:"data"`
	Href  string `json:"href"`
	Links Link   `json:"links"`
}

type PhotoCollection struct {
	Href     string `json:"href"`
	Items    []Item `json:"items"`
	Links    Link   `json:"links"`
	MetaData struct {
		TotalHits string `json:"total_hits"`
	} `json:"metadata"`
	Version string `json:"version"`
}

type SearchResponse struct {
	Collection PhotoCollection
}

type Href struct {
	Href string `json:"href"`
}

type AssetCollection struct {
	Href  string `json:"href"`
	Items []Href `json:"items"`
}

type AssetResponse struct {
	Collection AssetCollection `json:"collection"`
}

type MetadataResponse struct {
	Location string `json:"location"`
}

type VideoCaptionResponse struct {
	Location string `json:"location"`
}

type AlbumData struct {
	NasaID      string   `json:"nasa_id"`
	Album       []string `json:"album"`
	Keywords    []string `json:"keywords"`
	Title       string   `json:"title"`
	MediaType   string   `json:"media_type"`
	DateCreated string   `json:"date_created"`
	Center      string   `json:"center"`
	Description string   `json:"description"`
}

type AlbumItem struct {
	Data  []AlbumData `json:"data"`
	Href  string      `json:"href"`
	Links []Link      `json:"links"`
}

type AlbumResponse struct {
	Href     string      `json:"href"`
	Items    []AlbumItem `json:"items"`
	Links    []Link      `json:"links"`
	Metadata struct {
		TotalHits int `json:"total_hits"`
	} `json:"metadata"`
	Version string `json:"version"`
}
