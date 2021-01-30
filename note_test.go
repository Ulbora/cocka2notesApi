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
	n.OwnerEmail = "test@test.com"
	n.Title = "a note to end all notes and then some"
	n.Type = "checkbox"

	res, suc := api.GetUsersNotes("test@test.com")

	fmt.Println("GetUsersNotes: ", res)

	if res == nil || !suc {
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
