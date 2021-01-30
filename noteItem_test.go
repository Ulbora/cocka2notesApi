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

func TestNotesAPI_AddNoteItem(t *testing.T) {
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

	var ni NoteItem
	ni.NoteID = 4
	ni.Text = "lets do some note stuff"

	res := api.AddNoteItem(&ni)

	fmt.Println("AddNoteItem: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_AddNoteItemFail(t *testing.T) {
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

	var ni NoteItem
	ni.NoteID = 4
	ni.Text = "lets do some note stuff"

	res := api.AddNoteItem(&ni)

	fmt.Println("AddNoteItem: ", res)
	//fmt.Println("FailAddNoteItemList: ", sapi.FailAddNoteItemList)

	if res.Success {
		t.Fail()
	}
}

func TestNotesAPI_UpdateNoteItem(t *testing.T) {
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

	var ni NoteItem
	ni.ID = 4
	ni.NoteID = 4
	ni.Text = "lets do some note stuff and even more"

	res := api.UpdateNoteItem(&ni)

	fmt.Println("UpdateNoteItem: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_UpdateNoteItemFail(t *testing.T) {
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

	var ni NoteItem
	ni.ID = 4
	ni.NoteID = 4
	ni.Text = "lets do some note stuff and even more"

	res := api.UpdateNoteItem(&ni)

	fmt.Println("UpdateNoteItem: ", res)
	//fmt.Println("FailUpdateNoteItemList: ", sapi.FailUpdateNoteItemList)

	if res.Success {
		t.Fail()
	}
}

func TestNotesAPI_DeleteNoteItem(t *testing.T) {
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

	res := api.DeleteNoteItem(4)

	fmt.Println("DeleteNoteItem: ", res)

	if !res.Success {
		t.Fail()
	}
}
