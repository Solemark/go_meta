package clientRouter

import (
	"log"
	"net/http"
	"strconv"
)

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	var cl []Client = getClientsList()
	var id string = ""
	var fn string = ""
	var ln string = ""
	var ea string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "client_id" {
			id = value[0]
		}
		if key == "first_name" {
			fn = value[0]
		}
		if key == "last_name" {
			ln = value[0]
		}
		if key == "email_address" {
			ea = value[0]
		}
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	cl[i-1] = Client{FirstName: fn, LastName: ln, EmailAddress: ea, Visible: true}
	writeClientList(cl)
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}
