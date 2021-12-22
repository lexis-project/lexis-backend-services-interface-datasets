package datasetManager

import (
//        l "gitlab.com/cyclops-utilities/logging"
        "context"
        "net/http"
	"strings"
//        "bytes"
        "io/ioutil"
//        "github.com/segmentio/encoding/json"
	"strconv"
	"time"
//	"fmt"
        "github.com/go-openapi/runtime/middleware"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
        "code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi/operations/data_set_management"
        )

func (u *DatasetManager) FilePatch (ctx context.Context, params data_set_management.FilePatchParams) 	middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("tus", callTime)
	etype := data_set_management.NewFilePatchServiceUnavailable()
	b := params.Body
	req, err := http.NewRequest (http.MethodPatch, u.ddiapi+"/upload/"+params.ID, strings.NewReader(b))
//params.HTTPRequest.Body is already a reader
	if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        if a, ok := params.HTTPRequest.Header["Authorization"]; ok {
        req.Header.Set("Authorization", a[0])
        }
	if ct, ok :=  params.HTTPRequest.Header["Content-Type"]; ok {
		req.Header.Set("Content-Type", ct[0])
	}
	req.Header.Set("Content-Length", strconv.FormatInt(params.ContentLength, 10))
	req.Header.Set("Upload-offset", strconv.FormatInt(params.UploadOffset, 10))
	req.Header.Set("Tus-Resumable", params.TusResumable)
	if params.UploadChecksum != nil {
		req.Header.Set("Upload-Checksum", *params.UploadChecksum)
	}
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.FilePatchNoContentCode) { //201
		UploadOffset, err2 := strconv.ParseInt(res.Header["Upload-Offset"][0], 10, 64)
		if (err2 == nil) {
			u.monit.APIHitDone("support", callTime)
			return data_set_management.NewFilePatchNoContent().
				WithTusResumable(res.Header["Tus-Resumable"][0]).
				WithUploadOffset(UploadOffset).WithUploadExpires(res.Header["Upload-Expires"][0])
		} else {
			u.monit.APIHitDone("support", callTime)
			return etype.WithPayload(fillErrorResponse(err2))
		}
	} else if (res.StatusCode == data_set_management.FilePatchUnsupportedMediaTypeCode) { //415
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchUnsupportedMediaType().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchNotFoundCode) { //404
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchNotFound().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchGoneCode) { //410
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchGone().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchForbiddenCode) { //403
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchForbidden().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchBadRequestCode) { // 400
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchBadRequest().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchStatus460Code) { //460
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchStatus460().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilePatchUnauthorizedCode) { //401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchUnauthorized().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.FilePatchServiceUnavailableCode) { //503
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilePatchServiceUnavailable()
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
        	return data_set_management.NewFilePatchServiceUnavailable().WithPayload(&returnValError)
	}
}

func (u *DatasetManager) FilesDelete (ctx context.Context, params data_set_management.FilesDeleteParams) 	middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("tus", callTime) 
	etype := data_set_management.NewFilesDeleteServiceUnavailable()
//	b := params.Body
	req, err := http.NewRequest (http.MethodDelete, u.ddiapi+"/upload/"+params.ID, nil)
	if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        if a, ok := params.HTTPRequest.Header["Authorization"]; ok {
        req.Header.Set("Authorization", a[0])
        }
	req.Header.Set("Tus-Resumable", params.TusResumable)
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.FilesDeleteNoContentCode) { // 204
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesDeleteNoContent().
				WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilesDeletePreconditionFailedCode) { // 412
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesDeletePreconditionFailed().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilesDeleteUnauthorizedCode) { // 401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesDeleteUnauthorized().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.FilesDeleteServiceUnavailableCode) { // 503
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesDeleteServiceUnavailable().WithPayload(rvalue)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
        	return data_set_management.NewFilesDeleteServiceUnavailable().WithPayload(&returnValError)	
	}
}

func (u *DatasetManager) FilesHead (ctx context.Context, params data_set_management.FilesHeadParams) 	middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("tus", callTime) 
	etype := data_set_management.NewFilesHeadServiceUnavailable()
