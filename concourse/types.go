package concourse

type SortBy string

const (
	SortByNone   SortBy = "none"
	SortBySemver SortBy = "semver"
)

type Source struct {
	APIToken          string `json:"api_token"`
	ProductSlug       string `json:"product_slug"`
	ProductVersion    string `json:"product_version"`
	Bucket            string `json:"bucket"`
	Endpoint          string `json:"endpoint"`
	Region            string `json:"region"`
	ReleaseType       string `json:"release_type"`
	SortBy            SortBy `json:"sort_by"`
	SkipSSLValidation bool   `json:"skip_ssl_verification"`
	CopyMetadata      bool   `json:"copy_metadata"`
	Verbose           bool   `json:"verbose"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type Version struct {
	ProductVersion string `json:"product_version"`
}

type CheckResponse []Version

type InRequest struct {
	Source  Source   `json:"source"`
	Version Version  `json:"version"`
	Params  InParams `json:"params"`
}

type InParams struct {
	Globs  []string `json:"globs"`
	Unpack bool     `json:"unpack"`
}

type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type OutRequest struct {
	Params OutParams `json:"params"`
	Source Source    `json:"source"`
}

type OutParams struct {
	FileGlob       string `json:"file_glob"`
	MetadataFile   string `json:"metadata_file"`
	Override       bool   `json:"override"`
}

type OutResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata,omitempty"`
}
