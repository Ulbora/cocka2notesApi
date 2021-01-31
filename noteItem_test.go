package cocka2notesapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"
	"time"

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

	if res.Success || len(sapi.FailAddNoteItemList) != 1 {
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

	if res.Success || len(sapi.FailUpdateNoteItemList) != 1 {
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

func TestNotesAPI_setSavedTextItem(t *testing.T) {
	var sapi NotesAPI

	var cbilst []NoteItem
	var cbi1 NoteItem
	cbi1.ID = 1
	cbi1.NoteID = 5
	cbi1.Text = "bread"
	cbilst = append(cbilst, cbi1)

	var cbi2 NoteItem
	cbi2.ID = 2
	cbi2.NoteID = 5
	cbi2.Text = "milk"
	cbilst = append(cbilst, cbi2)

	var cb Note
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "text note 1"
	cb.Type = "note"
	api := sapi.GetNew()

	api.setSavedTextNote(&cb)

	var cbi3 NoteItem
	cbi3.NoteID = 5
	cbi3.Text = "cream"

	api.setSavedTextItem(&cbi3)

	if len(sapi.textNoteList[0].NoteItems) != 3 {
		t.Fail()
	}

}

func TestNotesAPI_setSavedTextItem2(t *testing.T) {
	var sapi NotesAPI

	var cbilst []NoteItem
	var cbi1 NoteItem
	cbi1.ID = 1
	cbi1.NoteID = 5
	cbi1.Text = "bread"
	cbilst = append(cbilst, cbi1)

	var cbi2 NoteItem
	cbi2.ID = 2
	cbi2.NoteID = 5
	cbi2.Text = "milk"
	cbilst = append(cbilst, cbi2)

	var cb Note
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "cb note 1"
	cb.Type = "note"
	api := sapi.GetNew()

	api.setSavedTextNote(&cb)

	var cbi3 NoteItem
	cbi3.ID = 2
	cbi3.NoteID = 5
	cbi3.Text = "cream"

	api.setSavedTextItem(&cbi3)

	if len(sapi.textNoteList[0].NoteItems) != 2 {
		t.Fail()
	}

}

func TestNotesAPI_GetFailAddNoteItemList(t *testing.T) {
	var sapi NotesAPI
	api := sapi.GetNew()
	lst := api.GetFailAddNoteItemList()
	if len(lst) != 0 {
		t.Fail()
	}
}

func TestNotesAPI_GetFailUpdateNoteItemList(t *testing.T) {
	var sapi NotesAPI
	api := sapi.GetNew()
	lst := api.GetFailUpdateNoteItemList()
	if len(lst) != 0 {
		t.Fail()
	}
}
