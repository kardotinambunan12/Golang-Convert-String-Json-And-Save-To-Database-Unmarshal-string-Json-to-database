package controllers

import (
	"net/http"
    //"strconv"
	"../../database"
	"../models"
	"github.com/labstack/echo"
)

    //penggunaan metode framework echo dengan gorm
func GetDataMasterBrandsController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, name, mobile_id, business_id, description, created_at, updated_at FROM brands").Rows()
	if err != nil {
		logs.Println(err)
		internal_server_error(c)
		return nil
	}
	defer rows.Close()
	each := models.JsonBrands{}
	result := []models.JsonBrands{}

	for rows.Next() {
		var id, name, mobile_id, business_id, description,  created_at, updated_at[]byte
		err := rows.Scan(&id, &name, &mobile_id, &business_id, &description, &created_at, &updated_at)
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil
		}
		each.Id                 = string(id)
		each.Name               = string(name)
		each.Mobile_id          = string(mobile_id)
		each.Business_id        = ConvertStringToInt(string(business_id))
		each.Description        = string(description)
		each.Created_at         = string(created_at)
        each.Updated_at         = string(updated_at)
		each.Additional_brands  = ConvertToMD5(string(id))

		
		result                  = append(result, each)
	}
	response := response_json{
		"data":   result,
		"status": status_200,
	}
	return c.JSON(http.StatusOK, response)

}

// func GetDataMasterBrandsController(c echo.Context) error {
// 	db := database.CreateCon()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name, mobile_id, business_id, description FROM brands ")
// 	if err != nil {
// 		logs.Println(err)
// 		internal_server_error(c)
// 		return nil
// 	}
// 		defer rows.Close()

// 	each   := models.MasterBrand{}
// 	result := []models.MasterBrand{}

//     //looping baris yang akan di ambil datanya  dari databse
// 	for rows.Next() {
// 		var id, mobile_id, name, description, business_id[]byte

// 		//periksa data atau isi item dari database dan arahkan  sesuai dengan  database
// 		err := rows.Scan(&id, &mobile_id, &name, &business_id, &description)
// 		if err != nil {
// 			logs.Println(err)
// 			internal_server_error(c)
// 			return nil
// 		}
// 		//menampilkan data untuk json sesuai dengan mdel
// 	//Additional      string `json:"additional"`
// 		each.Id            = string(id)
// 		each.Name          = string(name)
// 		each.Mobile_id     = string(mobile_id)
// 		each.Business_id   = ConvertStringToInt(string(business_id))
// 		each.Description   = string(description)

// 		result             = append(result, each)
// 	}
// 	response     := response_json{
// 		"data"   :   result,
// 		"status" :   status_200,
// 	}
// 	return c.JSON(200, response)
// }
   
   //penggunaan metode framework echo dengan gorm
func AddMasterBrandsController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	name        := c.FormValue("name")
	business_id := c.FormValue("business_id")
	mobile_id   := c.FormValue("mobile_id")
	description := c.FormValue("description")


	if name == "" {
		response := response_json{
			"status":  status_400,
			"message": "name_is_required",
		}
		return c.JSON(200, response)
	}
	if business_id == "" {
		response := response_json{
			"status":  status_400,
			"message": "business_id_is_required",
		}
		return c.JSON(200, response)
	}
	if mobile_id == "" {
		response := response_json{
			"status":  status_400,
			"message": "mobile_id_is_required",
		}
		return c.JSON(200, response)
	}
	if description == "" {
		response := response_json{
			"status":  status_400,
			"message": "description_is_required",
		}
		return c.JSON(200, response)
	}

    //id_brands    := id
	create_brands := models.Brands{
		Name        : name,
		Business_id : ConvertStringToInt(business_id),
		Mobile_id   : mobile_id,
		Description : description,
		Created_at  : current_time("2006-01-02 15:04:05"),
		//Created_by  : ConvertStringToInt(string(created_by)),
	}
	db.NewRecord(create_brands)
	if error_insert := db.Create(&create_brands); error_insert.Error != nil {
		logs.Println(error_insert)
		internal_server_error(c)
		return nil
	}
	db.NewRecord(create_brands)

	response := response_json{
		"status": status_200,
	}
	return c.JSON(200, response)
}
    

    //metode penggunaan unmarshal
    
// func InserMasterBrendsController(c echo. Context) error  {
// 	//membuka koneksi ke database
// 	db := database.CreateCon()
// 	//menutup database
// 	defer db.Close()

// 	//ambil data dari hasil json
// 	payload := c.FormValue("data")

// 	 var data_brands_master []models.JsonBrands
// 	//menyimpan hasil json ke struct model
// 	json.Unmarshal([]byte(string(payload)), &data_brands_master)

// 	  if data_brands_master == nil{
// 		response        := response_json{
// 			"status"	: status_204,
// 		}
// 		return c.JSON(200, response)
//     }

// 	for _,  nilai   := range data_brands_master{

// 		name        := nilai.Nama
// 		business_id := ConvertStringToInt(nilai.Business_id)
// 		mobile_id   := nilai.Mobile_id
// 		description := nilai.Description
// 		created_by  := ConvertStringToInt(nilai.Created_by)

