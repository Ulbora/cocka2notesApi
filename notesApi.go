package cocka2notesapi

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

import (
	"bytes"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

//NotesAPI NotesAPI
type NotesAPI struct {
	proxy                          px.Proxy
	log                            *lg.Logger
	restURL                        string
	apiKey                         string
	headers                        *Headers
	checkboxNoteList               []*CheckboxNote
	textNoteList                   []*Note
	noteList                       []Note
	FailAddCheckboxNoteItemList    []CheckboxNoteItem
	FailUpdateCheckboxNoteItemList []CheckboxNoteItem
	FailAddNoteItemList            []NoteItem
	FailUpdateNoteItemList         []NoteItem
}

//GetNew GetNew
func (a *NotesAPI) GetNew() API {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	a.log = &l

	var p px.GoProxy
	a.proxy = &p

	return a
}

//SetLogger SetLogger
func (a *NotesAPI) SetLogger(l *lg.Logger) {
	a.log = l
}

func (a *NotesAPI) buildRequest(method string, url string, headers *Headers, aJSON []byte) *http.Request {
	if a.apiKey != "" {
		headers.Set("apiKey", a.apiKey)
	}
	var req *http.Request
	var err error
	if method == post || method == put {
		headers.Set("Content-Type", "application/json")
		req, err = http.NewRequest(method, url, bytes.NewBuffer(aJSON))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	a.log.Debug("err in build req: ", err)
	if err == nil {
		for k, v := range headers.headers {
			a.log.Debug("header: ", k, v)
			req.Header.Set(k, v)
		}
	}
	return req
}

//SetLogLevel SetLogLevel
func (a *NotesAPI) SetLogLevel(level int) {
	a.log.LogLevel = level
}

//SetRestURL SetRestURL
func (a *NotesAPI) SetRestURL(url string) {
	a.restURL = url
}

//SetAPIKey SetAPIKey
func (a *NotesAPI) SetAPIKey(key string) {
	a.apiKey = key
}

//SetHeader SetHeader
func (a *NotesAPI) SetHeader(head *Headers) {
	a.headers = head
}

//OverrideProxy OverrideProxy
func (a *NotesAPI) OverrideProxy(proxy px.Proxy) {
	a.proxy = proxy
}
