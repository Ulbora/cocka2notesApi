package cocka2notesapi

import (
	"sync"
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

// go mod init github.com/Ulbora/cocka2notesApi

const (
	post   = "POST"
	put    = "PUT"
	get    = "GET"
	delete = "DELETE"
)

//ResponseID ResponseID
type ResponseID struct {
	ID      int64  `json:"id"`
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//Response Response
type Response struct {
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//LoginResponse LoginResponse
type LoginResponse struct {
	Success bool   `json:"success"`
	Email   string `json:"email"`
}

//Headers Headers
type Headers struct {
	headers map[string]string
	mu      sync.Mutex
}

//Set Set
func (h *Headers) Set(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

//User User
type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	WebEnabled bool   `json:"webEnabled"`
}

//NoteUsers NoteUsers
type NoteUsers struct {
	OwnerEmail string `json:"ownerEmail"`
	NoteID     int64  `json:"noteId"`
	UserEmail  string `json:"userEmail"`
}

//TextNote TextNote
type TextNote struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	OwnerEmail string     `json:"ownerEmail"`
	NoteItems  []NoteItem `json:"noteTextItems"`
	LastUsed   time.Time  `json:"lastUsed"`
}

//Note Note
type Note struct {
	ID                int64              `json:"id"`
	Title             string             `json:"title"`
	Type              string             `json:"type"`
	OwnerEmail        string             `json:"ownerEmail"`
	NoteTextItems     []NoteItem         `json:"noteTextItems"`
	NoteCheckboxItems []CheckboxNoteItem `json:"noteCheckboxItems"`
	LastUsed          time.Time          `json:"lastUsed"`
}

//NoteItem NoteItem
type NoteItem struct {
	ID     int64  `json:"id"`
	Text   string `json:"text"`
	NoteID int64  `json:"noteId"`
}

//CheckboxNote CheckboxNote
type CheckboxNote struct {
	ID         int64              `json:"id"`
	Title      string             `json:"title"`
	Type       string             `json:"type"`
	OwnerEmail string             `json:"ownerEmail"`
	NoteItems  []CheckboxNoteItem `json:"noteCheckboxItems"`
	LastUsed   time.Time          `json:"lastUsed"`
}

//CheckboxNoteItem CheckboxNoteItem
type CheckboxNoteItem struct {
	ID      int64  `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
	NoteID  int64  `json:"noteId"`
}

//MailServer MailServer
type MailServer struct {
	ID          int64  `json:"id"`
	Host        string `json:"host"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	SenderEmail string `json:"senderEmail"`
}

//API API
type API interface {
	AddUser(u *User) *Response
	UpdateUser(u *User) *Response
	GetUser(email string) *User
	Login(u *User) *LoginResponse
	ResetPassword(u *User) *Response

	AddUserToNote(n *NoteUsers) *Response
	GetNoteUserList(noteID int64, ownerEmail string) *[]string

	AddNote(n *Note) *ResponseID
	UpdateNote(n *Note) *Response
	GetCheckboxNote(id int64) *CheckboxNote
	GetNote(id int64) *TextNote
	GetUsersNotes(email string) (*[]Note, bool)
	DeleteNote(id int64, ownerEmail string) *Response

	setSavedCheckboxNote(cb *CheckboxNote)
	getSavedCheckboxNote(id int64) *CheckboxNote
	setSavedTextNote(tx *TextNote)
	getSavedTextNote(id int64) *TextNote

	setSavedCheckboxItem(cb *CheckboxNoteItem)
	setSavedTextItem(cb *NoteItem)

	AddCheckboxItem(ni *CheckboxNoteItem) *ResponseID
	UpdateCheckboxItem(ni *CheckboxNoteItem) *Response
	DeleteCheckboxItem(id int64) *Response

	AddNoteItem(ni *NoteItem) *ResponseID
	UpdateNoteItem(ni *NoteItem) *Response
	DeleteNoteItem(id int64) *Response

	//Offline functionality below
	GetFailAddCheckboxNoteItemList() []CheckboxNoteItem
	GetFailUpdateCheckboxNoteItemList() []CheckboxNoteItem
	GetFailAddNoteItemList() []NoteItem
	GetFailUpdateNoteItemList() []NoteItem

	SetNoteList(noteList []Note)
}
