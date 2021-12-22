package datasetManager
  
import (
//        l "gitlab.com/cyclops-utilities/logging"
        "context"
        "net/http"
        "bytes"
	//"fmt"
        "io/ioutil"
	"strconv"
	"time"
        "github.com/segmentio/encoding/json"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )

func (u *DatasetManager) AddSupportUser (ctx context.Context, params data_set_management.AddSupportUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("support", callTime)
       etype := data_set_management.NewAddSupportUserInternalServerError()
        b, err := json.Marshal(params.Username)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/support", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.AddSupportUserCreatedCode) {
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserCreated()
        }
        //Unexpected response
       if res.StatusCode == data_set_management.AddSupportUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteSupportUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddSupportUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddSupportUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewAddSupportUserBadGateway().WithPayload(&returnValError)
        }

}

func (u *DatasetManager) DeleteSupportUser (ctx context.Context, params data_set_management.DeleteSupportUserParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("support", callTime)    
	etype := data_set_management.NewDeleteSupportUserInternalServerError()
        b, err := json.Marshal(params.Username)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/support", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.DeleteSupportUserNoContentCode) {
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserNoContent()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.DeleteSupportUserUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteSupportUserBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteSupportUserInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteSupportUserBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
                return data_set_management.NewDeleteSupportUserBadGateway().WithPayload(&returnValError)
        }
}

