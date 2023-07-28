package model

type AddContactRequest struct {
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Age            int    `json:"age,omitempty"`
	FullAddress    string `json:"full_address,omitempty"`
	DistrictNumber int    `json:"district_number,omitempty"`
	CountryName    string `json:"country_name,omitempty"`
	Region         string `json:"region,omitempty"`
	Number         string `json:"number,omitempty"`
	IMEI           string `json:"imei,omitempty"`
}

type GetContactRequest struct {
	Number string `json:"number,omitempty"`
	IMEI   string `json:"imei,omitempty"`
}
