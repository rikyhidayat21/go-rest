package controller

import (
	"memetpaten/go-rest/helper"
	"memetpaten/go-rest/model/web"
	"memetpaten/go-rest/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// process decode request body
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	// process create category
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:  http.StatusCreated,
		Status: "OK",
		Data: categoryResponse,
	}
	
	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// process decode request body
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	// find category by id
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	// process update category
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: categoryResponse,
	}

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// find category by id
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// process delete category
	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
	}

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// find category by id
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// process find category
	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: categoryResponse,
	}

	// write response
	helper.WriteToResponseBody(w, webResponse)	
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Param) {
	// because find all no need params, so we just send response
	// process find all category
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: categoryResponses,
	}

	// write response
	helper.WriteToResponseBody(w, webResponse)
}