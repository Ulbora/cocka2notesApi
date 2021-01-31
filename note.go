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
		if rtn == nil {
			rtn = &ncbn
		}
	}

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return rtn
}

//GetNote GetNote
func (a *NotesAPI) GetNote(id int64) *Note {
	var rtn *Note
	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/note/get/" + idStr
	a.log.Debug("url: ", url)
	var txn Note
	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &txn)
	if suc && stat == 200 {
		a.setSavedTextNote(&txn)
		rtn = &txn
	} else {
		rtn = a.getSavedTextNote(id)
		if rtn == nil {
			rtn = &txn
		}
	}
	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return rtn
}

//GetUsersNotes GetUsersNotes
func (a *NotesAPI) GetUsersNotes(email string) *[]Note {
	var rtn []Note
	var url = a.restURL + "/rs/note/get/all/" + email
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)
	if suc && stat == 200 {
		a.noteList = rtn
	} else {
		rtn = a.noteList
	}

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return &rtn
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
	var found bool
	for _, cbn := range a.checkboxNoteList {
		if cbn.ID == cb.ID {
			a.log.Debug("found : ", cbn)
			cbn.LastUsed = cb.LastUsed
			cbn.NoteItems = cb.NoteItems
			cbn.Title = cb.Title
			found = true
			break
		}
	}
	if !found {
		a.log.Debug("not found : ")
		a.checkboxNoteList = append(a.checkboxNoteList, cb)
	}
}

func (a *NotesAPI) getSavedCheckboxNote(id int64) *CheckboxNote {
	var cbn *CheckboxNote
	for _, cb := range a.checkboxNoteList {
		if cb.ID == id {
			a.log.Debug("getting found cb note: ", cb)
			cbn = cb
			break
		}
	}
	return cbn
}

func (a *NotesAPI) setSavedTextNote(tn *Note) {
	var found bool
	for _, txn := range a.textNoteList {
		if txn.ID == tn.ID {
			a.log.Debug("found : ", txn)
			txn.LastUsed = tn.LastUsed
			txn.NoteItems = tn.NoteItems
			txn.Title = tn.Title
			found = true
			break
		}
	}
	if !found {
		a.log.Debug("not found : ")
		a.textNoteList = append(a.textNoteList, tn)
	}
}

func (a *NotesAPI) getSavedTextNote(id int64) *Note {
	var txn *Note
	for _, tn := range a.textNoteList {
		if tn.ID == id {
			a.log.Debug("getting found cb note: ", tn)
			txn = tn
			break
		}
	}
	return txn
}
