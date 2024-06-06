package utils

type SarifReport struct {
	Version string `json:"version"`
	Schema  string `json:"$schema"`
	Runs    []Run  `json:"runs"`
}

type Run struct {
	Tool      Tool       `json:"tool"`
	Artifacts []Artifact `json:"artifacts"`
	Results   []Result   `json:"results"`
}

type Tool struct {
	Driver Driver `json:"driver"`
}

type Driver struct {
	Name  string `json:"name"`
	URI   string `json:"informationUri"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID          string      `json:"id"`
	Description Description `json:"shortDescription"`
	HelpURI     string      `json:"helpUri"`
	Properties  Properties  `json:"properties"`
}

type Description struct {
	Text string `json:"text"`
}

type Properties struct {
	Category string `json:"category"`
}

type Artifact struct {
	Location Location `json:"location"`
}

type Location struct {
	URI string `json:"uri"`
}

type Result struct {
	Level     string     `json:"level"`
	Message   Message    `json:"message"`
	RuleID    string     `json:"ruleId"`
	Locations []Location `json:"locations"`
	RuleIndex int        `json:"ruleIndex"`
}

type Message struct {
	Text string `json:"text"`
}

type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Region           Region           `json:"region"`
}

type ArtifactLocation struct {
	URI   string `json:"uri"`
	Index int    `json:"index"`
}

type Region struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}
