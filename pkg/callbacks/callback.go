package callbacks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func PrintErrorMessage(err error) {
	var respError struct {
		Code    string `json:"code"`
		Details string `json:"details"`
		Message string `json:"message"`
		Subcode string `json:"subcode"`
	}
	if httpErr, ok := err.(*azcore.ResponseError); ok {
		body, _ := io.ReadAll(httpErr.RawResponse.Body)
		json.Unmarshal(body, &respError)
		fmt.Printf("Error: %s\n", respError.Message)
		os.Exit(1)
	} else {
		log.Fatal(err)
	}
}
