package datasetManager

import (
	l "gitlab.com/cyclops-utilities/logging"
	"context"
	"net/http"
        "bytes"
	"io/ioutil"
	"github.com/segmentio/encoding/json"
	"strconv"
	"time"
	"github.com/go-openapi/runtime/middleware"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/server/statusManager"
	)

type DatasetManager struct {
	ddiapi string
	gridmapapi string
	sshfsapi string
	steeringapi string
	compressapi string
	datasizeapi string
        encryptionapi string
	httpclient http.Client
	monit    *statusManager.StatusManager
	baseurl string
}

func New (api string, gridmap string, sshfs string, steering string, compress string, datasize string, encryption string, monit *statusManager.StatusManager, baseurl string) *DatasetManager {
	ep := []string{"certificate", "cloudnfs","compress","datasetManager","gridmap","project","replicate", "dataSize",
			"sshfs", "stagingDownload","steering","support","tus","user", "duplicate"}
	for _, s := range ep {
		monit.InitEndpoint(s)
	}
	return &DatasetManager{
		ddiapi: api,
		gridmapapi: gridmap,
		sshfsapi: sshfs,
		steeringapi: steering,
                compressapi: compress,
		datasizeapi: datasize,
                encryptionapi: encryption,
		httpclient: http.Client{},
		monit: monit,
		baseurl: baseurl,
	}
}

func fillErrorResponse (e error) *models.ErrorResponse {
	s := e.Error()
        returnValError := models.ErrorResponse{
                        ErrorString: &s,
		}
	return &returnValError
}

func fillErrorResponseFromBody (body []byte) (*models.ErrorResponse, error) {
	rvalue := models.ErrorResponse{}
	err := json.Unmarshal(body, &rvalue)
	return &rvalue, err
}

type AbstractParams interface {
	// Response
}

func (u *DatasetManager) performWp3Request (params AbstractParams, bearer string, method string, path string) (*http.Response, *models.ErrorResponse ) {
	b, err := json.Marshal(params)
        if err != nil {
                return nil, fillErrorResponse(err)
        }
        req, err := http.NewRequest (method, u.ddiapi+path, bytes.NewBuffer(b))
        if err != nil {
                return nil, fillErrorResponse(err)
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", bearer)
        res, err := u.httpclient.Do(req)
        if err != nil {
                return nil, fillErrorResponse(err)
        }
	return res, nil
}

func (u *DatasetManager) performQueryDatasets (params *models.MetadataQuery, bearer string, method string) (*http.Response, *models.ErrorResponse ) {
	return u.performWp3Request (params, bearer, method, "/dataset/search/metadata")
}

func (u *DatasetManager) CheckPermission (ctx context.Context, params data_set_management.CheckPermissionParams) middleware.Responder {
	callTime := time.Now()
	u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewCheckPermissionServiceUnavailable()
        bearer := params.HTTPRequest.Header["Authorization"][0]
        method := http.MethodGet
        res, err1 := u.performWp3Request (params.Access, bearer, method, "/dataset/checkpermission")
        if err1 != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(err1)
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime) 
                return etype.WithPayload(fillErrorResponse(err))
        }
        if res.StatusCode != data_set_management.CheckPermissionOKCode { // 200
	 if res.StatusCode == data_set_management.CheckPermissionForbiddenCode { // 403
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime) 
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckPermissionForbidden().WithPayload(rvalue)
        } else
	        if res.StatusCode == data_set_management.CheckPermissionUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime) 
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckPermissionUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CheckPermissionServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime) 
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckPermissionServiceUnavailable().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CheckPermissionBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime) 
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckPermissionBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.CheckPermissionBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime) 
                return data_set_management.NewCheckPermissionBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime) 
                return data_set_management.NewCheckPermissionBadGateway().WithPayload(&returnValError)
        }

	}
	status := "200"
	responsebody := data_set_management.CheckPermissionOKBody{&status}
	u.monit.APIHitDone("datasetManager", callTime) 
	return data_set_management.NewCheckPermissionOK().WithPayload(&responsebody)

}

func (u *DatasetManager) QueryDatasets (ctx context.Context, params data_set_management.QueryDatasetsParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewQueryDatasetsInternalServerError()
	res, err1 :=u.performQueryDatasets(params.MetadataQuery, params.HTTPRequest.Header["Authorization"][0], http.MethodPost) 
	if err1 != nil {
		u.monit.APIHitDone("datasetManager", callTime) 
		return etype.WithPayload(err1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}

	if res.StatusCode != data_set_management.QueryDatasetsOKCode { //200
             if res.StatusCode == data_set_management.QueryDatasetsUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewQueryDatasetsUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.QueryDatasetsBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewQueryDatasetsBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.QueryDatasetsServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewQueryDatasetsServiceUnavailable().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.QueryDatasetsInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			l.Warning.Printf("Error in response: "+string(body))
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewQueryDatasetsInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.QueryDatasetsBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewQueryDatasetsBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewQueryDatasetsBadGateway().WithPayload(&returnValError)
        }

	}
	rvalue := []*models.MetadataQueryResponse{}
	err = json.Unmarshal(body, &rvalue)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}

	return data_set_management.NewQueryDatasetsOK().WithPayload(rvalue)
}

