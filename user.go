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

//AddUser AddUser
func (a *NotesAPI) AddUser(u *User) *Response {
	var rtn Response
	var aurl = a.restURL + "/rs/user/add"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(u)
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

//UpdateUser UpdateUser
func (a *NotesAPI) UpdateUser(u *User) *Response {
	var rtn Response
	var url = a.restURL + "/rs/user/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(u)
	if err == nil {
		reqacuu := a.buildRequest(put, url, a.headers, aJSON)
		acuusuc, stat := a.proxy.Do(reqacuu, &rtn)
		a.log.Debug("suc: ", acuusuc)
		a.log.Debug("stat: ", stat)
		if !acuusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//AddUserToNote AddUserToNote
func (a *NotesAPI) AddUserToNote(n *NoteUsers) *Response {
	var rtn Response
	var aurl = a.restURL + "/rs/note/user/add"
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

//GetNoteUserList GetNoteUserList
func (a *NotesAPI) GetNoteUserList(noteID int64, ownerEmail string) *[]string {
	var rtn []string
	idStr := strconv.FormatInt(noteID, 10)
	var url = a.restURL + "/rs/note/users/" + idStr + "/" + ownerEmail
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetUser GetUser
func (a *NotesAPI) GetUser(email string) *User {
	var rtn User
	var url = a.restURL + "/rs/user/get/" + email
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//Login Login
func (a *NotesAPI) Login(u *User) *LoginResponse {
	var rtn LoginResponse
	var aurl = a.restURL + "/rs/user/login"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(u)
	if err == nil {
		reqacu := a.buildRequest(post, aurl, a.headers, aJSON)
		sucaacu, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", sucaacu)
		a.log.Debug("stat: ", stat)
	}
	a.log.Debug("rtn: ", rtn)

	return &rtn
}
