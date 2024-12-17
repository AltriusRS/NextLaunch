package translations

type Translation struct {
	Code  string `json:"code"`
	File  string `json:"file"`
	Sha25 string `json:"sha256"`
	Size  int    `json:"size"`
}

type Translations struct {
	Version string        `json:"version"`
	Results int           `json:"results"`
	Files   []Translation `json:"files"`
}
