package cocka2notesapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

func TestNotesAPI_AddCheckboxItem(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var ni CheckboxNoteItem
	ni.NoteID = 4
	ni.Checked = true
	ni.Text = "lets do some note stuff"

	res := api.AddCheckboxItem(&ni)

	fmt.Println("AddCheckboxItem: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_AddCheckboxItemFail(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":false}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var ni CheckboxNoteItem
	ni.NoteID = 4
	ni.Checked = true
	ni.Text = "lets do some note stuff"

	res := api.AddCheckboxItem(&ni)

	fmt.Println("AddCheckboxItem: ", res)
	//fmt.Println("FailAddCheckboxNoteList: ", sapi.FailAddCheckboxNoteList)

	if res.Success || len(sapi.FailAddCheckboxNoteList) != 1 {
		t.Fail()
	}
}

func TestNotesAPI_UpdateCheckboxItem(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var ni CheckboxNoteItem
	ni.ID = 5
	ni.NoteID = 4
	ni.Checked = true
	ni.Text = "lets do some note stuff and even more"

	res := api.UpdateCheckboxItem(&ni)

	fmt.Println("UpdateCheckboxItem: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_UpdateCheckboxItemFail(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":false}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var ni CheckboxNoteItem
	ni.ID = 5
	ni.NoteID = 4
	ni.Checked = true
	ni.Text = "lets do some note stuff and even more"

	res := api.UpdateCheckboxItem(&ni)

	fmt.Println("UpdateCheckboxItem: ", res)
	//fmt.Println("FailUpdateCheckboxNoteList: ", sapi.FailUpdateCheckboxNoteList)

	if res.Success || len(sapi.FailUpdateCheckboxNoteList) != 1{
		t.Fail()
	}
}

func TestNotesAPI_DeleteCheckboxItem(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var ni CheckboxNoteItem
	ni.ID = 5
	ni.NoteID = 4
	ni.Checked = true
	ni.Text = "lets do some note stuff and even more"

	res := api.DeleteCheckboxItem(5)

	fmt.Println("DeleteCheckboxItem: ", res)

	if !res.Success {
		t.Fail()
	}
}
