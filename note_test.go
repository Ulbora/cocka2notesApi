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

func TestNotesAPI_AddNote(t *testing.T) {
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

	var n Note
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes"
	n.Type = "checkbox"

	res := api.AddNote(&n)

	fmt.Println("AddNote: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_AddNoteFail(t *testing.T) {
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

	var n Note
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes"
	n.Type = "checkbox"

	res := api.AddNote(&n)

	fmt.Println("AddNote: ", res)

	if res.Success {
		t.Fail()
	}
}

func TestNotesAPI_UpdateNote(t *testing.T) {
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

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.UpdateNote(&n)

	fmt.Println("UpdateNote: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_UpdateNoteFail(t *testing.T) {
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

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.UpdateNote(&n)

	fmt.Println("UpdateNote: ", res)

	if res.Success {
		t.Fail()
	}
}

func TestNotesAPI_GetCheckboxNote(t *testing.T) {
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
		"id": 1,
		"title": "some more note",
		"type": "checkbox",
		"ownerEmail": "tester@tester.com",
		"noteItems": [
			{
				"id": 1,
				"text": "some note",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 2,
				"text": "some note2",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 3,
				"text": "some note3",
				"checked": false,
				"noteId": 1
			}
		]
	}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.GetCheckboxNote(1)

	fmt.Println("GetNote: ", res)
	fmt.Println("GetNote item: ", res.NoteItems)

	if res == nil {
		t.Fail()
	}
}

func TestNotesAPI_GetCheckboxNoteFail(t *testing.T) {
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
		"id": 1,
		"title": "some more note",
		"type": "checkbox",
		"ownerEmail": "tester@tester.com",
		"noteItems": [
			{
				"id": 1,
				"text": "some note",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 2,
				"text": "some note2",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 3,
				"text": "some note3",
				"checked": false,
				"noteId": 1
			}
		]
	}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.GetCheckboxNote(1)

	fmt.Println("GetNote: ", res)
	fmt.Println("GetNote item: ", res.NoteItems)

	if res == nil {
		t.Fail()
	}
}

func TestNotesAPI_GetNote(t *testing.T) {
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
		"id": 1,
		"title": "some more note",
		"type": "checkbox",
		"ownerEmail": "tester@tester.com",
		"noteItems": [
			{
				"id": 1,
				"text": "some note",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 2,
				"text": "some note2",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 3,
				"text": "some note3",
				"checked": false,
				"noteId": 1
			}
		]
	}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.GetNote(1)

	fmt.Println("GetNote: ", res)
	fmt.Println("GetNote item: ", res.NoteItems)

	if res == nil {
		t.Fail()
	}
}

func TestNotesAPI_GetNoteFail(t *testing.T) {
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
		"id": 1,
		"title": "some more note",
		"type": "checkbox",
		"ownerEmail": "tester@tester.com",
		"noteItems": [
			{
				"id": 1,
				"text": "some note",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 2,
				"text": "some note2",
				"checked": false,
				"noteId": 1
			},
			{
				"id": 3,
				"text": "some note3",
				"checked": false,
				"noteId": 1
			}
		]
	}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.GetNote(1)

	fmt.Println("GetNote: ", res)
	fmt.Println("GetNote item: ", res.NoteItems)

	if res == nil {
		t.Fail()
	}
}

func TestNotesAPI_GetUsersNotes(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": 1,
			"title": "some more note",
			"type": "checkbox",
			"ownerEmail": "tester@tester.com",
			"lastUsed": "2021-03-01T00:00:00Z"
		},
		{
			"id": 3,
			"title": "some note",
			"type": "note",
			"ownerEmail": "tester@tester.com",
			"lastUsed": "2021-01-03T00:00:00Z"
		}
	]`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "tester@tester.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res, suc := api.GetUsersNotes("tester@tester.com")

	fmt.Println("GetUsersNotes: ", res)
	fmt.Println("GetUsersNotes suc: ", suc)

	if res == nil && suc {
		t.Fail()
	}
}

func TestNotesAPI_GetUsersNotesFail(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": 1,
			"title": "some more note",
			"type": "checkbox",
			"ownerEmail": "tester@tester.com",
			"lastUsed": "2021-03-01T00:00:00Z"
		},
		{
			"id": 3,
			"title": "some note",
			"type": "note",
			"ownerEmail": "tester@tester.com",
			"lastUsed": "2021-01-03T00:00:00Z"
		}
	]`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	// var crt sdbi.Category
	// crt.Description = "test"
	// crt.Name = "stuff"

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res, suc := api.GetUsersNotes("test@test.com")

	fmt.Println("GetUsersNotes: ", res)

	if res == nil || suc {
		t.Fail()
	}
}

