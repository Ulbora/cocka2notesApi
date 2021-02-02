package cocka2notesapi

import (
	"encoding/json"
	"strconv"
)

/*
   Six910 is a shopping cart and E-commerce system.

   Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
   All rights reserved.

   Copyright (C) 2020 Ken Williamson
   All rights reserved.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//AddNote AddNote
func (a *NotesAPI) AddNote(n *Note) *ResponseID {
	var rtn ResponseID
	var aurl = a.restURL + "/rs/note/add"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(n)
	if err == nil {
		reqacu := a.buildRequest(post, aurl, a.headers, aJSON)
		sucaacu, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", sucaacu)
		a.log.Debug("stat: ", stat)
		if !sucaacu {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateNote UpdateNote
func (a *NotesAPI) UpdateNote(n *Note) *Response {
	var rtn Response
	var url = a.restURL + "/rs/note/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(n)
	if err == nil {
		reqspu := a.buildRequest(put, url, a.headers, aJSON)
		spusuc, stat := a.proxy.Do(reqspu, &rtn)
		a.log.Debug("suc: ", spusuc)
		a.log.Debug("stat: ", stat)
		if !spusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetCheckboxNote GetCheckboxNote
func (a *NotesAPI) GetCheckboxNote(id int64) *CheckboxNote {
	var rtn *CheckboxNote

	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/note/get/" + idStr
	a.log.Debug("url: ", url)

	var ncbn CheckboxNote
	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &ncbn)
	if suc && stat == 200 {
		a.setSavedCheckboxNote(&ncbn)
		rtn = &ncbn
	} else {
		rtn = a.getSavedCheckboxNote(id)
	}

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return rtn
}

//GetNote GetNote
func (a *NotesAPI) GetNote(id int64) *TextNote {
	var rtn *TextNote
	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/note/get/" + idStr
	a.log.Debug("url: ", url)
	var txn TextNote
	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &txn)
	if suc && stat == 200 {
		a.setSavedTextNote(&txn)
		rtn = &txn
	} else {
		rtn = a.getSavedTextNote(id)
	}
	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return rtn
}

//GetUsersNotes GetUsersNotes
func (a *NotesAPI) GetUsersNotes(email string) (*[]Note, bool) {
	var rtn []Note
	var suc bool
	var url = a.restURL + "/rs/note/get/all/" + email
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)
	if suc && stat == 200 {
		a.noteList = rtn
		suc = true
	} else {
		rtn = a.noteList
	}

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return &rtn, suc
}

//DeleteNote DeleteNote
func (a *NotesAPI) DeleteNote(id int64, ownerEmail string) *Response {
	var rtn Response
	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/note/delete/" + idStr + "/" + ownerEmail
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, a.headers, nil)
	dspsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dspsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

func (a *NotesAPI) setSavedCheckboxNote(cb *CheckboxNote) {
	for i := range a.noteList {
		if a.noteList[i].ID == cb.ID {
			a.log.Debug("found : ", a.noteList[i])
			a.noteList[i].Title = cb.Title
			a.log.Debug("updated : ", a.noteList[i])
			break
		}
	}
}

func (a *NotesAPI) getSavedCheckboxNote(id int64) *CheckboxNote {
	var cbn CheckboxNote
	for i := range a.noteList {
		a.log.Debug("getting cb note: ", a.noteList[i])
		if a.noteList[i].ID == id {
			a.log.Debug("getting found cb note: ", a.noteList[i])
			cbn.ID = a.noteList[i].ID
			cbn.LastUsed = a.noteList[i].LastUsed
			cbn.OwnerEmail = a.noteList[i].OwnerEmail
			cbn.Title = a.noteList[i].Title
			cbn.Type = a.noteList[i].Type
			cbn.NoteItems = a.noteList[i].NoteCheckboxItems
			// var ilst []CheckboxNoteItem
			// if rec, ok := a.noteList[i].NoteItems.([]map[string]interface{}); ok {
			// 	for _, r := range rec {
			// 		var ci CheckboxNoteItem
			// 		for key, val := range r {
			// 			if key == "id" {
			// 				ci.ID = val.(int64)
			// 			} else if key == "noteId" {
			// 				ci.NoteID = val.(int64)
			// 			} else if key == "checked" {
			// 				ci.Checked = val.(bool)
			// 			} else if key == "text" {
			// 				ci.Text = val.(string)
			// 			}
			// 			a.log.Debug("key: ", key)
			// 			a.log.Debug("val: ", val)
			// 		}
			// 		ilst = append(ilst, ci)
			// 	}
			// 	cbn.NoteItems = ilst
			// } else {
			// 	ilst = a.noteList[i].NoteItems.([]CheckboxNoteItem)
			// 	cbn.NoteItems = ilst
			// 	//fmt.Printf("record not a map[string]interface{}: %v\n", record)
			// }
			// //ilst := a.noteList[i].NoteItems.([]CheckboxNoteItem)

			break
		}
	}
	return &cbn
}

func (a *NotesAPI) setSavedTextNote(tn *TextNote) {
	for i := range a.noteList {
		if a.noteList[i].ID == tn.ID {
			a.log.Debug("found : ", a.noteList[i])
			a.noteList[i].Title = tn.Title
			a.noteList[i].NoteTextItems = tn.NoteItems
			a.log.Debug("updated : ", a.noteList[i])
			break
		}
	}
}

func (a *NotesAPI) getSavedTextNote(id int64) *TextNote {
	var txn TextNote
	for i := range a.noteList {
		if a.noteList[i].ID == id {
			a.log.Debug("getting found cb note: ", a.noteList[i])
			txn.ID = a.noteList[i].ID
			txn.LastUsed = a.noteList[i].LastUsed
			txn.OwnerEmail = a.noteList[i].OwnerEmail
			txn.Title = a.noteList[i].Title
			txn.Type = a.noteList[i].Type
			txn.NoteItems = a.noteList[i].NoteTextItems
			break
		}
	}
	return &txn
}

//SetNoteList SetNoteList
func (a *NotesAPI) SetNoteList(noteList []Note) {
	a.noteList = noteList
}

//GetNoteList GetNoteList
func (a *NotesAPI) GetNoteList() []Note {
	return a.noteList
}
