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

//AddNoteItem AddNoteItem
func (a *NotesAPI) AddNoteItem(ni *NoteItem) *ResponseID {
	var rtn ResponseID
	var aurl = a.restURL + "/rs/item/add"
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
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateNoteItem UpdateNoteItem
func (a *NotesAPI) UpdateNoteItem(ni *NoteItem) *Response {
	var rtn Response
	var aurl = a.restURL + "/rs/item/update"
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
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//DeleteNoteItem DeleteNoteItem
func (a *NotesAPI) DeleteNoteItem(id int64) *Response {
	var rtn Response
	idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/item/delete/" + idStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, a.headers, nil)
	dspsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dspsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
