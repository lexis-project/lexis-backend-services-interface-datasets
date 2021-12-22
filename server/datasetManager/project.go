package datasetManager

import (
//        l "gitlab.com/cyclops-utilities/logging"
        "context"
        "net/http"
        "bytes"
        "io/ioutil"
        "github.com/segmentio/encoding/json"
	"strconv"
	"time"
	//"fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )

func (u *DatasetManager) CreateProject (ctx context.Context, params data_set_management.CreateProjectParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)
	etype := data_set_management.NewCreateProjectInternalServerError()
        b, err := json.Marshal(params.ProjectInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/project", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CreateProjectCreatedCode) {
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectCreated()
        }
        //Unexpected response
     if res.StatusCode == data_set_management.CreateProjectUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CreateProjectBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.CreateProjectForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectForbidden().WithPayload(response)

        } else if res.StatusCode == data_set_management.CreateProjectConflictCode { // 409
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectConflict().WithPayload(response)
        } else if res.StatusCode == data_set_management.CreateProjectInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.CreateProjectBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewCreateProjectBadGateway().WithPayload(&returnValError)
        }

}

func (u *DatasetManager) DeleteProject (ctx context.Context, params data_set_management.DeleteProjectParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)   
	etype := data_set_management.NewDeleteProjectInternalServerError()
        b, err := json.Marshal(params.ProjectInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/project", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.DeleteProjectNoContentCode) {
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectNoContent()
        }
        //Unexpected response
        if res.StatusCode == data_set_management.DeleteProjectUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDeleteProjectUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteProjectBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectBadGateway().WithPayload(&returnValError)
        }

}

func (u *DatasetManager) AddProjectAdmin (ctx context.Context, params data_set_management.AddProjectAdminParams)  middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)   
	etype := data_set_management.NewAddProjectAdminInternalServerError()
        b, err := json.Marshal(params.ProjectAndUserInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/project/admin", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.AddProjectAdminCreatedCode) { // 201
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminCreated()
        }
        //Unexpected response
        if res.StatusCode == data_set_management.AddProjectAdminUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.AddProjectAdminBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminBadRequest().WithPayload(response)
	 } else if res.StatusCode == data_set_management.AddProjectAdminConflictCode { // 409
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminConflict().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddProjectAdminInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectAdminBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectAdminBadGateway().WithPayload(&returnValError)
        }

}

func (u *DatasetManager) DeleteProjectAdmin (ctx context.Context, params data_set_management.DeleteProjectAdminParams)  middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)   
	etype := data_set_management.NewDeleteProjectAdminInternalServerError() // 500
        b, err := json.Marshal(params.ProjectAndUserInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/project/admin", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.DeleteProjectAdminNoContentCode) {
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminNoContent()
        }
        //Unexpected response
        if res.StatusCode == data_set_management.DeleteProjectAdminUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteProjectAdminBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminBadRequest().WithPayload(response)
	} else if res.StatusCode == data_set_management.DeleteProjectAdminConflictCode { // 409
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminConflict().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectAdminInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectAdminBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectAdminBadGateway().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) AddProjectUser (ctx context.Context, params data_set_management.AddProjectUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)   
	etype := data_set_management.NewAddProjectUserInternalServerError()
        b, err := json.Marshal(params.ProjectAndUserInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/project/user", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.AddProjectUserCreatedCode) {
                rvalue := models.ErrorResponse{}
                err := json.Unmarshal(body, &rvalue)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserCreated().WithPayload(&rvalue)
        }
        //Unexpected response
       if res.StatusCode == data_set_management.AddProjectUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.AddProjectUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserBadRequest().WithPayload(response)
	} else if res.StatusCode == data_set_management.AddProjectUserConflictCode { // 409
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserConflict().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddProjectUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddProjectUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewAddProjectUserBadGateway().WithPayload(&returnValError)
        }


}

func (u *DatasetManager) DeleteProjectUser (ctx context.Context, params data_set_management.DeleteProjectUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("project", callTime)   
	etype := data_set_management.NewDeleteProjectUserInternalServerError()
        b, err := json.Marshal(params.ProjectAndUserInfo)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/project/admin", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("project", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.DeleteProjectUserNoContentCode) {
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserNoContent()
        }
        //Unexpected response
        if res.StatusCode == data_set_management.DeleteProjectUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteProjectUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserBadRequest().WithPayload(response)
	} else if res.StatusCode == data_set_management.DeleteProjectUserConflictCode { // 409
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserConflict().WithPayload(response)	
        } else if res.StatusCode == data_set_management.DeleteProjectUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("project", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteProjectUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("project", callTime)
                return data_set_management.NewDeleteProjectUserBadGateway().WithPayload(&returnValError)
        }

}

