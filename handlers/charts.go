package handlers

import (
	"eth2-exporter/services"
	"eth2-exporter/types"
	"eth2-exporter/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var chartsTemplate = template.Must(template.New("charts").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/charts.html"))
var genericChartTemplate = template.Must(template.New("chart").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/genericchart.html"))
var chartsUnavailableTemplate = template.Must(template.New("chart").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/chartsunavailable.html"))
var slotVizTemplate = template.Must(template.New("slotViz").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/slotViz.html"))

func getChartMeta(chartVar string) (string, string) {
	switch chartVar {
	case "blocks":
		return "Ethereum 2.0 Blocks Chart - History of Daily Blocks Proposed",
			"This chart illustrates the history of daily blocks proposed on the Ethereum 2.0 network."
	case "validators":
		return "Ethereum 2.0 Validators Chart - Active Validators Statistics",
			"This chart illustrates the history of daily active validators on the Ethereum 2.0 network."
	case "staked_ether":
		return "Ethereum 2.0 Staked Ether Chart",
			"This chart illustrates the history of daily staked Ether, which is the sum of all effective balances."
	case "average_balance":
		return "Ethereum 2.0 Validator Balance Chart",
			"This chart illustrates the average daily validator balance on the Ethereum 2.0 network."
	case "network_liveness":
		return "Ethereum 2.0 Network Liveness Chart",
			"This chart illustrates the Ethereum 2.0 network liveness."
	case "participation_rate":
		return "Ethereum 2.0 Participation Rate Chart",
			"This chart illustrates how many of the validators expected to attest to blocks are actually doing so."
	case "stake_effectiveness":
		return "Ethereum 2.0 Stake Effectiveness Chart",
			"This chart illustrates the relation between the sum of all effective balances and the sum of all balances."
	case "balance_distribution":
		return "Ethereum 2.0 Balance Distribution Chart",
			"This chart illustrates balance distribution in the current Ethereum 2.0 epoch."
	case "effective_balance_distribution":
		return "Ethereum 2.0 Effective Balance Distribution Chart",
			"This chart illustrates the effective balance distribution in the current Ethereum 2.0 epoch."
	case "performance_distribution_365d":
		return "Ethereum 2.0 Income Distribution Chart",
			"This chart illustrates the current state of income performances of the last 365 days."
	case "deposits_distribution":
		return "Eth1 Deposit Distribution Chart",
			"This chart illustrates validator distributions by Eth1 deposit addresses."
	case "deposits":
		return "ETH2 and ETH1 Daily Deposit Chart",
			"This chart illustrates the daily amount of deposited ETH."
	default:
		return "Ethereum 2.0 Network Charts | Ethscan.org",
			"Ethereum chart shows Ethereum statistics, ETH2 validator data and other info pertaining to Ethereum network. Ether chart for professionals and beginners."
	}
}

func getChartHeader(chartVar string) string {
	switch chartVar {
	case "blocks":
		return "Ethereum 2.0 Blocks Chart"
	case "validators":
		return "Ethereum 2.0 Validators Chart"
	case "staked_ether":
		return "Staked Ether Chart"
	case "average_balance":
		return "Validator Balance Chart"
	case "network_liveness":
		return "Network Liveness Chart"
	case "participation_rate":
		return "Network Liveness Chart"
	case "stake_effectiveness":
		return "ETH2 Stake Effectiveness Chart"
	case "balance_distribution":
		return "ETH2 Balance Distribution Chart"
	case "effective_balance_distribution":
		return "ETH2 Effective Balance Distribution Chart"
	case "performance_distribution_365d":
		return "ETH2 Income Distribution Chart"
	case "deposits_distribution":
		return "ETH1 Deposit Distribution Chart"
	case "deposits":
		return "ETH2 and ETH1 Daily Deposit Chart"
	default:
		return "Charts from the Ethereum 2.0 Network"
	}
}

// Charts uses a go template for presenting the page to show charts
func Charts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "stats", "/charts", "Charts")
	vars := mux.Vars(r)
	data.Meta.Title, data.Meta.Description = getChartMeta(vars["chart"])

	chartsPageData := services.LatestChartsPageData()
	if chartsPageData == nil {
		err := chartsUnavailableTemplate.ExecuteTemplate(w, "layout", data)
		if err != nil {
			logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
			http.Error(w, "Internal server error", 503)
			return
		}
		return
	}

	data.Data = chartsPageData
	chartsTemplate = template.Must(template.New("charts").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/charts.html"))
	err := chartsTemplate.ExecuteTemplate(w, "layout", data)
	if err != nil {
		logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// Chart renders a single chart
func Chart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chartVar := vars["chart"]
	switch chartVar {
	case "slotviz":
		SlotViz(w, r)
	default:
		GenericChart(w, r)
	}
}

// GenericChart uses a go template for presenting the page of a generic chart
func GenericChart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chartVar := vars["chart"]

	w.Header().Set("Content-Type", "text/html")
	data := InitPageData(w, r, "stats", "/charts", "Chart")
	data.Meta.Title, data.Meta.Description = getChartMeta(chartVar)

	chartsPageData := services.LatestChartsPageData()
	if chartsPageData == nil {
		err := chartsUnavailableTemplate.ExecuteTemplate(w, "layout", data)
		if err != nil {
			logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
			http.Error(w, "Internal server error", 503)
			return
		}
		return
	}

	var chartData *types.GenericChartData
	for _, d := range *chartsPageData {
		if d.Path == chartVar {
			chartData = d.Data
			break
		}
	}

	if chartData == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	chartData.H1 = getChartHeader(chartVar)
	data.Meta.Path = "/charts/" + chartVar
	data.Data = chartData

	genericChartTemplate = template.Must(template.New("chart").Funcs(utils.GetTemplateFuncs()).ParseFiles("templates/layout.html", "templates/genericchart.html"))
	err := genericChartTemplate.ExecuteTemplate(w, "layout", data)
	if err != nil {
		logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// SlotViz renders a single page with a d3 slot (block) visualisation
func SlotViz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := InitPageData(w, r, "stats", "/charts", "Charts")
	data.Meta.Title = "Ethereum 2.0 Network Charts | Ethscan.org"
	data.Meta.Description = "Ethereum chart shows Ethereum statistics, ETH2 validator data and other info pertaining to Ethereum network. Ether chart for professionals and beginners."

	data.Data = nil
	err := slotVizTemplate.ExecuteTemplate(w, "layout", data)
	if err != nil {
		logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
