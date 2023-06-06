package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/Public", "Public")
	e.Static("Source", "Source")
	e.GET("/", Main_Page)
	e.GET("/Project-Detail/:id", Static_Project_Detail)
	e.POST("/", Project_Form_Value)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func Main_Page(c echo.Context) error {
	var tmpl, err = template.ParseFiles("Source/Views/Index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func Static_Project_Detail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"ID": id,
	}
	tmpl, err := template.ParseFiles("Source/Views/Project-Detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func Project_Form_Value(c echo.Context) error {
	Project_Title := c.FormValue("project-title")
	Start_Date := c.FormValue("start-date")
	Finsih_Date := c.FormValue("finish-date")
	Description := c.FormValue("description")
	Toggle_A := c.FormValue("input-value-js")
	Toggle_B := c.FormValue("input-value-bs")
	Toggle_C := c.FormValue("input-value-go")
	Toggle_D := c.FormValue("input-value-react")
	println(
		Project_Title, 
		Start_Date, 
		Finsih_Date, 
		Description, 
		Toggle_A, 
		Toggle_B, 
		Toggle_C, 
		Toggle_D)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