func TestNotesAPI_DeleteNote(t *testing.T) {
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

	var n Note
	n.ID = 4
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res := api.DeleteNote(5, "tester@tester.com")

	fmt.Println("DeleteNote: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestNotesAPI_setSavedCheckboxNote(t *testing.T) {
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
	cb.Title = "cb note 2"
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

	api.setSavedCheckboxNote(&cb)
	cb.Title = "cb note 2"

	//api.setSavedCheckboxNote(&cb)
	//fmt.Println("cb note list: ", sapi.checkboxNoteList[0])
	if sapi.noteList[0].Title != "cb note 2" {
		t.Fail()
	}

}

func TestNotesAPI_getSavedCheckboxNote(t *testing.T) {
	var sapi NotesAPI

	var cbilst []CheckboxNoteItem
	//var cbilst []interface{}
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

	// var cb CheckboxNote
	// cb.ID = 5
	// cb.LastUsed = time.Now()
	// cb.OwnerEmail = "tester@tst.com"
	// cb.NoteItems = cbilst
	// cb.Title = "cb note 1"
	// cb.Type = "checkbox"
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

	fcbn := api.getSavedCheckboxNote(5)
	fmt.Println("fcbn: ", fcbn)
	if fcbn.Title != "cb note 1" || len(fcbn.NoteItems) != 2 {
		t.Fail()
	}
}

func TestNotesAPI_setSavedTextNote(t *testing.T) {
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

	var cb TextNote
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "cb note 1"
	cb.Type = "note"
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

	//api.setSavedTextNote(&cb)
	cb.Title = "cb note 2"

	api.setSavedTextNote(&cb)

	ilst := sapi.noteList[0].NoteItems.([]NoteItem)
	fmt.Println("text note list: ", ilst[0])
	if len(ilst) != 2 {
		t.Fail()
	}

}

func TestNotesAPI_getSavedTextNote(t *testing.T) {
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

	var cb TextNote
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "text note 1"
	cb.Type = "note"
	api := sapi.GetNew()

	var ntlst []Note
	var n1 Note
	n1.ID = 5
	n1.LastUsed = time.Now()
	n1.OwnerEmail = "tester@tst.com"
	n1.NoteItems = cbilst
	n1.Title = "text note 1"
	n1.Type = "note"
	ntlst = append(ntlst, n1)
	sapi.noteList = ntlst

	//api.setSavedTextNote(&cb)

	fcbn := api.getSavedTextNote(5)
	if fcbn.Title != "text note 1" {
		t.Fail()
	}
}

func TestNotesAPI_SetNoteList(t *testing.T) {
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

	var cb TextNote
	cb.ID = 5
	cb.LastUsed = time.Now()
	cb.OwnerEmail = "tester@tst.com"
	cb.NoteItems = cbilst
	cb.Title = "text note 1"
	cb.Type = "note"
	api := sapi.GetNew()

	var ntlst []Note
	var n1 Note
	n1.ID = 5
	n1.LastUsed = time.Now()
	n1.OwnerEmail = "tester@tst.com"
	n1.NoteItems = cbilst
	n1.Title = "text note 1"
	n1.Type = "note"
	ntlst = append(ntlst, n1)
	//sapi.noteList = ntlst

	api.SetNoteList(ntlst)

	if len(sapi.noteList) != 1 {
		t.Fail()
	}

}
