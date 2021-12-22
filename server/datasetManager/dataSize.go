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
//      "fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )

func (u *DatasetManager) CheckSizeStatus (ctx context.Context, params data_set_management.CheckSizeStatusParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("dataSize", callTime)
        etype := data_set_management.NewCheckSizeStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.datasizeapi+"/data/size/"+params.RequestID.String(), nil)
        if err != nil {
                u.monit.APIHitDone("dataSize", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        res, err := u.httpclient.Do(req)
        if err != nil {
                u.monit.APIHitDone("dataSize", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckSizeStatusOKCode) {
                ret := models.DataSize {}
                err := json.Unmarshal(body, &ret)
                if err != nil {
                        u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusOK().WithPayload(&ret)
        }
	//Unexpected response
	if (res.StatusCode == data_set_management.CheckSizeStatusBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusBadRequest().WithPayload(response)
	} else if (res.StatusCode == data_set_management.CheckSizeStatusUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusUnauthorized().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckSizeStatusForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckSizeStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckSizeStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckSizeStatusTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckSizeStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewCheckSizeStatusInternalServerError().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) Size (ctx context.Context, params data_set_management.SizeParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("dataSize", callTime)
	etype := data_set_management.NewSizeInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("dataSize", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.datasizeapi+"/data/size", 
		bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("dataSize", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("dataSize", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.SizeCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
		return data_set_management.NewSizeCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.SizeUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.SizeBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.SizeForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.SizeNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.SizeRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.SizeTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.SizeInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("dataSize", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("dataSize", callTime)
                return data_set_management.NewSizeInternalServerError().WithPayload(&returnValError)
        }
}
