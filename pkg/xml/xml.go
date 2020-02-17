package xml

import (
	"encoding/xml"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/", Foo)
	http.ListenAndServe(":3000", nil)
}

//Foo function
func Foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	x, err := xml.MarshalIndent(profile, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
