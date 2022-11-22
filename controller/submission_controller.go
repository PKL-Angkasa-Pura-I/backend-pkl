package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateSubmissionController(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	if file.Size >= 2200000 {
		return c.JSON(400, map[string]interface{}{
			"messages": "max file is 1 MB",
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if filetype != "application/pdf" {
		return c.JSON(400, map[string]interface{}{
			"messages": "file is not .pdf",
		})
	}

	defer src.Close()

	submission := model.Submission{}
	if err := c.Bind(&submission); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	filename := "../uploads/submission/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
	submission.SubmissionPathFile = filename

	res, err := ce.Svc.CreateSubmissionService(submission)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = os.WriteFile(filename, filebyte, 0777)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write new file",
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":        "success",
		"code_submission": res,
	})
}
