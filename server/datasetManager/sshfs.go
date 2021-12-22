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

func (u *DatasetManager) CreateSSHFSExport (ctx context.Context, params data_set_management.CreateSSHFSExportParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("sshfs", callTime)
	etype := data_set_management.NewCreateSSHFSExportServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.sshfsapi+"/sshfsexport", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CreateSSHFSExportCreatedCode) {
		rvalue := data_set_management.CreateSSHFSExportCreatedBody{}
                err := json.Unmarshal(body, &rvalue)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
		return data_set_management.NewCreateSSHFSExportCreated().WithPayload(&rvalue)
	}
	//Unexpected response
        if res.StatusCode == data_set_management.CreateSSHFSExportUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewCreateSSHFSExportUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CreateSSHFSExportBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewCreateSSHFSExportBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CreateSSHFSExportForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewCreateSSHFSExportForbidden().WithPayload(response)

          } else if res.StatusCode == data_set_management.CreateSSHFSExportServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewCreateSSHFSExportServiceUnavailable().WithPayload(response)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewCreateSSHFSExportServiceUnavailable().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) DeleteSSHFSExport (ctx context.Context, params data_set_management.DeleteSSHFSExportParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("sshfs", callTime)
	etype := data_set_management.NewDeleteSSHFSExportServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.sshfsapi+"/sshfsexport", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("sshfs", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.DeleteSSHFSExportNoContentCode) {
		u.monit.APIHitDone("sshfs", callTime)
		return data_set_management.NewDeleteSSHFSExportNoContent()
	}
	//Unexpected response
        if res.StatusCode == data_set_management.DeleteSSHFSExportUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteSSHFSExportBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DeleteSSHFSExportForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DeleteSSHFSExportNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportNotFound().WithPayload(response)
          } else if res.StatusCode == data_set_management.DeleteSSHFSExportServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("sshfs", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportServiceUnavailable().WithPayload(response)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("sshfs", callTime)
                return data_set_management.NewDeleteSSHFSExportServiceUnavailable().WithPayload(&returnValError)
        }
}
