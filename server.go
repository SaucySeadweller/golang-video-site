package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

//RadioButtons defines the buttons available to the user
type RadioButtons struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

//PageVariables defines the variables to be displayed on the webpage
type PageVariables struct {
	PageTitle         string
	PageRadioButtonss []RadioButtons
	Answer            string
}

func main() {
	//	http.HandleFunc("/", DisplayButtons)
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
	}
}

//DisplayButtons creates the buttons that the user clicks on to make their choice.
//func DisplayButtons(w http.ResponseWriter, r *http.Request) {
//	Title := "Video Player"
//	MyRadioButtonss := []RadioButtons{}
//	MyPageVariables := PageVariables{
//		PageTitle:         Title,
//		PageRadioButtonss: MyRadioButtonss,
//	}
//	t, err := template.ParseFiles("site.html")
//	if err != nil {
//		log.Print("Template parsing error: ", err)
//	}
//	err = t.Execute(w, MyPageVariables)
//	if err != nil {
//		log.Print("Template executing error: ", err)
//	}
//}
