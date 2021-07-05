package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/disturb16/go-sqlite-service/internal/api/v1/dto"
	"github.com/disturb16/go-sqlite-service/internal/persons/entity"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	"github.com/sanservices/apilogger/v2"
)

func (h Handler) persons(c echo.Context) error {
	limitParam := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitParam)
	ctx := c.Request().Context()

	pp, err := h.service.Persons(ctx, limit)
	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatServiceOutput, err)
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	return helper.RespondOk(c, pp)
}

func (h Handler) person(c echo.Context) error {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	ctx := c.Request().Context()

	if err != nil || id < 1 {
		respErr := errors.New("parameter id is not valid")
		return helper.RespondError(c, http.StatusBadRequest, respErr)
	}

	p, err := h.service.Person(ctx, id)

	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	return helper.RespondOk(c, p)
}

func (h Handler) savePerson(c echo.Context) error {

	params := &dto.RegisterUserDto{}
	ctx := c.Request().Context()
	var err error

	err = helper.DecodeBody(c, &params.Body)
	if err != nil {
		apilogger.Error(c.Request().Context(), apilogger.LogCatUnmarshalReq, err)
		return helper.RespondError(c, http.StatusBadRequest, errParametersNotValid)
	}

	_, err = h.service.SavePerson(ctx, params.Body.Name, params.Body.Age)
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (h Handler) updatePerson(c echo.Context) error {

	ctx := c.Request().Context()
	params := &dto.UpdateUserDto{}
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil || id < 1 {
		respErr := errors.New("parameter id is not valid")
		return helper.RespondError(c, http.StatusBadRequest, respErr)
	}

	err = helper.DecodeBody(c, &params.Body)
	if err != nil {
		apilogger.Error(c.Request().Context(), apilogger.LogCatUnmarshalReq, err)
		return helper.RespondError(c, http.StatusBadRequest, errParametersNotValid)
	}

	p := entity.Person{
		ID:   id,
		Name: params.Body.Name,
		Age:  params.Body.Age,
	}

	err = h.service.UpdatePerson(ctx, p)
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
