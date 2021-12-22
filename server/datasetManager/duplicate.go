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
//        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/staging"
        )

func (u *DatasetManager) CheckDuplicationStatus (ctx context.Context, 
		params staging.CheckDuplicationStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("duplicate", callTime)
	etype := staging.NewCheckDuplicationStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/duplicate/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("duplicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("duplicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckDuplicationStatusOKCode) {
		ret := staging.CheckDuplicationStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
		return staging.NewCheckDuplicationStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.CheckDuplicationStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.CheckDuplicationStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDuplicationStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDuplicationStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDuplicationStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewCheckDuplicationStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Duplicate (ctx context.Context, params staging.DuplicateParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("duplicate", callTime)
	etype := staging.NewDuplicateInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("duplicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.steeringapi+"/duplicate", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("duplicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("duplicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.DuplicateCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
		return staging.NewDuplicateCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.DuplicateUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.DuplicateBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.DuplicateForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateForbidden().WithPayload(response)
	  } else if res.StatusCode == staging.DuplicateNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.DuplicateRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.DuplicateTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == staging.DuplicateInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("duplicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("duplicate", callTime)
                return staging.NewDuplicateInternalServerError().WithPayload(&returnValError)
        }
}

