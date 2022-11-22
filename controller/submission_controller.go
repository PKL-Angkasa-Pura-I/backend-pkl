package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func (ce *EchoController) GetAllSubmissionController(c echo.Context) error {

	submissions := ce.Svc.GetAllSubmissionService()

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": submissions,
	})
}

func (ce *EchoController) GetOneSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": res,
	})
}

func (ce *EchoController) GetFileSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	}

	name_file := fmt.Sprintf("%s%s", res.CodeSubmission, ".pdf")

	return c.Attachment(res.SubmissionPathFile, name_file)
}
