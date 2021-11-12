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

const (
	Create  string = "Create"
	Update  string = "Update"
	Delete  string = "Delete"
	Genesis string = "Genesis"
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
	Index               int64  `json:"index"`
	Timestamp           int64  `json:"timestamp"`
	PreviousHash        string `json:"previousHash"`
	Hash                string `json:"hash"`
	BlockType           string `json:"blockType"`
	CreateOriginHash    string `json:"createOriginHash"`
	Nonce               int64  `json:"nonce"`
} // Block

type Blockchain struct {
	UserId int     `json:"userId"`
	Chain  []Block `json:"chain"`
} // Blockchain
