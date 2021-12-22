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

func (u *DatasetManager) CheckPIDStatus (ctx context.Context, 
		params data_set_management.CheckPIDStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("replicate", callTime)   
	etype := data_set_management.NewCheckPIDStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/pid/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CheckPIDStatusOKCode) {
		ret := models.DataReplication {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
		return data_set_management.NewCheckPIDStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.CheckPIDStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckPIDStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckPIDStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckPIDStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckPIDStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckPIDStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckReplicateStatus (ctx context.Context, 
		params data_set_management.CheckReplicateStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("replicate", callTime)
	etype := data_set_management.NewCheckReplicateStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/replicate/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CheckReplicateStatusOKCode) {
		ret := models.DataReplication {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
		return data_set_management.NewCheckReplicateStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.CheckReplicateStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckReplicateStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckReplicateStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckReplicateStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckReplicateStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewCheckReplicateStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) PID (ctx context.Context, params data_set_management.PIDParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("replicate", callTime)
	etype := data_set_management.NewPIDInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.steeringapi+"/pid/assign", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.PIDCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
		return data_set_management.NewPIDCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.PIDUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.PIDBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PIDForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PIDNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PIDRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PIDTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PIDInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewPIDInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Replicate (ctx context.Context, params data_set_management.ReplicateParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("replicate", callTime)
	etype := data_set_management.NewReplicateInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.steeringapi+"/replicate", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("replicate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.ReplicateCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
		return data_set_management.NewReplicateCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.ReplicateUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.ReplicateBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.ReplicateForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.ReplicateNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.ReplicateRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.ReplicateTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.ReplicateInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("replicate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("replicate", callTime)
                return data_set_management.NewReplicateInternalServerError().WithPayload(&returnValError)
        }
}
