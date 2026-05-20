package output

import (
	"encoding/json"
	"fmt"
	"ipcli/internal/models"
)

func Print(data *models.ApiResponse, query string) {
	title := query
	if title == "" {
		title = "Current IP"
	}

	fmt.Println("================================")
	fmt.Printf("        IP Lookup: %s\n", title)
	fmt.Println("================================")

	fmt.Printf("%-10s = %s\n", "IP", data.Ip)
	fmt.Printf("%-10s = %s\n", "Country", data.Country)
	fmt.Printf("%-10s = %s\n", "Region", data.Region)
	fmt.Printf("%-10s = %s\n", "City", data.City)
	fmt.Printf("%-10s = %.4f\n", "Latitude", data.Lat)
	fmt.Printf("%-10s = %.4f\n", "Longitude", data.Lon)
	fmt.Printf("%-10s = %s\n", "ISP", data.Isp)

	fmt.Println("================================")
}

func JsonOutput(data *models.ApiResponse) {
	b, _ := json.MarshalIndent(data, "", "")
	fmt.Println(string(b))
}
