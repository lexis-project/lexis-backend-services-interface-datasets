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
//	"fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )


func (u *DatasetManager) CreateUser (ctx context.Context, params data_set_management.CreateUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("user", callTime) 
	etype := data_set_management.NewCreateUserInternalServerError()
	b, err := json.Marshal(params.UserInfo)
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/user", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CreateUserCreatedCode) {
		u.monit.APIHitDone("user", callTime)
		return data_set_management.NewCreateUserCreated()
	}
	//Unexpected response
        if res.StatusCode == data_set_management.CreateUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CreateUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CreateUserForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserForbidden().WithPayload(response)
        } else if res.StatusCode == data_set_management.CreateUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.CreateUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserBadGateway().WithPayload(&returnValError)
          } else if res.StatusCode == data_set_management.CreateUserServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserServiceUnavailable().WithPayload(response)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewCreateUserBadGateway().WithPayload(&returnValError)
        }

}

func (u *DatasetManager) DeleteUser (ctx context.Context, params data_set_management.DeleteUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("user", callTime)
	etype := data_set_management.NewDeleteUserInternalServerError()
        b, err := json.Marshal(params.User)
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/user", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("user", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.DeleteUserNoContentCode) {
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserNoContent()
        }
        //Unexpected response
        if res.StatusCode == data_set_management.DeleteUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("user", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("user", callTime)
                return data_set_management.NewDeleteUserBadGateway().WithPayload(&returnValError)
        }

}

