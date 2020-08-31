package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func MasterBrand(g *echo.Group) {
	DEFINE_URL := "/brands"
	g.GET(DEFINE_URL+"/get_data/", controllers.GetDataMasterBrandsController)
	g.POST(DEFINE_URL+"/insert/", controllers.AddMasterBrandsController)
	g.POST(DEFINE_URL+"/update/:id/", controllers.UpdateMasterBrandsControllers)
    g.POST(DEFINE_URL+"/delete/:id/", controllers.DeleteMasterBrandsControllers)
}
