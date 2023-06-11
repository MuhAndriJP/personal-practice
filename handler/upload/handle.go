package upload

import (
	"io"
	"os"

	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type Upload struct{}

func (u *Upload) Handle(c echo.Context) error {
	errors := []string{}
	file, err := c.FormFile("file")
	if err != nil {
		errors = append(errors, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		errors = append(errors, err.Error())
	}
	defer src.Close()

	// fileByte, _ := ioutil.ReadAll(src)
	// fileType := http.DetectContentType(fileByte)

	dstPath := "./handler/upload/file/" + file.Filename

	dst, err := os.Create(dstPath)
	if err != nil {
		errors = append(errors, err.Error())
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
			Data: map[string]interface{}{
				"error": errors,
			},
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
	})
}

func NewUpload() *Upload {
	return &Upload{}
}
