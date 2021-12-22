package datasetManager
  
import (
//        l "gitlab.com/cyclops-utilities/logging"
        "context"
        "net/http"
        //"bytes"
        "io/ioutil"
        "github.com/segmentio/encoding/json"
	"strconv"
	"time"
//	"fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/staging"
        )

 
func (u *DatasetManager) CheckCloudNFSExportAddStatus (ctx context.Context, params staging.CheckCloudNFSExportAddStatusParams) middleware.Responder {
	callTime := time.Now()
	u.monit.APIHit("cloudnfs", callTime)
	etype := staging.NewCheckCloudNFSExportAddStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/cloud/add/"+params.Param, nil)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckCloudNFSExportAddStatusOKCode) {
		ret := staging.CheckCloudNFSExportAddStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
		return staging.NewCheckCloudNFSExportAddStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if (res.StatusCode == staging.CheckCloudNFSExportAddStatusBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusBadRequest().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportAddStatusUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return staging.NewCheckCloudNFSExportAddStatusUnauthorized().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportAddStatusForbiddenCode) { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusForbidden().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportAddStatusNotFoundCode) { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusNotFound().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportAddStatusRequestURITooLongCode) { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusRequestURITooLong().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportAddStatusInternalServerErrorCode) { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusInternalServerError().WithPayload(response)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusInternalServerError().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) CheckCloudNFSExportRemoveStatus (ctx context.Context, params staging.CheckCloudNFSExportRemoveStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("cloudnfs", callTime)
	etype := staging.NewCheckCloudNFSExportRemoveStatusInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/cloud/remove/"+params.Param, nil)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusOKCode) {
		ret := staging.CheckCloudNFSExportRemoveStatusOKBody {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
		return staging.NewCheckCloudNFSExportRemoveStatusOK().WithPayload(&ret)
	}
	//Unexpected response
	if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportRemoveStatusBadRequest().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportRemoveStatusUnauthorized().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusNotFoundCode) { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportRemoveStatusNotFound().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusRequestURITooLongCode) { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportRemoveStatusRequestURITooLong().WithPayload(response)
	} else if (res.StatusCode == staging.CheckCloudNFSExportRemoveStatusInternalServerErrorCode) { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportRemoveStatusInternalServerError().WithPayload(response)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCheckCloudNFSExportAddStatusInternalServerError().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) CloudNFSExportAdd (ctx context.Context, params staging.CloudNFSExportAddParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("cloudnfs", callTime)
	etype := staging.NewCloudNFSExportAddInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/cloud/add/"+params.Param, nil)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CloudNFSExportAddCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
		return staging.NewCloudNFSExportAddCreated().WithPayload(&ret)
	}
	//Unexpected response
	if (res.StatusCode == staging.CloudNFSExportAddBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddBadRequest().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportAddUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddUnauthorized().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportAddForbiddenCode) { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddForbidden().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportAddRequestURITooLongCode) { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddRequestURITooLong().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportAddInternalServerErrorCode) { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddInternalServerError().WithPayload(response)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportAddInternalServerError().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) CloudNFSExportRemove (ctx context.Context, params staging.CloudNFSExportRemoveParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("cloudnfs", callTime)
	etype := staging.NewCloudNFSExportRemoveInternalServerError()
        req, err := http.NewRequest (http.MethodGet, u.steeringapi+"/cloud/remove/"+params.Param, nil)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("cloudnfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == staging.CloudNFSExportRemoveCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
		return staging.NewCloudNFSExportRemoveCreated().WithPayload(&ret)
	}
	//Unexpected response
	if (res.StatusCode == staging.CloudNFSExportRemoveBadRequestCode) { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportRemoveBadRequest().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportRemoveUnauthorizedCode) { // 401
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportRemoveUnauthorized().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportRemoveRequestURITooLongCode) { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportRemoveRequestURITooLong().WithPayload(response)
	} else if (res.StatusCode == staging.CloudNFSExportRemoveInternalServerErrorCode) { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("cloudnfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportRemoveInternalServerError().WithPayload(response)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("cloudnfs", callTime)
                return staging.NewCloudNFSExportRemoveInternalServerError().WithPayload(&returnValError)
	}
}
