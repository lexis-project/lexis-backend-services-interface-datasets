package datasetManager

import (
//	l "gitlab.com/cyclops-utilities/logging"
	"context"
	"net/http"
        "bytes"
	"io/ioutil"
	"strconv"
	"time"
	"github.com/go-openapi/runtime/middleware"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
	)

func (u *DatasetManager) Certificate (ctx context.Context, params data_set_management.CertificateParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("certificate", callTime)
	etype := data_set_management.NewCertificateInternalServerError()
	var b []byte
	req, err := http.NewRequest (http.MethodGet, u.ddiapi+"/cert", bytes.NewBuffer(b))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
	res, err := u.httpclient.Do(req)
	if err != nil {
		u.monit.APIHitDone("certificate", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
	if res.StatusCode == data_set_management.CertificateOKCode {
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateOK().WithPayload(res.Body)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			u.monit.APIHitDone("certificate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
        }

	if res.StatusCode == data_set_management.CertificateUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("certificate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateUnauthorized().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.CertificateServiceUnavailableCode { //503
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("certificate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateServiceUnavailable().WithPayload(rvalue)
        } else if res.StatusCode == data_set_management.DownloadDatasetInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("certificate", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateInternalServerError().WithPayload(response)
        } else if res.StatusCode == data_set_management.CertificateBadGatewayCode { // 502, from nginx, html
                s:="502 error (Bad Gateway) from wp3 backend"
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateBadGateway().WithPayload(&returnValError)

        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("certificate", callTime)
                return data_set_management.NewCertificateBadGateway().WithPayload(&returnValError)
        }
}