//	b := params.Body
	req, err := http.NewRequest (http.MethodHead, u.ddiapi+"/upload/"+params.ID, nil)
	if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        if a, ok := params.HTTPRequest.Header["Authorization"]; ok {
        req.Header.Set("Authorization", a[0])
        }
	req.Header.Set("Tus-Resumable", params.TusResumable)
        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.FilesHeadOKCode) { // 200
		UploadOffset, err2 := strconv.ParseInt(res.Header["Upload-Offset"][0], 10, 64)
		if (err2 != nil) {
			u.monit.APIHitDone("support", callTime)
			return etype.WithPayload(fillErrorResponse(err2))
		}
		UploadLength, err3 := strconv.ParseInt(res.Header["Upload-Length"][0], 10, 64)
		if (err3 != nil) {
			u.monit.APIHitDone("support", callTime)
			return etype.WithPayload(fillErrorResponse(err3))
		}
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadOK().
				WithTusResumable(res.Header["Tus-Resumable"][0]).
				WithUploadOffset(UploadOffset).
				WithUploadLength(UploadLength).
				WithCacheControl(res.Header["Cache-Control"][0])
	} else if (res.StatusCode == data_set_management.FilesHeadUnauthorizedCode) { // 401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadUnauthorized().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.FilesHeadNotFoundCode) { // 404
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadNotFound().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilesHeadGoneCode) { // 410
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadGone().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilesHeadForbiddenCode) { // 403
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadForbidden().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.FilesHeadServiceUnavailableCode) { // 503
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadServiceUnavailable().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.FilesHeadPreconditionFailedCode) { // 412
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewFilesHeadPreconditionFailed().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
        	return data_set_management.NewFilesHeadServiceUnavailable().WithPayload(&returnValError)	
	}
}

func (u *DatasetManager) OptionsDatasetUpload (ctx context.Context, 
	params data_set_management.OptionsDatasetUploadParams) 	middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("tus", callTime) 
	etype := data_set_management.NewOptionsDatasetUploadServiceUnavailable()
//	b := params.Body
	req, err := http.NewRequest (http.MethodOptions, u.ddiapi+"/upload/", nil)
	if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        if a, ok := params.HTTPRequest.Header["Authorization"]; ok {
        req.Header.Set("Authorization", a[0])
        }

        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.OptionsDatasetUploadOKCode) { // 200
		TusMaxSize, err2 := strconv.ParseInt(res.Header["Tus-Max-Size"][0], 10, 64)
		if (err2 != nil) {
			u.monit.APIHitDone("support", callTime)
			return etype.WithPayload(fillErrorResponse(err2))
		}
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewOptionsDatasetUploadOK().
				WithTusResumable(res.Header["Tus-Resumable"][0]).
				WithTusMaxSize(TusMaxSize).
				WithTusVersion(res.Header["Tus-Version"][0]).
				WithTusExtension(res.Header["Tus-Extension"][0]).
				WithTusChecksumAlgorithm(res.Header["Tus-Checksum-Algorithm"][0])
	} else 	if (res.StatusCode == data_set_management.OptionsDatasetUploadNoContentCode) { // 204
		TusMaxSize, err2 := strconv.ParseInt(res.Header["Tus-Max-Size"][0], 10, 64)
		if (err2 != nil) {
			u.monit.APIHitDone("support", callTime)
			return etype.WithPayload(fillErrorResponse(err2))
		}
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewOptionsDatasetUploadNoContent().
				WithTusResumable(res.Header["Tus-Resumable"][0]).
				WithTusMaxSize(TusMaxSize).
				WithTusVersion(res.Header["Tus-Version"][0]).
				WithTusExtension(res.Header["Tus-Extension"][0]).
				WithTusChecksumAlgorithm(res.Header["Tus-Checksum-Algorithm"][0])
	} else if (res.StatusCode == data_set_management.OptionsDatasetUploadUnauthorizedCode) { // 401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewOptionsDatasetUploadUnauthorized().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.OptionsDatasetUploadServiceUnavailableCode) { // 503
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewOptionsDatasetUploadServiceUnavailable().WithPayload(rvalue)
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
        	return data_set_management.NewOptionsDatasetUploadServiceUnavailable().WithPayload(&returnValError)	
	}	
}

