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
	"sync"

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
	noteList                       []Note
	FailAddCheckboxNoteItemList    []CheckboxNoteItem
	FailUpdateCheckboxNoteItemList []CheckboxNoteItem
	FailAddNoteItemList            []NoteItem
	FailUpdateNoteItemList         []NoteItem
	retryCall                      bool
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

//FlushFailedCache FlushFailedCache
func (a *NotesAPI) FlushFailedCache() {
	if len(a.FailAddCheckboxNoteItemList) > 0 {
		a.log.Debug("FailAddCheckboxNoteItemList: ", a.FailAddCheckboxNoteItemList)
		var suc = true
		a.retryCall = true
		var wg sync.WaitGroup
		for i := range a.FailAddCheckboxNoteItemList {
			wg.Add(1)
			go func(cb CheckboxNoteItem) {
				defer wg.Done()
				res := a.AddCheckboxItem(&cb)
				a.log.Debug("AddCheckboxItem suc: ", res.Success)
				if !res.Success {
					suc = false
				}
			}(a.FailAddCheckboxNoteItemList[i])
		}
		a.log.Debug("waiting for go routines in add checkitem")
		wg.Wait()
		a.retryCall = false
		if suc {
			a.FailAddCheckboxNoteItemList = a.FailAddCheckboxNoteItemList[:0]
			a.log.Debug("FailAddCheckboxNoteItemList size after clear : ", len(a.FailAddCheckboxNoteItemList))
		}
	}

	if len(a.FailUpdateCheckboxNoteItemList) > 0 {
		a.log.Debug("FailUpdateCheckboxNoteItemList: ", a.FailUpdateCheckboxNoteItemList)
		var suc = true
		a.retryCall = true
		var wg sync.WaitGroup
		for i := range a.FailUpdateCheckboxNoteItemList {
			wg.Add(1)
			go func(cb CheckboxNoteItem) {
				defer wg.Done()
				res := a.UpdateCheckboxItem(&cb)
				a.log.Debug("UpdateCheckboxItem suc: ", res.Success)
				if !res.Success {
					suc = false
				}
			}(a.FailUpdateCheckboxNoteItemList[i])
		}
		a.log.Debug("waiting for go routines in update checkitem")
		wg.Wait()
		a.retryCall = false
		if suc {
			a.FailUpdateCheckboxNoteItemList = a.FailUpdateCheckboxNoteItemList[:0]
			a.log.Debug("FailUpdateCheckboxNoteItemList size after clear : ", len(a.FailUpdateCheckboxNoteItemList))
		}
	}

	if len(a.FailAddNoteItemList) > 0 {
		a.log.Debug("FailAddNoteItemList: ", a.FailAddNoteItemList)
		var suc = true
		a.retryCall = true
		var wg sync.WaitGroup
		for i := range a.FailAddNoteItemList {
			wg.Add(1)
			go func(ni NoteItem) {
				defer wg.Done()
				res := a.AddNoteItem(&ni)
				a.log.Debug("AddNoteItem suc: ", res.Success)
				if !res.Success {
					suc = false
				}
			}(a.FailAddNoteItemList[i])
		}
		a.log.Debug("waiting for go routines in add item")
		wg.Wait()
		a.retryCall = false
		if suc {
			a.FailAddNoteItemList = a.FailAddNoteItemList[:0]
			a.log.Debug("FailAddNoteItemList size after clear : ", len(a.FailAddNoteItemList))
		}
	}

	if len(a.FailUpdateNoteItemList) > 0 {
		a.log.Debug("FailUpdateNoteItemList: ", a.FailUpdateNoteItemList)
		var suc = true
		a.retryCall = true
		var wg sync.WaitGroup
		for i := range a.FailUpdateNoteItemList {
			wg.Add(1)
			go func(ni NoteItem) {
				defer wg.Done()
				res := a.UpdateNoteItem(&ni)
				a.log.Debug("UpdateNoteItem suc: ", res.Success)
				if !res.Success {
					suc = false
				}
			}(a.FailUpdateNoteItemList[i])
		}
		a.log.Debug("waiting for go routines in update item")
		wg.Wait()
		a.retryCall = false
		if suc {
			a.FailUpdateNoteItemList = a.FailUpdateNoteItemList[:0]
			a.log.Debug("FailUpdateNoteItemList size after clear : ", len(a.FailUpdateNoteItemList))
		}
	}

}
