package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Upload_Image(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		file, err := c.FormFile("Upload_Image")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		defer src.Close()

		temp, err := ioutil.TempFile("Uploads", "Image-*.png")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		defer temp.Close()

		if _, err = io.Copy(temp, src); err !=nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		data := temp.Name()
		File_Name := data[8:] //mengurangi digit string dari "uploads/" yang jumlahnya 8

		c.Set("File_Data", File_Name)

		return next(c)
	}
}
