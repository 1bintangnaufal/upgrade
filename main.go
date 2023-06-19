package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
	connection "upgrade/Connection"
	middleware "upgrade/Middleware"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	ID     int
	Author string

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

type User struct {
	ID         int
	First_Name string
	Last_Name  string
	Email      string
	Password   string
}

type Data_Session struct {
	Login_State bool
	First_Name  string
}

var User_Session = Data_Session{}

func main() {
	connection.Database_Connect()

	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.Static("/Public", "Public")
	e.Static("/Uploads", "Uploads")

	e.GET("/", Main_Page)
	e.GET("/Project-Detail/:id", Project_Detail)
	e.GET("/Edit-Project/:id", Edit_Project)

	e.GET("/Login", Logging_in)
	e.GET("/Register", Registration)

	e.POST("/", middleware.Upload_Image(Project_Form_Value))
	e.POST("/Delete-Project/:id", Delete_Project)
	e.POST("/Edit-Project/:id", middleware.Upload_Image(Save_Changes))

	e.POST("/Login", Login)
	e.POST("/Register", Register)

	e.POST("/Logout", Logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func Main_Page(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["Login_State"] == true {
		User_Session.Login_State = sess.Values["Login_State"].(bool)
		User_Session.First_Name = sess.Values["First_Name"].(string)
		AuthorID := sess.Values["ID"].(int)

		data, _ := connection.Conn.Query(context.Background(), "SELECT tb_project.id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image, tb_user.first_name as author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_project.author_id = $1 ORDER BY tb_project.id", AuthorID)

		var result []Project

		for data.Next() {
			var each = Project{}
			err := data.Scan(&each.ID, &each.Project_Title, &each.Start_Date, &each.Finish_Date, &each.Description, &each.Toggle_A, &each.Toggle_B, &each.Toggle_C, &each.Toggle_D, &each.Image, &each.Author)

			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			}

			each.Duration = each.Finish_Date.Sub(each.Start_Date)
			each.Formatted_Duration = Duration_Formatting(each.Duration)

			result = append(result, each)
		}

		Data := map[string]interface{}{
			"Projects":      result,
			"Flash_Status":  sess.Values["Status"],
			"Flash_Message": sess.Values["Message"],
			"Data_Session":  User_Session,
		}

		delete(sess.Values, "Message")
		delete(sess.Values, "Status")
		sess.Save(c.Request(), c.Response())

		var tmpl, err = template.ParseFiles("Views/Index.html")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		Map_Duration_Formatting := template.FuncMap{
			"Duration_Formatting": Duration_Formatting,
		}

		tmpl = tmpl.Funcs(Map_Duration_Formatting)

		return tmpl.Execute(c.Response(), Data)

	} else {
		User_Session.Login_State = sess.Values["Login_State"] == false
		User_Session.First_Name = ""

		data, _ := connection.Conn.Query(context.Background(), "SELECT tb_project.id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image, tb_user.first_name as author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id ORDER BY tb_project.id")

		var result []Project

		for data.Next() {
			var each = Project{}
			err := data.Scan(&each.ID, &each.Project_Title, &each.Start_Date, &each.Finish_Date, &each.Description, &each.Toggle_A, &each.Toggle_B, &each.Toggle_C, &each.Toggle_D, &each.Image, &each.Author)

			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			}

			each.Duration = each.Finish_Date.Sub(each.Start_Date)
			each.Formatted_Duration = Duration_Formatting(each.Duration)

			result = append(result, each)
		}

		Data := map[string]interface{}{
			"Projects":      result,
			"Flash_Status":  sess.Values["Status"],
			"Flash_Message": sess.Values["Message"],
			"Data_Session":  User_Session,
		}

		delete(sess.Values, "Message")
		delete(sess.Values, "Status")
		sess.Save(c.Request(), c.Response())

		var tmpl, err = template.ParseFiles("Views/Index.html")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		Map_Duration_Formatting := template.FuncMap{
			"Duration_Formatting": Duration_Formatting,
		}

		tmpl = tmpl.Funcs(Map_Duration_Formatting)

		return tmpl.Execute(c.Response(), Data)
	}
}