func (u *DatasetManager) DeleteDatasetByMetadata (ctx context.Context, params data_set_management.DeleteDatasetByMetadataParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewDeleteDatasetByMetadataInternalServerError()
	res, err1 :=u.performQueryDatasets(params.MetadataQuery, params.HTTPRequest.Header["Authorization"][0], http.MethodDelete)
        if err1 != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(err1)
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	if res.StatusCode != data_set_management.DeleteDatasetByMetadataNoContentCode { // 204
            if res.StatusCode == data_set_management.DeleteDatasetByMetadataUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataUnauthorized().WithPayload(rvalue)
	} else if res.StatusCode == data_set_management.DeleteDatasetByMetadataForbiddenCode { // 403
		response := data_set_management.DeleteDatasetByMetadataForbiddenBody{}
		err := json.Unmarshal(body, &response)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataForbidden().WithPayload(&response)	
        } else if res.StatusCode == data_set_management.DeleteDatasetByMetadataBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteDatasetByMetadataServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataServiceUnavailable().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DeleteDatasetByMetadataInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DeleteDatasetByMetadataBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetByMetadataBadGateway().WithPayload(&returnValError)
	}
	}
	u.monit.APIHitDone("datasetManager", callTime)
	return data_set_management.NewDeleteDatasetByMetadataNoContent()
}

func (u *DatasetManager) CreateDataset (ctx context.Context, params data_set_management.CreateDatasetParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewCreateDatasetInternalServerError()
	bearer := params.HTTPRequest.Header["Authorization"][0]
	method := http.MethodPost
	res, err1 := u.performWp3Request (params.Dataset, bearer, method, "/dataset")
        if err1 != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(err1)
        }
	defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        if res.StatusCode != data_set_management.CreateDatasetCreatedCode &&
           res.StatusCode != data_set_management.CreateDatasetOKCode {
/*                s:=string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
                return data_set_management.NewDeleteDatasetByMetadataServiceUnavailable().WithPayload(&returnValError)
*/
		if res.StatusCode == data_set_management.CreateDatasetUnauthorizedCode { // 401
		        rvalue , err := fillErrorResponseFromBody (body)
		        if err != nil {
				u.monit.APIHitDone("datasetManager", callTime)
		                return etype.WithPayload(fillErrorResponse(err))
		        }
			u.monit.APIHitDone("datasetManager", callTime)
		        return data_set_management.NewCreateDatasetUnauthorized().WithPayload(rvalue)
	   	} else if res.StatusCode == data_set_management.CreateDatasetForbiddenCode { //403
		        rvalue , err := fillErrorResponseFromBody (body)
		        if err != nil {
				u.monit.APIHitDone("datasetManager", callTime)
		                return etype.WithPayload(fillErrorResponse(err))
		        }
			u.monit.APIHitDone("datasetManager", callTime)
		        return data_set_management.NewCreateDatasetForbidden().WithPayload(rvalue)
		} else if res.StatusCode == data_set_management.CreateDatasetServiceUnavailableCode { //503
		        rvalue , err := fillErrorResponseFromBody (body)
		        if err != nil {
				u.monit.APIHitDone("datasetManager", callTime)
		                return etype.WithPayload(fillErrorResponse(err))
		        }
		        return data_set_management.NewCreateDatasetServiceUnavailable().WithPayload(rvalue)
		} else if res.StatusCode == data_set_management.CreateDatasetBadRequestCode { // 400
		        response, err := fillErrorResponseFromBody(body)
		        if err != nil {
				u.monit.APIHitDone("datasetManager", callTime)
		                return etype.WithPayload(fillErrorResponse(err))
		        }
		        return data_set_management.NewCreateDatasetBadRequest().WithPayload(response)
		} else if res.StatusCode == data_set_management.CreateDatasetInternalServerErrorCode { // 500
		        response, err := fillErrorResponseFromBody(body)
		        if err != nil {
				u.monit.APIHitDone("datasetManager", callTime)
		                return etype.WithPayload(fillErrorResponse(err))
		        }
			u.monit.APIHitDone("datasetManager", callTime)
		        return data_set_management.NewCreateDatasetInternalServerError().WithPayload(response)
		} else if res.StatusCode == data_set_management.CreateDatasetBadGatewayCode { // 502, from nginx, html
		        s:="502 error (Bad Gateway) from wp3 backend"
		        returnValError := models.ErrorResponse{
		                ErrorString: &s,
		        }
			u.monit.APIHitDone("datasetManager", callTime)
		        return data_set_management.NewCreateDatasetBadGateway().WithPayload(&returnValError)

		} else {
		        s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
		        returnValError := models.ErrorResponse{
		                ErrorString: &s,
		        }
			u.monit.APIHitDone("datasetManager", callTime)
		        return data_set_management.NewCreateDatasetBadGateway().WithPayload(&returnValError)
		}

	}
	rvalue := models.ItemCreatedResponse{}
        err = json.Unmarshal(body, &rvalue)
	if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	u.monit.APIHitDone("datasetManager", callTime)
	if res.StatusCode == data_set_management.CreateDatasetCreatedCode {
        	return data_set_management.NewCreateDatasetCreated().WithPayload(&rvalue)
	} else {
		return data_set_management.NewCreateDatasetOK().WithPayload(&rvalue)
	}
}

