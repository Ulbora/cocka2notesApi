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

func TestNotesAPI_FlushFailedCache(t *testing.T) {
	var sapi NotesAPI

	var lst []CheckboxNoteItem
	var cb1 CheckboxNoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailAddCheckboxNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailAddCheckboxNoteItemList: ", sapi.FailAddCheckboxNoteItemList)

	if len(sapi.FailAddCheckboxNoteItemList) > 0 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCacheFail(t *testing.T) {
	var sapi NotesAPI

	var lst []CheckboxNoteItem
	var cb1 CheckboxNoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailAddCheckboxNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailAddCheckboxNoteItemList: ", sapi.FailAddCheckboxNoteItemList)

	if len(sapi.FailAddCheckboxNoteItemList) != 1 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache2(t *testing.T) {
	var sapi NotesAPI

	var lst []CheckboxNoteItem
	var cb1 CheckboxNoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailUpdateCheckboxNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailUpdateCheckboxNoteItemList: ", sapi.FailUpdateCheckboxNoteItemList)

	if len(sapi.FailUpdateCheckboxNoteItemList) > 0 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache2Fail(t *testing.T) {
	var sapi NotesAPI

	var lst []CheckboxNoteItem
	var cb1 CheckboxNoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailUpdateCheckboxNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailUpdateCheckboxNoteItemList: ", sapi.FailUpdateCheckboxNoteItemList)

	if len(sapi.FailUpdateCheckboxNoteItemList) != 1 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache3(t *testing.T) {
	var sapi NotesAPI

	var lst []NoteItem
	var cb1 NoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailAddNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailAddNoteItemList: ", sapi.FailAddNoteItemList)

	if len(sapi.FailAddNoteItemList) > 0 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache3Fail(t *testing.T) {
	var sapi NotesAPI

	var lst []NoteItem
	var cb1 NoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailAddNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailAddNoteItemList: ", sapi.FailAddNoteItemList)

	if len(sapi.FailAddNoteItemList) != 1 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache4(t *testing.T) {
	var sapi NotesAPI

	var lst []NoteItem
	var cb1 NoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailUpdateNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailUpdateNoteItemList: ", sapi.FailUpdateNoteItemList)

	if len(sapi.FailUpdateNoteItemList) > 0 {
		t.Fail()
	}

}

func TestNotesAPI_FlushFailedCache4Fail(t *testing.T) {
	var sapi NotesAPI

	var lst []NoteItem
	var cb1 NoteItem
	lst = append(lst, cb1)
	//var cb2 CheckboxNoteItem
	//lst = append(lst, cb2)
	sapi.FailUpdateNoteItemList = lst

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
	//api := sapi.GetNew()

	api.FlushFailedCache()

	fmt.Println("sapi.FailUpdateNoteItemList: ", sapi.FailUpdateNoteItemList)

	if len(sapi.FailUpdateNoteItemList) != 1 {
		t.Fail()
	}

}
