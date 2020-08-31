package models

type MasterBrand struct {
	Id          string `json:"id"`
	Business_id int    `json:"business_id"`
	Mobile_id   string `json:"mobile_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Additional  string `json:"additional"`
}

type Brands struct {
	Id          int    `gorm:"AUTO_INCREMENT"`
	Name        string `gorm: "type:varchar(50)"`
	Mobile_id   string `gorm: "type:int(11)"`
	Business_id int    `gorm: "type:varchar(50)"`
	Description string `gorm:"type:text"`
	Created_at  string `gorm:"type:varchar(100)"`
	//Created_by  int    `gorm:"type:int(11)"`
}

func (Brands) TableName() string {
	return "brands"
}

type JsonBrands struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Business_id       int    `json:"business_id"`
	Mobile_id         string `Json:"mobile_id"`
	Description       string `json:"description"`
	Created_by        int    `json:"created_by"`
	Created_at        string `json:"Created_at"`
	Updated_at        string `json:"updated_at"`
	Additional_brands string `json:"additional_brands"`
}
