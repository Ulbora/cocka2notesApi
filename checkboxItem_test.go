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

	if res.Success || len(sapi.FailAddCheckboxNoteItemList) != 1 {
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

	if res.Success || len(sapi.FailUpdateCheckboxNoteItemList) != 1 {
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

func TestNotesAPI_setSavedCheckboxItem(t *testing.T) {
	var sapi NotesAPI

	var cbilst []CheckboxNoteItem
	var cbi1 CheckboxNoteItem
	cbi1.Checked = true
	cbi1.ID = 1
	cbi1.NoteID = 5
	cbi1.Text = "bread"
	cbilst = append(cbilst, cbi1)

	var cbi2 CheckboxNoteItem
	cbi2.Checked = false
	cbi2.ID = 2
	cbi2.NoteID = 5
	cbi2.Text = "milk"
	cbilst = append(cbilst, cbi2)

	var cb CheckboxNote
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "cb note 1"
	cb.Type = "checkbox"
	api := sapi.GetNew()

	var ntlst []Note
	var n1 Note
	n1.ID = 5
	n1.LastUsed = time.Now()
	n1.OwnerEmail = "tester@tst.com"
	n1.NoteItems = cbilst
	n1.Title = "cb note 1"
	n1.Type = "checkbox"
	ntlst = append(ntlst, n1)
	sapi.noteList = ntlst

	//api.setSavedCheckboxNote(&cb)

	var cbi3 CheckboxNoteItem
	cbi3.Checked = false
	cbi3.NoteID = 5
	cbi3.Text = "cream"

	api.setSavedCheckboxItem(&cbi3)

	ilst := sapi.noteList[0].NoteItems.([]CheckboxNoteItem)

	if len(ilst) != 3 {
		t.Fail()
	}

}

func TestNotesAPI_setSavedCheckboxItem2(t *testing.T) {
	var sapi NotesAPI

	var cbilst []CheckboxNoteItem
	var cbi1 CheckboxNoteItem
	cbi1.Checked = true
	cbi1.ID = 1
	cbi1.NoteID = 5
	cbi1.Text = "bread"
	cbilst = append(cbilst, cbi1)

	var cbi2 CheckboxNoteItem
	cbi2.Checked = false
	cbi2.ID = 2
	cbi2.NoteID = 5
	cbi2.Text = "milk"
	cbilst = append(cbilst, cbi2)

	var cb CheckboxNote
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "cb note 1"
	cb.Type = "checkbox"
	api := sapi.GetNew()


	var ntlst []Note
	var n1 Note
	n1.ID = 5
	n1.LastUsed = time.Now()
	n1.OwnerEmail = "tester@tst.com"
	n1.NoteItems = cbilst
	n1.Title = "cb note 1"
	n1.Type = "checkbox"
	ntlst = append(ntlst, n1)
	sapi.noteList = ntlst

	//api.setSavedCheckboxNote(&cb)

	var cbi3 CheckboxNoteItem
	cbi3.ID = 2
	cbi3.Checked = true
	cbi3.NoteID = 5
	cbi3.Text = "cream"

	api.setSavedCheckboxItem(&cbi3)

	ilst := sapi.noteList[0].NoteItems.([]CheckboxNoteItem)

	if len(ilst) != 2 {
		t.Fail()
	}

}

func TestNotesAPI_GetFailAddCheckboxNoteItemList(t *testing.T) {
	var sapi NotesAPI
	api := sapi.GetNew()
	lst := api.GetFailAddCheckboxNoteItemList()
	if len(lst) != 0 {
		t.Fail()
	}
}

func TestNotesAPI_GetFailUpdateCheckboxNoteItemList(t *testing.T) {
	var sapi NotesAPI
	api := sapi.GetNew()
	lst := api.GetFailUpdateCheckboxNoteItemList()
	if len(lst) != 0 {
		t.Fail()
	}
}
