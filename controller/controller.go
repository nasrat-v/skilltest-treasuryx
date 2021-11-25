package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// default constructor
func New() Controller {
	return Controller{}
}

func (x *Controller) Payment(context *gin.Context) {
	_, err := x.decodePaymentRequest(context.Request.Body)
	if err != nil {
		context.Abort()
		//context.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
	}
	/*if body := x.validateUploadRecordingForm(context); body == nil {
		return
	} else {
		defer body.File.Close()
		/*triggerInstanceId := context.Param("trigger_instance_id")

		event, err := x.findEventByUniqueEvent(context, body.UniqueEventId)
		if err != nil {
			return
		}
		eventId := strconv.Itoa(event.Id)
		metadatum, err := x.findMetadatumByEvent(context, eventId)
		if err != nil {
			return
		}
		metadatumId := strconv.Itoa(metadatum.Id)
		recording, err := x.createRecording(context, eventId, metadatumId, body.RecordingType)
		if err != nil {
			return
		}
		recordingId := strconv.Itoa(recording.Id)
		zip, err := x.unzipArchive(context, body)
		if err != nil {
			return
		}
		destPath := x.fullFilepath(triggerInstanceId, eventId, recordingId)
		fullpath, err := x.uploadUnzippedArchive(context, destPath, zip)
		if err != nil {
			return
		}
		x.saveCloudInfos(context, recordingId, fullpath)
	}*/
}

func (x *Controller) decodePaymentRequest(body io.ReadCloser) (Payment, error) {
	var payment Payment

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&payment)
	fmt.Println(payment)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return payment, err
	}
	return payment, nil
}
