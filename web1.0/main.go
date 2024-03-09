package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/carloseabade/web1.0/model"
)

var (
	contacts = model.NewContactsSet()
	message  = ""
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	})

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		query := getQuery(r)

		values := map[string]any{
			"Messages": getMessageIfAny(),
			"Content": map[string]any{
				"ContactsSet": getContactsSet(query),
				"Query":       query,
			},
		}

		renderTemplate(w, values, "templates/index.html", "templates/content.html")
	})

	http.HandleFunc("/contacts/new", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			values := map[string]any{
				"Content": map[string]any{
					"Contact": nil,
					"Error":   nil,
				},
			}
			renderTemplate(w, values, "templates/index.html", "templates/new.html")
		case http.MethodPost:
			r.ParseForm()
			c := model.NewContact(r.Form.Get("first_name"), r.Form.Get("last_name"), r.Form.Get("phone"), r.Form.Get("email"))
			err := contacts.Add(c)
			if err != nil {
				values := map[string]any{
					"Content": map[string]any{
						"Contact": *c,
						"Error":   err,
					},
				}
				renderTemplate(w, values, "templates/index.html", "templates/new.html")
			}
			createMessage("Created New Contact!")
			http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		}
	})

	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w io.Writer, data any, filename ...string) {
	tmpl := template.Must(template.ParseFiles(filename...))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func getQuery(r *http.Request) string {
	r.ParseForm()
	return r.Form.Get("q")
}

func getContactsSet(query string) model.ContactsSet {
	contactsSet := model.ContactsSet{}
	if len(query) > 0 {
		contactsSet = contacts.SearchByName(query)
	} else {
		contactsSet = contacts.All()
	}
	return contactsSet
}

func createMessage(msg string) {
	message = msg
}

func getMessageIfAny() string {
	defer func() {
		message = ""
	}()
	return message
}
