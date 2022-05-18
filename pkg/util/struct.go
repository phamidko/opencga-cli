package util

type Cellbase struct {
	APIVersion string          `json:"apiVersion"`
	Time       int             `json:"time"`
	Params     Params          `json:"params"`
	Responses  []CellResponses `json:"responses"`
}
type Params struct {
	Species string `json:"species"`
	Limit   string `json:"limit"`
}

type CellResponses struct {
	Time            int       `json:"time"`
	NumResults      int       `json:"numResults"`
	Results         []Results `json:"results"`
	NumTotalResults int       `json:"numTotalResults"`
	NumMatches      int       `json:"numMatches"`
	NumInserted     int       `json:"numInserted"`
	NumUpdated      int       `json:"numUpdated"`
	NumDeleted      int       `json:"numDeleted"`
	ID              string    `json:"id"`
}

type Results struct {
	Program     string `json:"Program"`
	GitCommit   string `json:"Git commit"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
	GitBranch   string `json:"Git branch"`
}

//
// OpenCGA
//
type Opencga struct {
	APIVersion     string             `json:"apiVersion"`
	Time           int                `json:"time"`
	Params         Params             `json:"params"`
	Responses      []OpencgaResponses `json:"responses"`
	FederationNode []FederationNode   `json:"federationNode"`
}

type OpencgaResponses struct {
	Time            int       `json:"time"`
	NumResults      int       `json:"numResults"`
	Results         []Results `json:"results"`
	NumTotalResults int       `json:"numTotalResults"`
	NumMatches      int       `json:"numMatches"`
	NumInserted     int       `json:"numInserted"`
	NumUpdated      int       `json:"numUpdated"`
	NumDeleted      int       `json:"numDeleted"`
	ID              string    `json:"id"`
}

type FederationNode struct {
	ID      string `json:"id"`
	URI     string `json:"uri"`
	Commit  string `json:"commit"`
	Version string `json:"version"`
}