func (u *DatasetManager) DownloadDataset (ctx context.Context, params data_set_management.DownloadDatasetParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewDownloadDatasetInternalServerError()
	b, err := json.Marshal(params.JSON)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}
	req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/dataset/download", bytes.NewBuffer(b))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
	res, err := u.httpclient.Do(req)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	if res.StatusCode == data_set_management.DownloadDatasetOKCode {
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDownloadDatasetOK().WithPayload(res.Body)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }

	if res.StatusCode == data_set_management.DownloadDatasetUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDownloadDatasetUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DownloadDatasetForbiddenCode { //403 
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDownloadDatasetForbidden().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DownloadDatasetServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDownloadDatasetServiceUnavailable().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DownloadDatasetBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDownloadDatasetBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.DownloadDatasetInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDownloadDatasetInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.DownloadDatasetBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDownloadDatasetBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDownloadDatasetBadGateway().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Listing (ctx context.Context, params data_set_management.ListingParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewListingServiceUnavailable()
	b, err := json.Marshal(params.JSON)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}
	req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/dataset/listing", bytes.NewBuffer(b))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
	res, err := u.httpclient.Do(req)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}

	if res.StatusCode != data_set_management.ListingOKCode { //200
             if res.StatusCode == data_set_management.ListingUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewListingUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.ListingForbiddenCode { //403 
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewListingForbidden().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.ListingServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewListingServiceUnavailable().WithPayload(rvalue)
	} else if res.StatusCode == data_set_management.ListingBadRequestCode { // 400
		response, err := fillErrorResponseFromBody(body)
		if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
			return etype.WithPayload(fillErrorResponse(err))
		}
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewListingBadRequest().WithPayload(response)
        } else if res.StatusCode == data_set_management.ListingBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewListingBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewListingBadGateway().WithPayload(&returnValError)
        }

	}
	rvalue := models.DatasetContent{}
	err = json.Unmarshal(body, &rvalue)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
		return etype.WithPayload(fillErrorResponse(err))
	}
	u.monit.APIHitDone("datasetManager", callTime)
	return data_set_management.NewListingOK().WithPayload(&rvalue)
}

func (u *DatasetManager) DeleteDataset (ctx context.Context, params data_set_management.DeleteDatasetParams) middleware.Responder {
        callTime := time.Now()
        u.monit.APIHit("datasetManager", callTime)
	etype := data_set_management.NewDeleteDatasetInternalServerError()
	b, err := json.Marshal(params.JSON)
        if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/dataset", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		u.monit.APIHitDone("datasetManager", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }

        if res.StatusCode == data_set_management.DeleteDatasetCreatedCode { // 201
		rvalue := data_set_management.DeleteDatasetCreatedBody{}
		err := json.Unmarshal(body, &rvalue)
		if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
			return etype.WithPayload(fillErrorResponse(err))
		}
		u.monit.APIHitDone("datasetManager", callTime)
		return data_set_management.NewDeleteDatasetCreated().WithPayload(&rvalue)
	}
	if res.StatusCode == data_set_management.DeleteDatasetNoContentCode { // 204
		u.monit.APIHitDone("datasetManager", callTime)
		return data_set_management.NewDeleteDatasetNoContent()
	} else if res.StatusCode == data_set_management.DeleteDatasetUnauthorizedCode { // 401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
		return data_set_management.NewDeleteDatasetUnauthorized().WithPayload(rvalue)
	} else if res.StatusCode == data_set_management.DeleteDatasetForbiddenCode { // 403
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetForbidden().WithPayload(rvalue)
	} else if res.StatusCode == data_set_management.DeleteDatasetServiceUnavailableCode { // 503
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
		return data_set_management.NewDeleteDatasetServiceUnavailable().WithPayload(rvalue)
	} else if res.StatusCode == data_set_management.DeleteDatasetBadRequestCode { // 400
		response, err := fillErrorResponseFromBody(body)
		if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
			return etype.WithPayload(fillErrorResponse(err))
		}
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetBadRequest().WithPayload(response)
	} else if res.StatusCode == data_set_management.DeleteDatasetInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("datasetManager", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetInternalServerError().WithPayload(response)
	} else if res.StatusCode == data_set_management.DeleteDatasetBadGatewayCode { // 502, from nginx, html
		s:="502 error (Bad Gateway) from wp3 backend"
		returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
                return data_set_management.NewDeleteDatasetBadGateway().WithPayload(&returnValError)

	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("datasetManager", callTime)
		return data_set_management.NewDeleteDatasetBadGateway().WithPayload(&returnValError)
	}

}