// 		//proses_insert data
// 		sql := "INSERT INTO brands(name, mobile_id, business_id, description, created_by, created_at) VALUES  (?, ?, ?, ?, ?, ?) "
// 		insert_brands, err := db.Prepare(sql)
// 		if err != nil {
// 			logs.Println(err)
// 			internal_server_error(c)
// 			return nil
// 		}
// 		// menutup koneksi insert data brands
// 		defer insert_brands.Close()

// 		//melakukan eksekusi insert data brands ke database
// 		 insert_brands.Exec(name, mobile_id, business_id, description, created_by, current_time("2006-01-02 15:04:05"))
// 		if err != nil {
// 			logs.Println(err)
// 			internal_server_error(c)
// 			return nil
// 		}

// 	}
// 	response := response_json{
// 		"status" : status_200,
// 	}
// 	return c.JSON(http.StatusOK, response)

// }


func UpdateMasterBrandsControllers(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	//define FOrm
	id_brands   := c.Param("id")
	name        := c.FormValue("name")
	business_id := c.FormValue("business_id")
	mobile_id   := c.FormValue("mobile_id")
	description := c.FormValue("description")

   
	var update_brands models.Brands 
	data := db.Model(&update_brands).Where("md5(id) =?", id_brands).Updates(map[string]interface{}{
		"name"         : name,
		"business_id"  : business_id,
		"mobile_id"    : mobile_id,
		"description"  : description,
		"updated_at"   : current_time("2006-01-02 15:04:05"),
	})
	if data.Error != nil {
		logs.Println(data.Error)
		internal_server_error(c)
		return nil
	}else if data.RowsAffected == 0 {
       not_modified(c)
       return nil
	}
	response := response_json {
		"status"  : status_200,
	}
	return c.JSON(200, response)

}



  //penggunaan update menggunakan deploy. atau unmarshal

// func UpdateMasterBrandsController(c echo.Context) error {
// 	db  := database.CreateCon()

// 	defer db.Close()

// 	//ambil data dari string json
// 	payload :=c.FormValue("data")

// 	//konversi data string ke json
// 	var data_brands_master []models.JsonBrands

// 	//menyimpan hasil data dari json ke struct model
// 	 json.Unmarshal([]byte(string(payload)), &data_brands_master)

// 	 if data_brands_master == nil{
// 		response        := response_json{
// 			"status"	: status_204,
// 		}
// 		return c.JSON(200, response)
//     }
//      //proses looping value databrands ke database
// 	for _, nilai      := range data_brands_master{

// 		name          :=  nilai.Nama
// 		business_id   := ConvertStringToInt(nilai.Business_id)
// 		mobile_id     := nilai.Mobile_id
// 		description   := nilai.Description

// 		update_brands, err := db.Prepare("UPDATE brands SET name =?, business_id=?, mobile_id=?, description=?, updated_at=? WHERE mobile_id=?")
// 		defer update_brands.Close()
// 		if  err != nil {
// 			logs.Println(err)
// 			internal_server_error(c)
// 			return nil
// 		}
// 		update_brands.Exec(name, business_id, mobile_id,  description, current_time("2006-01-02 15:04:05"), mobile_id )

// 	}
// 	response := response_json{
// 		"status " : status_200,
// 	}
// 	return c.JSON(200, response)

// }



func DeleteMasterBrandsControllers(c echo.Context ) error{
	db := database.CreateCon()
	defer db.Close()


	id_brands   := c.Param("id")
	var master_brands models.Brands
	data := db.Unscoped().Where("md5(id) = ? ", id_brands ).Delete(&master_brands)
	if data.Error != nil {
		logs.Println(data.Error)
		internal_server_error(c)
		return nil
	}else if data.Error == nil{
		response := response_json{
			"status":  status_204,

	}
	return c.JSON(200, response)
   }
	response  := response_json {
		"status"  :status_200,
	}
	return c.JSON(http.StatusOK, response)
}


    //metode penggunaan unmarshal

// func DeleteMasterBrandsController(c echo.Context) error {
// 	db := database.CreateCon()
// 	defer db.Close()

// 	//ambil data string json
// 	payload := c.FormValue("data")

// 	//konversi data dari json bentuk string
// 	var data_brands_master []models.JsonBrands

// 	//menyimpan hasil dari json ke struct modol
// 	json.Unmarshal([]byte(string(payload)), &data_brands_master)
// 	for _, nilai   := range data_brands_master {

// 		mobile_id  := nilai.Mobile_id

// 		//proses delete data

// 		sql  := "DELETE FROM brands WHERE mobile_id =?"
// 		delete_brands, err := db.Prepare(sql)
// 		if err != nil {
// 			logs.Println(err)
// 			internal_server_error(c)
// 			return nil
// 		}
// 		defer delete_brands.Close()
// 		delete_brands.Exec(mobile_id)

// 	}
// 	response := response_json {
// 		"status"  : status_200,
// 	}
// 	return c.JSON(200, response)
// }
