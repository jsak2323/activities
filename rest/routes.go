package rest

import (
	"activities/helpers"
	"activities/service"
	"context"
	"flag"
	"net/http"
	"strconv"

	// "bitbucket.org/ruparupa/go-src/helpers"
	"github.com/labstack/echo"
)

type Rest struct {
	svc service.Service
}

func New(svcs service.Service,
) *Rest {
	return &Rest{
		svc: svcs,
	}

}

func (re *Rest) Routes(e *echo.Echo) *echo.Echo {
	// const db = *sql.DB
	e.GET("/activities/all", re.GetAllActivities)
	e.GET("/activities/:id", re.GetDetailActivities)
	e.POST("/activities", re.PostActivity)
	e.PUT("/activities", re.PutActivity)
	e.DELETE("/activities/:id", re.DeleteActivity)
	return e
}

func (re *Rest) GetAllActivities(e echo.Context) error {
	var (
		response = Response{}
		ctx      = e.Request().Context()
	)
	if ctx == nil {
		ctx = context.Background()
	}

	acts, err := re.svc.SelectAllActivities(ctx)
	if err != nil {
		title := "[Failed call service select all activities]"
		writeLogError(title, err.Error(), "error")
		response.Errors = Errors{
			Code:     401,
			Title:    title,
			Messages: []string{err.Error()},
		}
		response.Messages = append(response.Messages, "error")
		return e.JSON(http.StatusInternalServerError, response)
	}

	data := make([]*ResponseActivity, 0, len(acts))
	for _, v := range acts {
		data = append(data, &ResponseActivity{
			ID:           v.ID,
			ActivityType: "select-all-activities",
			responseActivityAttributes: responseActivityAttributes{
				Name:   v.Name,
				Act:    v.Act,
				Status: v.Status,
			},
		})
	}

	response.Data = data
	response.Messages = append(response.Messages, "success")
	return e.JSON(http.StatusOK, response)
}

func (re *Rest) GetDetailActivities(e echo.Context) error {
	var (
		id, _    = strconv.ParseInt(e.Param("id"), 10, 64)
		params   = service.Activity{ID: id}
		response = Response{}
		ctx      = e.Request().Context()
	)
	if ctx == nil {
		ctx = context.Background()
	}

	actt, err := re.svc.GetActivityByID(ctx, &params)
	if err != nil {
		title := "[Failed call service select activity by id]"
		writeLogError(title, err.Error(), "error")
		response.Errors = Errors{
			Code:     401,
			Title:    title,
			Messages: []string{err.Error()},
		}
		response.Messages = append(response.Messages, "error")
		return e.JSON(http.StatusInternalServerError, response)
	}

	data := &ResponseActivity{
		ID:           actt.ID,
		ActivityType: "select-activity-by-id",
		responseActivityAttributes: responseActivityAttributes{
			Name:   actt.Name,
			Act:    actt.Act,
			Status: actt.Status,
		}}

	response.Data = data
	response.Messages = append(response.Messages, "success")
	return e.JSON(http.StatusOK, response)
}

func (re *Rest) PostActivity(e echo.Context) error {
	var (
		response = Response{}
		params   RequestActivity
		ctx      = e.Request().Context()
	)
	if ctx == nil {
		ctx = context.Background()
	}

	e.Bind(&params)

	acts, err := re.svc.InsertActivity(ctx, &service.Activity{
		Name:   params.Name,
		Act:    params.Act,
		Status: params.Status,
	})
	if err != nil {
		title := "[Failed call service insert activity]"
		writeLogError(title, err.Error(), "error")
		response.Errors = Errors{
			Code:     401,
			Title:    title,
			Messages: []string{err.Error()},
		}
		response.Messages = append(response.Messages, "error")
		return e.JSON(http.StatusInternalServerError, response)
	}

	resp := &ResponseActivity{
		ID:           acts.ID,
		ActivityType: "insert-activities",
		responseActivityAttributes: responseActivityAttributes{
			Name:   params.Name,
			Act:    params.Act,
			Status: params.Status,
		},
	}

	response.Data = resp
	response.Messages = append(response.Messages, "success")
	return e.JSON(http.StatusOK, response)
}

func (re *Rest) PutActivity(e echo.Context) error {
	var (
		response = Response{}
		params   service.Activity
		ctx      = e.Request().Context()
	)

	if ctx == nil {
		ctx = context.Background()
	}
	_ = params

	e.Bind(&params)

	err := re.svc.UpdateActivity(ctx, &service.Activity{
		ID:     params.ID,
		Name:   params.Name,
		Act:    params.Act,
		Status: params.Status,
	})
	if err != nil {
		title := "[Failed call service update activity]"
		writeLogError(title, err.Error(), "error")
		response.Errors = Errors{
			Code:     401,
			Title:    title,
			Messages: []string{err.Error()},
		}
		response.Messages = append(response.Messages, "error")
		return e.JSON(http.StatusInternalServerError, response)
	}

	resp := &ResponseActivity{
		ID:           params.ID,
		ActivityType: "update-activities",
		responseActivityAttributes: responseActivityAttributes{
			Name:   params.Name,
			Act:    params.Act,
			Status: params.Status,
		},
	}

	response.Data = resp
	response.Messages = append(response.Messages, "success")
	return e.JSON(http.StatusOK, response)
}

func (re *Rest) DeleteActivity(e echo.Context) error {
	var (
		response = Response{}
		id, _    = strconv.ParseInt(e.Param("id"), 10, 64)
		params   = service.Activity{ID: id}
		ctx      = e.Request().Context()
	)
	if ctx == nil {
		ctx = context.Background()
	}
	e.Bind(&params)

	err := re.svc.DeleteActivity(ctx, &service.Activity{
		ID: params.ID,
	})
	if err != nil {
		title := "[Failed call service delete activity]"
		writeLogError(title, err.Error(), "error")
		response.Errors = Errors{
			Code:     401,
			Title:    title,
			Messages: []string{err.Error()},
		}
		response.Messages = append(response.Messages, err.Error())
		return e.JSON(http.StatusInternalServerError, response)
	}

	// response.Data = make(map[string]interface{})
	response.Data = "OK"
	response.Messages = append(response.Messages, "success")
	return e.JSON(http.StatusOK, response)
}

func writeLogError(title, msgError, logFileName string) {
	flag.Parse()
	projectFolder := flag.Lookup("folder").Value.(flag.Getter).Get().(string)

	helpers.WriteToLogfile(projectFolder+"logs/"+logFileName+".log", title, ", Error: "+msgError)
}
