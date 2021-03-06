package cocka2notesapi

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
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

//AddCheckboxItem AddCheckboxItem
func (a *NotesAPI) AddCheckboxItem(ni *CheckboxNoteItem) *ResponseID {
	var rtn ResponseID
	var aurl = a.restURL + "/rs/checkbox/add"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(ni)
	if err == nil {
		reqacu := a.buildRequest(post, aurl, a.headers, aJSON)
		suc, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		if !suc {
			rtn.Code = int64(stat)
		}
		if (!suc || stat != 200) && !a.retryCall {
			a.FailAddCheckboxNoteItemList = append(a.FailAddCheckboxNoteItemList, *ni)
			a.setSavedCheckboxItem(ni)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateCheckboxItem UpdateCheckboxItem
func (a *NotesAPI) UpdateCheckboxItem(ni *CheckboxNoteItem) *Response {
	var rtn Response
	var aurl = a.restURL + "/rs/checkbox/update"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(ni)
	if err == nil {
		reqacu := a.buildRequest(put, aurl, a.headers, aJSON)
		suc, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		if !suc {
			rtn.Code = int64(stat)
		}
		if (!suc || stat != 200) && !a.retryCall {
			a.FailUpdateCheckboxNoteItemList = append(a.FailUpdateCheckboxNoteItemList, *ni)
			a.setSavedCheckboxItem(ni)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//DeleteCheckboxItem DeleteCheckboxItem
func (a *NotesAPI) DeleteCheckboxItem(id int64) *Response {
	var rtn Response
	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/checkbox/delete/" + idStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, a.headers, nil)
	dspsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dspsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

func (a *NotesAPI) setSavedCheckboxItem(cbi *CheckboxNoteItem) {
	//update existing item
	a.log.Debug("saving in updating: ", *cbi)
	if cbi.ID != 0 {
		for i := range a.noteList {
			a.log.Debug("saving in updating in for loop: ", a.noteList[i])
			if a.noteList[i].ID == cbi.NoteID {
				a.log.Debug("found cb and updating: ", a.noteList[i])
				for ii := range a.noteList[i].NoteCheckboxItems {
					a.log.Debug("found cb item: ", a.noteList[i].NoteCheckboxItems[ii])
					if a.noteList[i].NoteCheckboxItems[ii].ID == cbi.ID {
						a.log.Debug("found cb item and updating: ", a.noteList[i].NoteCheckboxItems[ii])
						a.log.Debug("updating to: ", *cbi)
						a.noteList[i].NoteCheckboxItems[ii].Checked = cbi.Checked
						a.noteList[i].NoteCheckboxItems[ii].Text = cbi.Text
						break
					}
				}
				break
			}
		}
	} else if cbi.NoteID != 0 {
		//add the item as new
		rand.Seed(time.Now().UnixNano())
		var nid = rand.Int63n(30000)
		cbi.ID = nid
		for i := range a.noteList {
			if a.noteList[i].ID == cbi.NoteID {
				a.log.Debug("adding new cb item: ", *cbi)
				//ilst := a.noteList[i].NoteItems.([]CheckboxNoteItem)
				//a.log.Debug("existing cb item: ", ilst)
				a.noteList[i].NoteCheckboxItems = append(a.noteList[i].NoteCheckboxItems, *cbi)
				//a.noteList[i].NoteItems = ilst
				break
			}
		}
	}
}

//GetFailAddCheckboxNoteItemList GetFailAddCheckboxNoteItemList
func (a *NotesAPI) GetFailAddCheckboxNoteItemList() []CheckboxNoteItem {
	return a.FailAddCheckboxNoteItemList
}

//GetFailUpdateCheckboxNoteItemList GetFailUpdateCheckboxNoteItemList
func (a *NotesAPI) GetFailUpdateCheckboxNoteItemList() []CheckboxNoteItem {
	return a.FailUpdateCheckboxNoteItemList
}
