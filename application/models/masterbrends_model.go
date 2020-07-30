package models


type MasterBrand struct {
	Id        	  		string `json:"id"`
	Business_id        	int `json:"business_id"`
	Mobile_id        	string `json:"mobile_id"`
	Name        	  	string `json:"name"`
	Description        	string `json:"description"`
	Additional        	string `json:"additional"`
}


type JsonBrands struct{

	Nama          string `json:"nama"`
	Business_id   string `json:"business_id"`
	Mobile_id     string `Json:"mobile_id`
	Description   string `json:"description"`
	Created_by     string `json:"created_by"`
}