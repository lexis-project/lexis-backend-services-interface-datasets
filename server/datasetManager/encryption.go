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
	//"fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )

func (u *DatasetManager) Encrypt (ctx context.Context, params data_set_management.EncryptParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewEncryptInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/encrypt", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.EncryptCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewEncryptCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.EncryptUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewEncryptUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.EncryptBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewEncryptBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.EncryptForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewEncryptForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.EncryptNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewEncryptNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.EncryptRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewEncryptRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.EncryptTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewEncryptTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.EncryptInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewEncryptInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewEncryptInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Decrypt (ctx context.Context, params data_set_management.DecryptParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewDecryptInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/decrypt", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.DecryptCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewDecryptCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.DecryptUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewDecryptUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.DecryptBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecryptBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecryptForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptInternalServerError().WithPayload(&returnValError)
        }
}


func (u *DatasetManager) Compress (ctx context.Context, params data_set_management.CompressParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCompressInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/compress", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CompressCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewCompressCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.CompressUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCompressUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CompressBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCompressBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCompressForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) Decompress (ctx context.Context, params data_set_management.DecompressParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewDecompressInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/decompress", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.DecompressCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewDecompressCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.DecompressUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewDecompressUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.DecompressBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecompressBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecompressForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecompressForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecompressNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecompressNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecompressRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecompressRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecompressTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecompressTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecompressInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecompressInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecompressInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) DecryptDecompress (ctx context.Context, params data_set_management.DecryptDecompressParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewDecryptDecompressInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/decrypt_decompress", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.DecryptDecompressCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewDecryptDecompressCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.DecryptDecompressUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewDecryptDecompressUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.DecryptDecompressBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecryptDecompressBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptDecompressForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewDecryptDecompressForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptDecompressNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptDecompressNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptDecompressRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptDecompressRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptDecompressTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptDecompressTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.DecryptDecompressInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptDecompressInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewDecryptDecompressInternalServerError().WithPayload(&returnValError)
        }
}


func (u *DatasetManager) CompressEncrypt (ctx context.Context, params data_set_management.CompressEncryptParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCompressEncryptInternalServerError()
	b, err := json.Marshal(params.Parameters)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req, err := http.NewRequest (http.MethodPost, u.encryptionapi+"/compress_encrypt", bytes.NewBuffer(b))
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.CompressEncryptCreatedCode) {
		ret := models.RequestID {}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
		return data_set_management.NewCompressEncryptCreated().WithPayload(&ret)
	}
	//Unexpected response
	if res.StatusCode == data_set_management.CompressEncryptUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCompressEncryptUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CompressEncryptBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCompressEncryptBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressEncryptForbiddenCode { // 403
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCompressEncryptForbidden().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressEncryptNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressEncryptNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressEncryptRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressEncryptRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressEncryptTooManyRequestsCode { // 429
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressEncryptTooManyRequests().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CompressEncryptInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressEncryptInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCompressEncryptInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckCompressionEncryptionStatus (ctx context.Context, params data_set_management.CheckCompressionEncryptionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckCompressionEncryptionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/compress_encrypt/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckCompressionEncryptionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionEncryptionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckCompressionEncryptionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckCompressionEncryptionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckCompressionEncryptionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckCompressionEncryptionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionEncryptionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionEncryptionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionEncryptionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionEncryptionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionEncryptionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionEncryptionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionEncryptionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckCompressionStatus (ctx context.Context, params data_set_management.CheckCompressionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckCompressionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/compress/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckCompressionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckCompressionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckCompressionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckCompressionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckCompressionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckCompressionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckCompressionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckDecompressionStatus (ctx context.Context, params data_set_management.CheckDecompressionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckDecompressionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/decompress/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckDecompressionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecompressionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckDecompressionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckDecompressionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckDecompressionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckDecompressionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecompressionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecompressionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecompressionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecompressionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecompressionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecompressionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecompressionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckDecryptionDecompressionStatus (ctx context.Context, params data_set_management.CheckDecryptionDecompressionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckDecryptionDecompressionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/decrypt_decompress/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionDecompressionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckDecryptionDecompressionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckDecryptionDecompressionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionDecompressionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionDecompressionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionDecompressionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionDecompressionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionDecompressionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckDecryptionStatus (ctx context.Context, params data_set_management.CheckDecryptionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckDecryptionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/decrypt/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckDecryptionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckDecryptionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckDecryptionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckDecryptionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckDecryptionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckDecryptionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckDecryptionStatusInternalServerError().WithPayload(&returnValError)
        }
}

func (u *DatasetManager) CheckEncryptionStatus (ctx context.Context, params data_set_management.CheckEncryptionStatusParams) middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("encryption", callTime)
	etype := data_set_management.NewCheckEncryptionStatusInternalServerError()

        req, err := http.NewRequest (http.MethodGet, u.encryptionapi+"/encrypt/"+params.RequestID.String(), nil)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", params.HTTPRequest.Header["Authorization"][0])
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("encryption", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if (res.StatusCode == data_set_management.CheckEncryptionStatusOKCode) {
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckEncryptionStatusOK()
        }
        //Unexpected response
	if res.StatusCode == data_set_management.CheckEncryptionStatusUnauthorizedCode { // 401
                rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }		
                return data_set_management.NewCheckEncryptionStatusUnauthorized().WithPayload(rvalue)
	  } else if res.StatusCode == data_set_management.CheckEncryptionStatusBadRequestCode { // 400
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
                return data_set_management.NewCheckEncryptionStatusBadRequest().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckEncryptionStatusNotFoundCode { // 404
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckEncryptionStatusNotFound().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckEncryptionStatusRequestURITooLongCode { // 414
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckEncryptionStatusRequestURITooLong().WithPayload(response)
	  } else if res.StatusCode == data_set_management.CheckEncryptionStatusInternalServerErrorCode { // 500
                response, err := fillErrorResponseFromBody(body)
                if err != nil {
			u.monit.APIHitDone("encryption", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckEncryptionStatusInternalServerError().WithPayload(response)
        } else {
                s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("encryption", callTime)
                return data_set_management.NewCheckEncryptionStatusInternalServerError().WithPayload(&returnValError)
        }
}

