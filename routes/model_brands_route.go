package routes

import(
	"../application/controllers"
	"github.com/labstack/echo"
)

func MasterBrand (g *echo.Group){
	DEFINE_URL :="/brands"
	g.GET(DEFINE_URL+"/get_data/", controllers.GetDataMasterBrandsController)
	g.POST(DEFINE_URL+"/insert/", controllers.InserMasterBrendsController)
	g.POST(DEFINE_URL+"/update/",controllers.UpdateMasterBrands)
	g.POST(DEFINE_URL+"/delete/", controllers.DeleteMasterBrandsController)
}

