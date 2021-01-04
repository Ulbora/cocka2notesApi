package cocka2notesapi

import (
	"encoding/json"
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
