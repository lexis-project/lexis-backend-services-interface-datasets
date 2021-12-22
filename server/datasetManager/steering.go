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
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/staging"
        )

func (u *DatasetManager) CheckDeletionStatus (ctx context.Context, params staging.CheckDeletionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("steering", callTime)
	etype := staging.NewCheckDeletionStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/delete/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckDeletionStatusOKCode) {
		ret := staging.CheckDeletionStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
		return staging.NewCheckDeletionStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.CheckDeletionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.CheckDeletionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDeletionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDeletionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.CheckDeletionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckDeletionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckStageStatus (ctx context.Context, params staging.CheckStageStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("steering", callTime)
	etype := staging.NewCheckStageStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/stage/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckStageStatusOKCode) {
		ret := staging.CheckStageStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
		return staging.NewCheckStageStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.CheckStageStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.CheckStageStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.CheckStageStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.CheckStageStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.CheckStageStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewCheckStageStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Delete (ctx context.Context, params staging.DeleteParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("steering", callTime)
	etype := staging.NewDeleteInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.steeringapi+"/delete", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.DeleteCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
		return staging.NewDeleteCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.DeleteUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return staging.NewDeleteUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.DeleteBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return staging.NewDeleteBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.DeleteForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return staging.NewDeleteForbidden().WithPayload(response)
	  } else if res.StatusCode == staging.DeleteNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewDeleteNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.DeleteRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewDeleteRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.DeleteTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewDeleteTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == staging.DeleteInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewDeleteInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewDeleteInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Stage (ctx context.Context, params staging.StageParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("steering", callTime)
	etype := staging.NewStageInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.steeringapi+"/stage", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.StageCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
		return staging.NewStageCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.StageUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.StageBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.StageForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageForbidden().WithPayload(response)
	  } else if res.StatusCode == staging.StageNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.StageRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.StageTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == staging.StageInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("steering", callTime)
                return staging.NewStageInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) StagingInfo (ctx context.Context, params staging.StagingInfoParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("steering", callTime)
	etype := staging.NewCheckStageStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/info", nil)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("steering", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.StagingInfoOKCode) {
		var ret [] string
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("steering", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("steering", callTime)
		return staging.NewStagingInfoOK().WithPayload(ret)
	}
	//Unexpected response
	return nil
}
