package validator

/*func PaymentRequest(context *gin.Context) *UploadForm {
	var body UploadForm

	debtorIban := context.Request.FormValue("debtor_iban")
	debtorName := context.Request.FormValue("debtor_name")
	creditorIban := context.Request.FormValue("creditor_iban")
	creditorName := context.Request.FormValue("creditor_name")
	if debtorIban == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 Bad Request",
			"message": "Missing or invalid uniqueEventId",
			"error":   "uniqueEventId invalid",
			"success": false,
		})
		return nil
	}
	body.File = file
	body.FileHeader = fileHeader
	body.UniqueEventId = uniqueEventId
	body.RecordingType = recordingType
	return &body
}*/
