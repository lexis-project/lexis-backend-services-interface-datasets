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

func (u *DatasetManager) AddGridmapEntry (ctx context.Context, params data_set_management.AddGridmapEntryParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("gridmap", callTime)
	etype := data_set_management.NewAddGridmapEntryServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.gridmapapi+"/gridmap", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.AddGridmapEntryCreatedCode) {
		u.monit.APIHitDone("gridmap", callTime)
		return data_set_management.NewAddGridmapEntryCreated()
	}
	//Unexpected response
        if res.StatusCode == data_set_management.AddGridmapEntryUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewAddGridmapEntryUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.AddGridmapEntryForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewAddGridmapEntryForbidden().WithPayload(response)
        } else if res.StatusCode == data_set_management.AddGridmapEntryBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewAddGridmapEntryBadGateway().WithPayload(&returnValError)
          } else if res.StatusCode == data_set_management.AddGridmapEntryServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewAddGridmapEntryServiceUnavailable().WithPayload(response)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewAddGridmapEntryBadGateway().WithPayload(&returnValError)
        }
}


func (u *DatasetManager) RemoveGridmapEntry (ctx context.Context, params data_set_management.RemoveGridmapEntryParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("gridmap", callTime)
	etype := data_set_management.NewRemoveGridmapEntryServiceUnavailable()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.gridmapapi+"/gridmap", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("gridmap", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.RemoveGridmapEntryNoContentCode) {
		u.monit.APIHitDone("gridmap", callTime)
		return data_set_management.NewRemoveGridmapEntryNoContent()
	}
	//Unexpected response
        if res.StatusCode == data_set_management.RemoveGridmapEntryUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewRemoveGridmapEntryUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.RemoveGridmapEntryForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewRemoveGridmapEntryForbidden().WithPayload(response)
        } else if res.StatusCode == data_set_management.RemoveGridmapEntryBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewRemoveGridmapEntryBadGateway().WithPayload(&returnValError)
          } else if res.StatusCode == data_set_management.RemoveGridmapEntryServiceUnavailableCode { // 503
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("gridmap", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewRemoveGridmapEntryServiceUnavailable().WithPayload(response)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("gridmap", callTime)
                return data_set_management.NewRemoveGridmapEntryBadGateway().WithPayload(&returnValError)
        }
}
