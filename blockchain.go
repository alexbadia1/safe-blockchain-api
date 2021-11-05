package main

//================================================================================
// Certificate Type Enums
//================================================================================

const (
	education          string = "Education"
	employment         string = "Employement"
	tradeCertification string = "Trade Certification"
	other              string = "Other"
)

//================================================================================
// Blockchain Structs
//================================================================================

type Block struct {
	UserId              string `json:"userId"`
	CertificateToken    string `json:"certificate_token"`
	CertificateUrl      string `json:"certificate_url"`
	CertificateCatgeory string `json:"certificate_category"`
	InstitionName       string `json:"institution_name"`
	DegreeName          string `json:"degree_name"`
	DateRange           string `json:"date_range"`
	Description         string `json:"description"`
} // Block

type Blockchain struct {
	userId int
	chain  []Block
} // Blockchain
