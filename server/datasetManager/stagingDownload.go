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

func (u *DatasetManager) PostStagingDownload (ctx context.Context, params data_set_management.PostStagingDownloadParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("stagingDownload", callTime)
	etype := data_set_management.NewPostStagingDownloadServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
	if err != nil {
		u.monit.APIHitDone("stagingDownload", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/staging/download", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("stagingDownload", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("stagingDownload", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	if res.StatusCode == data_set_management.PostStagingDownloadOKCode {
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadOK().WithPayload(res.Body)
	}
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	//Unexpected response
	if res.StatusCode == data_set_management.PostStagingDownloadUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("stagingDownload", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.PostStagingDownloadBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("stagingDownload", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadBadRequest().WithPayload(response)
          } else if res.StatusCode == data_set_management.PostStagingDownloadForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
                        u.monit.APIHitDone("stagingDownload", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PostStagingDownloadBadGatewayCode { // 502
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("stagingDownload", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadBadGateway().WithPayload(response)
	  } else if res.StatusCode == data_set_management.PostStagingDownloadServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("stagingDownload", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadServiceUnavailable().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("stagingDownload", callTime)
                return data_set_management.NewPostStagingDownloadServiceUnavailable().WithPayload(&returnValError)
        }
}

