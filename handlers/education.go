package handlers

import (
	"eth2-exporter/utils"
	"html/template"
	"net/http"
)

var educationServicesTemplate = template.Must(template.New("educationServices").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/educationServices.html"))

func EducationServices(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/education", "Ethereum 2.0 Education Services Overview")
	data.Meta.Title = "Ethereum 2.0 Education Materials"
	data.Meta.Description = "New to Ethereum 2.0 and want to know more? Here's the list of all education materials related to the ETH2 network and how it works."

	err = educationServicesTemplate.ExecuteTemplate(w, "layout", data)
	if err != nil {
		logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", 503)
		return
	}
}
