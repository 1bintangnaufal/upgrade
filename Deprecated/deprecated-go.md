``` go
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
		Label_A: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B: `<div class="d-flex flex-row align-items-center gap-1">
            				<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
            				<p class="m-0">Bootstrap</p>
        				</div>
						`,
		Label_C: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D: `<div class="d-flex flex-row align-items-center gap-1">
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
		Label_A: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D: `<div class="d-flex flex-row align-items-center gap-1">
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
		Label_A: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D: `<div class="d-flex flex-row align-items-center gap-1">
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
		Label_A: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-golang fa-lg fa-fw"></i>
							<p class="m-0">Go</p>
						</div>
						`,
		Label_D: `<div class="d-flex flex-row align-items-center gap-1">
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
		Label_A: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-square-js fa-lg fa-fw"></i>
							<p class="m-0">Javascript</p>
						</div>
						`,
		Label_B: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>
							<p class="m-0">Bootstrap</p>
						</div>
						`,
		Label_C: `<div class="d-flex flex-row align-items-center gap-1">
            				<i class="fa-brands fa-golang fa-lg fa-fw"></i>
            				<p class="m-0">Go</p>
        				</div>
						`,
		Label_D: `<div class="d-flex flex-row align-items-center gap-1">
							<i class="fa-brands fa-react fa-lg fa-fw"></i>
							<p class="m-0">React</p>
						</div>
						`,
	},
}