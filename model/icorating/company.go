package model

type ICORatingCompany struct {
	Title       string
	Category    string
	Description string
	Website     string
	Whitepaper  string
	Refs        []string
	ICO         map[string]string
	Raitings    map[string]string
}
