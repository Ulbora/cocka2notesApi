package cocka2notesapi

import ()

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

//GetMailServer GetMailServer
func (a *NotesAPI) GetMailServer() *MailServer {
	var rtn MailServer
	//idStr := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/mailserver/get"
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, a.headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)
	// if suc && stat == 200 {
	// 	a.Note = rtn
	// } else {
	// 	rtn = a.Note
	// }
	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