func (u *DatasetManager) PostDatasetUpload (ctx context.Context, 
	params data_set_management.PostDatasetUploadParams) 	middleware.Responder {
        callTime := time.Now()
	u.monit.APIHit("tus", callTime) 
	etype := data_set_management.NewPostDatasetUploadServiceUnavailable()
//	b := params.Body
	req, err := http.NewRequest (http.MethodPost, u.ddiapi+"/upload/", params.HTTPRequest.Body)
	if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }

	if a, ok := params.HTTPRequest.Header["Authorization"]; ok {
	req.Header.Set("Authorization", a[0])
	}
	if params.ContentLength != nil {
		req.Header.Set("Content-Length", strconv.FormatInt(*params.ContentLength, 10))
	}
	if params.UploadLength != nil {
		req.Header.Set("Upload-Length", strconv.FormatInt(*params.UploadLength, 10))
	}
	req.Header.Set("Tus-Resumable", params.TusResumable)
	if params.UploadMetadata != nil {
		req.Header.Set("Upload-Metadata", *params.UploadMetadata)
	}
	if params.UploadConcat != nil {
		req.Header.Set("Upload-Concat", *params.UploadConcat)
	}
	if params.UploadDeferLength!= nil {
		req.Header.Set("Upload-Defer-Length", strconv.FormatInt(*params.UploadDeferLength, 10))
	}
	if params.UploadChecksum!= nil {
		req.Header.Set("Upload-Checksum", *params.UploadChecksum)
	}
	if params.UploadOffset != nil {
		req.Header.Set("Upload-Offset", strconv.FormatInt(*params.UploadOffset, 10))
	}

        res, err := u.httpclient.Do(req)
        if err != nil {
		u.monit.APIHitDone("support", callTime)
                return etype.WithPayload(fillErrorResponse(err))
        }
        defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (res.StatusCode == data_set_management.PostDatasetUploadCreatedCode) { // 201
		u.monit.APIHitDone("support", callTime)
		// proxy location
		l := res.Header["Location"][0]
		if u.baseurl[len(u.baseurl)-1] =='/' { //avoid double /
			l = u.baseurl + params.HTTPRequest.URL.Path[1:] + strings.TrimPrefix(l, u.ddiapi+"/upload/")
		} else {
			l = u.baseurl + params.HTTPRequest.URL.Path + strings.TrimPrefix(l, u.ddiapi+"/upload/")
		}
		return data_set_management.NewPostDatasetUploadCreated().
				WithTusResumable(res.Header["Tus-Resumable"][0]).
				WithLocation(l)
	} else if (res.StatusCode == data_set_management.FilesHeadUnauthorizedCode) { // 401
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewPostDatasetUploadUnauthorized().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.PostDatasetUploadRequestEntityTooLargeCode) { // 413
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewPostDatasetUploadRequestEntityTooLarge().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else if (res.StatusCode == data_set_management.PostDatasetUploadServiceUnavailableCode) { // 503
		rvalue , err := fillErrorResponseFromBody (body)
                if err != nil {
			u.monit.APIHitDone("support", callTime)
                        return etype.WithPayload(fillErrorResponse(err))
                }
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewPostDatasetUploadServiceUnavailable().WithPayload(rvalue)
	} else if (res.StatusCode == data_set_management.PostDatasetUploadPreconditionFailedCode) { // 412
		u.monit.APIHitDone("support", callTime)
		return data_set_management.NewPostDatasetUploadPreconditionFailed().
			WithTusResumable(res.Header["Tus-Resumable"][0])
	} else {
		s:="Unexpected status code from wp3 backend: "+strconv.Itoa(res.StatusCode)+"; response: "+string(body)
                returnValError := models.ErrorResponse{
                        ErrorString: &s,
                }
		u.monit.APIHitDone("support", callTime)
        	return data_set_management.NewPostDatasetUploadServiceUnavailable().WithPayload(&returnValError)	
	}
}
