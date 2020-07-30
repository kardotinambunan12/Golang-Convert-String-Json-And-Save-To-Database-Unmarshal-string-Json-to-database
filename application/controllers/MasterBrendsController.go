package controllers

import(
	"net/http"
	"encoding/json"
	"../models"
	"../../database"
	"github.com/labstack/echo"
)


func GetDataMasterBrandsController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, mobile_id, business_id, description FROM brands ")
	if err != nil {
		logs.Println(err)
		internal_server_error(c)
		return nil
	}
		defer rows.Close()

	each   := models.MasterBrand{}
	result := []models.MasterBrand{}

	for rows.Next() {
		var id, mobile_id, name, description, business_id[]byte
		err := rows.Scan(&id, &mobile_id, &name, &business_id, &description)
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil 
		}
	//Additional      string `json:"additional"`
		each.Id            = string(id)
		each.Name          = string(name)
		each.Mobile_id     = string(mobile_id)
		each.Business_id   = ConvertStringToInt(string(business_id))
		each.Description   = string(description)

		result          = append(result, each)
	}
	response     := response_json{
		"data"   :   result,
		"status" :   status_200,	
	}
	return c.JSON(200, response)
}

func InserMasterBrendsController(c echo. Context) error  {
	//membuka koneksi ke database
	db := database.CreateCon()
	//menutup database
	defer db.Close()

	//ambil data dari hasil json
	payload := c.FormValue("data")

	 var data_brands_master []models.JsonBrands
	//menyimpan hasil json ke struct model
	json.Unmarshal([]byte(string(payload)), &data_brands_master)
	for _,  nilai := range data_brands_master{

		name        := nilai.Nama
		business_id := ConvertStringToInt(nilai.Business_id)
		mobile_id   := nilai.Mobile_id
		description := nilai.Description
		created_by  := ConvertStringToInt(nilai.Created_by)

		//proses_insert data
		 sql := "INSERT INTO brands(name, mobile_id, business_id, description, created_by, created_at) VALUES  (?, ?, ?, ?, ?, ?) "
		insert_brands, err := db.Prepare(sql)
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil
		}
		defer insert_brands.Close()
		 insert_brands.Exec(name, mobile_id, business_id, description, created_by, current_time("2006-01-02 15:04:05"))
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil
		}

	}
	response := response_json{
		"status" : status_200,
	}
	return c.JSON(http.StatusOK, response)


}

func UpdateMasterBrands(c echo.Context) error {
	db  := database.CreateCon()

	defer db.Close()

	//ambil data dari string json
	payload :=c.FormValue("data")

	//konversi data string ke json
	var data_brands_master []models.JsonBrands

	//menyimpan hasil data dari json ke struct model
	 json.Unmarshal([]byte(string(payload)), &data_brands_master)
	 if data_brands_master == nil{
		response := response_json{
			"status"	: status_204,
		}
		return c.JSON(200, response)
    } 
     //proses looping value databrands ke database
	for _, nilai := range data_brands_master{

		name          :=  nilai.Nama
		business_id   := ConvertStringToInt(nilai.Business_id)
		mobile_id     := nilai.Mobile_id
		description   := nilai.Description

		update_brands, err := db.Prepare("UPDATE brands SET name =?, business_id=?, mobile_id=?, description=?, updated_at=? WHERE mobile_id=?")
		defer update_brands.Close()
		if  err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil
		}
		update_brands.Exec(name, business_id, mobile_id,  description, current_time("2006-01-02- 15:04:05"), mobile_id )

	}

	response := response_json{
		"status " : status_200,
	}
	return c.JSON(200, response)

}
func DeleteMasterBrandsController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	//ambil data string json
	payload := c.FormValue("data")

	//konversi data dari json bentuk string 
	var data_brands_master []models.JsonBrands

	//menyimpan hasil dari json ke struct modol
	json.Unmarshal([]byte(string(payload)), &data_brands_master)
	for _, nilai   := range data_brands_master {

		mobile_id  := nilai.Mobile_id

		//proses delete data

		sql  := "DELETE FROM brands WHERE mobile_id =?"
		delete_brands, err := db.Prepare(sql)
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil 
		}
		defer delete_brands.Close()
		delete_brands.Exec(mobile_id)

	}
	response := response_json {
		"status"  : status_200,
	}
	return c.JSON(200, response)
}