func Project_Detail(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["Login_State"] == true {
		User_Session.Login_State = sess.Values["Login_State"].(bool)
		User_Session.First_Name = sess.Values["First_Name"].(string)
	} else {
		User_Session.Login_State = sess.Values["Login_State"] == false
		User_Session.First_Name = ""
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var Project_Detail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_project.id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image, tb_user.first_name as author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_project.id=$1", id).Scan(&Project_Detail.ID, &Project_Detail.Project_Title, &Project_Detail.Start_Date, &Project_Detail.Finish_Date, &Project_Detail.Description, &Project_Detail.Toggle_A, &Project_Detail.Toggle_B, &Project_Detail.Toggle_C, &Project_Detail.Toggle_D, &Project_Detail.Image, &Project_Detail.Author)

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
		"Data_Session":   User_Session,
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

// POST
func Project_Form_Value(c echo.Context) error {
	sess, _ := session.Get("session", c)
	Author := sess.Values["ID"].(int)

	Project_Title := c.FormValue("Project_Title")
	Start_Date := c.FormValue("Start_Date")
	Finish_Date := c.FormValue("Finish_Date")
	Description := c.FormValue("Description")

	Toggle_A := c.FormValue("Toggle_A") == "on"
	Toggle_B := c.FormValue("Toggle_B") == "on"
	Toggle_C := c.FormValue("Toggle_C") == "on"
	Toggle_D := c.FormValue("Toggle_D") == "on"

	Image := c.Get("File_Data").(string)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D, Image, Author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	fmt.Println(Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D)

	return c.Redirect(http.StatusFound, "/")
}

func Delete_Project(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusFound, "/")
}

func Edit_Project(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["Login_State"] == true {
		User_Session.Login_State = sess.Values["Login_State"].(bool)
		User_Session.First_Name = sess.Values["First_Name"].(string)
	} else {
		User_Session.Login_State = sess.Values["Login_State"] == false
		User_Session.First_Name = ""
	}

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
		"Data_Session":  User_Session,
	}

	tmpl, err_too := template.ParseFiles("Views/Edit-Project-Form.html")

	if err_too != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func Save_Changes(c echo.Context) error {
	sess, _ := session.Get("session", c)
	Author := sess.Values["ID"].(int)

	id, _ := strconv.Atoi(c.Param("id"))

	Project_Title := c.FormValue("Project_Title")
	Start_Date := c.FormValue("Start_Date")
	Finish_Date := c.FormValue("Finish_Date")
	Description := c.FormValue("Description")

	Toggle_A := c.FormValue("Toggle_A") == "on"
	Toggle_B := c.FormValue("Toggle_B") == "on"
	Toggle_C := c.FormValue("Toggle_C") == "on"
	Toggle_D := c.FormValue("Toggle_D") == "on"

	Image := c.Get("File_Data").(string)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_title=$1, start_date=$2, finish_date=$3, description=$4, toggle_a=$5, toggle_b=$6, toggle_c=$7, toggle_d=$8, image=$9, author_id=$10 WHERE id=$11", Project_Title, Start_Date, Finish_Date, Description, Toggle_A, Toggle_B, Toggle_C, Toggle_D, Image, Author, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	redirectURL := fmt.Sprintf("/?id=%d#ppc-container", id)
	return c.Redirect(http.StatusFound, redirectURL)
}

// GET
func Registration(c echo.Context) error {
	sess, _ := session.Get("session", c)

	Flash := map[string]interface{}{
		"Flash_Status":  sess.Values["Status"],
		"Flash_Message": sess.Values["Message"],
	}

	delete(sess.Values, "Message")
	delete(sess.Values, "Status")

	sess.Save(c.Request(), c.Response())
	
	var tmpl, err = template.ParseFiles("Views/Registration-Form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), Flash)
}

// POST
func Register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	First_Name := c.FormValue("First_Name")
	Last_Name := c.FormValue("Last_Name")
	Email := c.FormValue("Email")
	Password := c.FormValue("Password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(Password), 10)

	Existing_User := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT id, first_name, last_name, email, password FROM tb_user WHERE email=$1", Email).Scan(&Existing_User.ID, &Existing_User.First_Name, &Existing_User.Last_Name, &Existing_User.Email, &Existing_User.Password)
	if err == nil {
		return redirectWithMessage(c, "User with that email already exists. Please login or use another email", false, "/Register")
	}

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", First_Name, Last_Name, Email, passwordHash)
	if err != nil {
		redirectWithMessage(c, "Registration failed. Please try again", false, "/Register")
	}

	return redirectWithMessage(c, "Registration complete. Please login", true, "/Login")
}

// GET
func Logging_in(c echo.Context) error {
	sess, _ := session.Get("session", c)

	Flash := map[string]interface{}{
		"Flash_Status":  sess.Values["Status"],
		"Flash_Message": sess.Values["Message"],
	}

	delete(sess.Values, "Message")
	delete(sess.Values, "Status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("Views/Login-Form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), Flash)
}

// POST
func Login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	Email := c.FormValue("Email")
	Password := c.FormValue("Password")

	User := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT id, first_name, email, password FROM tb_user WHERE email = $1", Email).Scan(&User.ID, &User.First_Name, &User.Email, &User.Password)
	if err != nil {
		return redirectWithMessage(c, "Unregistered or incorrect email data", false, "/Login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(Password))
	if err != nil {
		return redirectWithMessage(c, "Incorrect Password", false, "/Login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 21600 //6 hours
	sess.Values["Message"] = "Logged in"
	sess.Values["Status"] = true
	sess.Values["ID"] = User.ID
	sess.Values["First_Name"] = User.First_Name
	sess.Values["Email"] = User.Email
	sess.Values["Login_State"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}

func redirectWithMessage(c echo.Context, Message string, Status bool, Path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["Message"] = Message
	sess.Values["Status"] = Status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, Path)
}
