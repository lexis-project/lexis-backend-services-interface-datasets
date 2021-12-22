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

func (u *DatasetManager) CompressToZip (ctx context.Context, params staging.CompressToZipParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("compress", callTime)
	etype := staging.NewCompressToZipServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
	if err != nil {
		u.monit.APIHitDone("compress", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.compressapi+"/zip", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("compress", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("compress", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CompressToZipCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
		return staging.NewCompressToZipCreated().WithPayload(&ret)
	}
	//Unexpected response
	if (res.StatusCode == staging.CompressToZipBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCompressToZipBadRequest().WithPayload(response)
	} else if (res.StatusCode == staging.CompressToZipUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCompressToZipUnauthorized().WithPayload(response)
	} else if (res.StatusCode == staging.CompressToZipBadGatewayCode) { // 502
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCompressToZipBadGateway().WithPayload(response)
	} else if (res.StatusCode == staging.CompressToZipServiceUnavailableCode) { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCompressToZipServiceUnavailable().WithPayload(response)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCompressToZipServiceUnavailable().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) CheckCompressToZipStatus (ctx context.Context, params staging.CheckCompressToZipStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("compress", callTime)
	etype := staging.NewCheckCompressToZipStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.compressapi+"/zip/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("compress", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
//        req.Header.Set("Content-Type", "application/json")
//        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("compress", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckCompressToZipStatusOKCode) {
		ret := staging.CheckCompressToZipStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
		return staging.NewCheckCompressToZipStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == staging.CheckCompressToZipStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == staging.CheckCompressToZipStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == staging.CheckCompressToZipStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == staging.CheckCompressToZipStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == staging.CheckCompressToZipStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("compress", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("compress", callTime)
                return staging.NewCheckCompressToZipStatusInternalServerError().WithPayload(&returnValError)
        }
}
