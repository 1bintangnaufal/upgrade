package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
	connection "upgrade/Connection"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type Project struct {
	ID int

	Project_Title string

	Start_Date  time.Time
	Finish_Date time.Time
	Duration    time.Duration

	Description string

	Toggle_A bool
	Toggle_B bool
	Toggle_C bool
	Toggle_D bool

	Image string

	Format_Start_Date  string
	Format_Finish_Date string
	Formatted_Duration string
}

func main() {
	connection.Database_Connect()

	e := echo.New()

	e.Static("/Public", "Public")

	e.GET("/", Main_Page)
	e.GET("/Project-Detail/:id", Project_Detail)
	e.GET("/Edit-Project/:id", Edit_Project)

	e.POST("/", Project_Form_Value)
	e.POST("/Delete-Project/:id", Delete_Project)
	e.POST("/Edit-Project/:id", Save_Changes)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func Main_Page(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image FROM tb_project ORDER BY id")

	var result []Project

	for data.Next() {
		var each = Project{}
		err := data.Scan(&each.ID, &each.Project_Title, &each.Start_Date, &each.Finish_Date, &each.Description, &each.Toggle_A, &each.Toggle_B, &each.Toggle_C, &each.Toggle_D, &each.Image)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.Duration = each.Finish_Date.Sub(each.Start_Date)
		each.Formatted_Duration = Duration_Formatting(each.Duration)

		result = append(result, each)
	}

	Projects := map[string]interface{}{
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("Views/Index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	Map_Duration_Formatting := template.FuncMap{
		"Duration_Formatting": Duration_Formatting,
	}

	tmpl = tmpl.Funcs(Map_Duration_Formatting)

	return tmpl.Execute(c.Response(), Projects)
}

func Project_Detail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var Project_Detail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image FROM tb_project WHERE id=$1", id).Scan(&Project_Detail.ID, &Project_Detail.Project_Title, &Project_Detail.Start_Date, &Project_Detail.Finish_Date, &Project_Detail.Description, &Project_Detail.Toggle_A, &Project_Detail.Toggle_B, &Project_Detail.Toggle_C, &Project_Detail.Toggle_D, &Project_Detail.Image)

	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Project not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}

	Project_Detail.Format_Start_Date = Project_Detail.Start_Date.Format("Jan 2, 2006")
	Project_Detail.Format_Finish_Date = Project_Detail.Finish_Date.Format("Jan 2, 2006")

	Project_Detail.Duration = Project_Detail.Finish_Date.Sub(Project_Detail.Start_Date)
	Project_Detail.Formatted_Duration = Duration_Formatting(Project_Detail.Duration)

	data := map[string]interface{}{
		"Project_Detail": Project_Detail,
	}

	tmpl, err_too := template.ParseFiles("Views/Project-Detail.html")

	if err_too != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err_too.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func Duration_Formatting(Duration time.Duration) string {
	if Duration <= 24*time.Hour {
		return "Less than a day"
	}

	Days := int(Duration.Hours() / 24)
	Weeks := Days / 7
	Months := Days / 30
	Years := Months / 12

	if Years > 1 {
		return fmt.Sprintf("%d years", Years)
	} else if Years == 1 {
		return "A year"
	} else if Months > 1 {
		return fmt.Sprintf("%d months", Months)
	} else if Months == 1 {
		return "A month"
	} else if Weeks > 1 {
		return fmt.Sprintf("%d weeks", Weeks)
	} else if Weeks == 1 {
		return "A week"
	} else if Days > 1 {
		return fmt.Sprintf("%d days", Days)
	} else {
		return "A day"
	}
}

func Project_Form_Value(c echo.Context) error {
	Project_Title := c.FormValue("Project_Title")
	Start_Date := c.FormValue("Start_Date")
	Finish_Date := c.FormValue("Finish_Date")
	Description := c.FormValue("Description")

	Toggle_A := c.FormValue("Toggle_A") == "on"
	Toggle_B := c.FormValue("Toggle_B") == "on"
	Toggle_C := c.FormValue("Toggle_C") == "on"
	Toggle_D := c.FormValue("Toggle_D") == "on"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D, "mobile-app.jpg")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	fmt.Println(Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Delete_Project(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Edit_Project(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var Previous_Data = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image FROM tb_project WHERE id=$1", id).Scan(&Previous_Data.ID, &Previous_Data.Project_Title, &Previous_Data.Start_Date, &Previous_Data.Finish_Date, &Previous_Data.Description, &Previous_Data.Toggle_A, &Previous_Data.Toggle_B, &Previous_Data.Toggle_C, &Previous_Data.Toggle_D, &Previous_Data.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	SD_F := Previous_Data.Start_Date.Format("2006-01-02")
	FD_F := Previous_Data.Finish_Date.Format("2006-01-02")

	data := map[string]interface{}{
		"Previous_Data": Previous_Data,
		"SD_F":          SD_F,
		"FD_F":          FD_F,
		"PDID":          Previous_Data.ID,
	}

	tmpl, err_too := template.ParseFiles("Views/Edit-Project-Form.html")

	if err_too != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func Save_Changes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Project_Title := c.FormValue("Project_Title")
	Start_Date := c.FormValue("Start_Date")
	Finish_Date := c.FormValue("Finish_Date")
	Description := c.FormValue("Description")

	Toggle_A := c.FormValue("Toggle_A") == "on"
	Toggle_B := c.FormValue("Toggle_B") == "on"
	Toggle_C := c.FormValue("Toggle_C") == "on"
	Toggle_D := c.FormValue("Toggle_D") == "on"

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_title=$1, start_date=$2, finish_date=$3, description=$4, toggle_a=$5, toggle_b=$6, toggle_c=$7, toggle_d=$8, image=$9 WHERE id=$10", Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D, "mobile-app.jpg", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	redirectURL := fmt.Sprintf("/?id=%d#ppc-container", id)
	return c.Redirect(http.StatusMovedPermanently, redirectURL)
}
