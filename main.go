package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Project_Title string
	Start_Date    string
	Finish_Date   string
	Duration      string
	Description   string
	Toggle_A      bool
	Toggle_B      bool
	Toggle_C      bool
	Toggle_D      bool
	Icon_A        string
	Icon_B        string
	Icon_C        string
	Icon_D        string
	Label_A       string
	Label_B       string
	Label_C       string
	Label_D       string
}

var Static_Project_Data = []Project{
	{
		Project_Title: "Mobile App - 2023",
		Start_Date:    "2023-01-17",
		Finish_Date:   "2023-05-17",
		Duration:      "4 months",
		Description:   "Quasi, iusto autem voluptas, facilis quidem aliquid sed harum provident iure nobis eaque sin accusantium excepturi consequatur, amet totam magni blanditiis. Voluptatibus natus placeat, maiores voluptates distinctio quia. Saepe porro maxime iste maiores voluptatem deserunt alias, debitis dolore odit eaque. Cupiditate architecto eaque vero debitis velit unde, cum haru consequatur iste. Neque molestias ducimus temporibus ex labore tempore magni, provident maiores? Obcaecati, placeat aliquid, officiis nulla voluptatem asperiores, ipsa unde a vel magni facere inventore necessitatibus.",
		Icon_A:        `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`,
		Icon_B:        `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`,
		Icon_C:        `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`,
		Icon_D:        `<i class="fa-brands fa-react fa-lg fa-fw"></i>`,
		Label_A:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B:       `<div class="d-flex flex-row align-items-center gap-1">
            				<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
            				<p class="m-0">Bootstrap</p>
        				</div>
						`,
		Label_C:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
	{
		Project_Title: "Web App - 2022",
		Start_Date:    "2022-04-11",
		Finish_Date:   "2022-06-22",
		Duration:      "2 months",
		Description:   "Quasi, iusto autem voluptas, facilis quidem aliquid sed harum provident iure nobis eaque sin accusantium excepturi consequatur, amet totam magni blanditiis. Voluptatibus natus placeat, maiores voluptates distinctio quia. Saepe porro maxime iste maiores voluptatem deserunt alias, debitis dolore odit eaque. Cupiditate architecto eaque vero debitis velit unde, cum haru consequatur iste. Neque molestias ducimus temporibus ex labore tempore magni, provident maiores? Obcaecati, placeat aliquid, officiis nulla voluptatem asperiores, ipsa unde a vel magni facere inventore necessitatibus.",
		Icon_A:        `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`,
		Icon_B:        `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`,
		Icon_C:        `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`,
		Icon_D:        `<i class="fa-brands fa-react fa-lg fa-fw"></i>`,
		Label_A:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
	{
		Project_Title: "Desktop App - 2022",
		Start_Date:    "2019-08-06",
		Finish_Date:   "2022-09-08",
		Duration:      "3 years",
		Description:   "Quasi, iusto autem voluptas, facilis quidem aliquid sed harum provident iure nobis eaque sin accusantium excepturi consequatur, amet totam magni blanditiis. Voluptatibus natus placeat, maiores voluptates distinctio quia. Saepe porro maxime iste maiores voluptatem deserunt alias, debitis dolore odit eaque. Cupiditate architecto eaque vero debitis velit unde, cum haru consequatur iste. Neque molestias ducimus temporibus ex labore tempore magni, provident maiores? Obcaecati, placeat aliquid, officiis nulla voluptatem asperiores, ipsa unde a vel magni facere inventore necessitatibus.",
		Icon_A:        `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`,
		Icon_B:        `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`,
		Icon_C:        `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`,
		Icon_D:        `<i class="fa-brands fa-react fa-lg fa-fw"></i>`,
		Label_A:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
	{
		Project_Title: "Server Building - 2022",
		Start_Date:    "2021-03-07",
		Finish_Date:   "2022-05-24",
		Duration:      "1 year",
		Description:   "Quasi, iusto autem voluptas, facilis quidem aliquid sed harum provident iure nobis eaque sin accusantium excepturi consequatur, amet totam magni blanditiis. Voluptatibus natus placeat, maiores voluptates distinctio quia. Saepe porro maxime iste maiores voluptatem deserunt alias, debitis dolore odit eaque. Cupiditate architecto eaque vero debitis velit unde, cum haru consequatur iste. Neque molestias ducimus temporibus ex labore tempore magni, provident maiores? Obcaecati, placeat aliquid, officiis nulla voluptatem asperiores, ipsa unde a vel magni facere inventore necessitatibus.",
		Icon_A:        `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`,
		Icon_B:        `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`,
		Icon_C:        `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`,
		Icon_D:        `<i class="fa-brands fa-react fa-lg fa-fw"></i>`,
		Label_A:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
	{
		Project_Title: "Pecel Lele",
		Start_Date:    "2023-06-07",
		Finish_Date:   "2023-06-07",
		Duration:      "Less than a day",
		Description:   "Quasi, iusto autem voluptas, facilis quidem aliquid sed harum provident iure nobis eaque sin accusantium excepturi consequatur, amet totam magni blanditiis. Voluptatibus natus placeat, maiores voluptates distinctio quia. Saepe porro maxime iste maiores voluptatem deserunt alias, debitis dolore odit eaque. Cupiditate architecto eaque vero debitis velit unde, cum haru consequatur iste. Neque molestias ducimus temporibus ex labore tempore magni, provident maiores? Obcaecati, placeat aliquid, officiis nulla voluptatem asperiores, ipsa unde a vel magni facere inventore necessitatibus.",
		Icon_A:        `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`,
		Icon_B:        `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`,
		Icon_C:        `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`,
		Icon_D:        `<i class="fa-brands fa-react fa-lg fa-fw"></i>`,
		Label_A:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C:       `<div class="d-flex flex-row align-items-center gap-1">
            				<i class="fa-brands fa-golang fa-lg fa-fw"></i>
            				<p class="m-0">Go</p>
        				</div>
						`,
		Label_D:       `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
}

func main() {
	e := echo.New()
	e.Static("/Public", "Public")
	e.GET("/", Main_Page)
	e.GET("/Project-Detail/:id", Static_Project_Detail)
	e.POST("/", Project_Form_Value)
	e.POST("/Delete-Project/:id", Delete_Project)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func Main_Page(c echo.Context) error {
	var tmpl, err = template.ParseFiles("Views/Index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	Projects := map[string]interface{}{
		"Projects": Static_Project_Data,
	}
	return tmpl.Execute(c.Response(), Projects)
}

func Static_Project_Detail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var Project_Detail = Project{}
	for i, data := range Static_Project_Data {
		if id == i {
			Project_Detail = Project{
				Project_Title: data.Project_Title,
				Start_Date:    data.Start_Date,
				Finish_Date:   data.Finish_Date,
				Duration:      data.Duration,
				Description:   data.Description,
				Label_A:       data.Label_A,
				Label_B:       data.Label_B,
				Label_C:       data.Label_C,
				Label_D:       data.Label_D,
			}
		}
	}
	data := map[string]interface{}{
		"Project":        Project_Detail,
		"Date_Formation": Date_Formation,
	}
	tmpl, err := template.New("Project-Detail.html").Funcs(template.FuncMap{"Date_Formation": Date_Formation}).ParseFiles("Views/Project-Detail.html")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}
	err = tmpl.Execute(c.Response(), data)
	if err != nil {
		fmt.Println("Template execution error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	return nil
}

func Date_Formation(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "Invalid date format"
	}
	return t.Format("Jan 2, 2006")
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

func Fa_I(Input_Value string) string {
	if Input_Value == "Toggle_A" {
		return `<i class="fa-brands fa-square-js fa-lg fa-fw"></i>`
	} else if Input_Value == "Toggle_B" {
		return `<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>`
	} else if Input_Value == "Toggle_C" {
		return `<i class="fa-brands fa-golang fa-lg fa-fw"></i>`
	} else if Input_Value == "Toggle_D" {
		return `<i class="fa-brands fa-react fa-lg fa-fw"></i>`
	} else {
		return ""
	}

}

func Fa_I_Labels(Input_Value string) string {
	if Input_Value == "Toggle_A" {
		return `
		<div class="d-flex flex-row align-items-center gap-1">
			<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
			<p class="m-0">Javascript</p>
		</div>
		`
	} else if Input_Value == "Toggle_B" {
		return `
		<div class="d-flex flex-row align-items-center gap-1">
            <i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
            <p class="m-0">Bootstrap</p>
        </div>
		`
	} else if Input_Value == "Toggle_C" {
		return `
		<div class="d-flex flex-row align-items-center gap-1">
            <i class="fa-brands fa-golang fa-lg fa-fw"></i>
            <p class="m-0">Go</p>
        </div>
		`
	} else if Input_Value == "Toggle_D" {
		return `
		<div class="d-flex flex-row align-items-center gap-1">
			<i class="fa-brands fa-react fa-lg fa-fw"></i>
			<p class="m-0">React</p>
		</div>
		`
	} else {
		return ""
	}
}

func Project_Form_Value(c echo.Context) error {
	Project_Title := c.FormValue("Project_Title")
	Start_Date := c.FormValue("Start_Date")
	Finish_Date := c.FormValue("Finish_Date")
	Description := c.FormValue("Description")
	Toggle_A := c.FormValue("Toggle_A")
	Toggle_B := c.FormValue("Toggle_B")
	Toggle_C := c.FormValue("Toggle_C")
	Toggle_D := c.FormValue("Toggle_D")
	Icon_A := Fa_I(Toggle_A)
	Icon_B := Fa_I(Toggle_B)
	Icon_C := Fa_I(Toggle_C)
	Icon_D := Fa_I(Toggle_D)
	Label_A := Fa_I_Labels(Toggle_A)
	Label_B := Fa_I_Labels(Toggle_B)
	Label_C := Fa_I_Labels(Toggle_C)
	Label_D := Fa_I_Labels(Toggle_D)
	Start_Date_Format, _ := time.Parse("2006-01-02", Start_Date)
	Finish_Date_Format, _ := time.Parse("2006-01-02", Finish_Date)
	Duration := Finish_Date_Format.Sub(Start_Date_Format)
	Duration_Format := Duration_Formatting(Duration)
	var Render_New_Project = Project{
		Project_Title: Project_Title,
		Start_Date:    Start_Date,
		Finish_Date:   Finish_Date,
		Description:   Description,
		Icon_A:        Icon_A,
		Icon_B:        Icon_B,
		Icon_C:        Icon_C,
		Icon_D:        Icon_D,
		Label_A:       Label_A,
		Label_B:       Label_B,
		Label_C:       Label_C,
		Label_D:       Label_D,
		Duration:      Duration_Format,
	}
	fmt.Println(Render_New_Project)
	Static_Project_Data = append(Static_Project_Data, Render_New_Project)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Delete_Project(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Static_Project_Data = append(Static_Project_Data[:id], Static_Project_Data[id+1:]... )
	return c.Redirect(http.StatusMovedPermanently, "/")
}
