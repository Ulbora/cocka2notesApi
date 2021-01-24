package cocka2notesapi

import (
	"bytes"
	"fmt"
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNotesAPI_GetMailServer(t *testing.T) {
	var sapi NotesAPI
	//sapi.SetAPIKey("123")

	sapi.SetRestURL("http://localhost:3000")

	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var head Headers
	sapi.SetHeader(&head)

	api := sapi.GetNew()
	sapi.SetLogLevel(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": 4,
		"host":"testerhost",
		"username": "tester@tester.com",
		"password": "ddddd",
		"port": "333",
		"senderEmail": "tester@tester.com"

	}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	res := api.GetMailServer()

	fmt.Println("GetMailServer: ", res)

	if res == nil || res.SenderEmail != "tester@tester.com" {
		t.Fail()
	}
}
