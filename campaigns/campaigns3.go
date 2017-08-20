/*
	Command Line application to fetch results from zoho Campaigns to get desired results

		Build the application
			- go build campaings.go

		Run the application
			- ./campaigns -action=openedcontacts -fromIndex=1 -range=10 -status=all

		Type help for usage
			- ./campaigns --help

		Usage of ./campaigns:
		-action string
				Enter the action required (default "openedcontacts")
		-fromIndex int
				Enter index of the campaign (default 1)
		-range int
				Enter no of campaigns to get (default 10)
		-status string
				Enter the status of campaign (default "all")

	The output would be saved to campaigns.csv
*/

package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// RecentCamps represents JSON data for /api/recentcampaigns
type RecentCamps struct {
	Fromindex        string `json:"fromindex"`
	TotalRecordCount string `json:"total_record_count"`
	Code             string `json:"code"`
	RecentCampaigns  []struct {
		CampaignStatus string `json:"campaign_status"`
		CampaignKey    string `json:"campaign_key"`
	} `json:"recent_campaigns"`
	Range         string `json:"range"`
	CampaignCount string `json:"campaign_count"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

// CampData represents JSON data for individual Campaigns
type CampData struct {
	Code          string `json:"code"`
	ListOfDetails []struct {
		ContactEmailAddress string `json:"contactemailaddress"`
		Sentdate            string `json:"sentdate"`
		LastSeen            string `json:"last seen yymmdd"`
		Title               string `json:"title"`
	} `json:"list_of_details"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {

	// Command line flags
	upTo := flag.Int("range", 10, "Enter no of campaigns to get")
	from := flag.Int("fromIndex", 1, "Enter index of the campaign")
	status := flag.String("status", "all", "Enter the status of campaign")
	action := flag.String("action", "openedcontacts", "Enter the action required")
	flag.Parse()

	// Create File
	fname := "campaigns.csv"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatal("Unable to create File, error :", err)

	}
	defer file.Close()

	// CSV Encoder
	writer := csv.NewWriter(file)

	// Write Header
	header := []string{"CampaignID", "CampaignStatus", "ContactEmailAddress", "SentDate", "LastSeen", "Title"}
	if err = writer.Write(header); err != nil {
		log.Warn("unable to write header", err)
	}

	// Url params
	url := "https://campaigns.zoho.com/api/recentcampaigns"
	authtoken := "9ea6d981d30cd3880d72f618bf475a05"

	log.WithFields(log.Fields{
		"Url":       url,
		"AuthToken": authtoken,
		"Status":    *status,
		"FromIndex": *from,
		"Range":     *upTo,
		"Action":    *action,
	}).Info("Fetching Response with following data;")

	// Get Recent Campaigns
	resp, err := http.Get(url +
		"?authtoken=" + authtoken +
		"&scope=CampaignsAPI&status=" + *status +
		"&fromindex=" + strconv.Itoa(*from) +
		"&resfmt=JSON&range=" + strconv.Itoa(*upTo))

	if err != nil {
		log.Fatal("Unable to fetch from /api/recentcampaigns, error :", err)
	}
	defer resp.Body.Close()

	// Decode Response
	var recent RecentCamps
	if err = json.NewDecoder(resp.Body).Decode(&recent); err != nil {
		log.Fatal("Unable to Decode response, error ", err)
	}

	if recent.Status == "error" {
		log.Fatal("Unable to fetch from the url: ", url, "error msg: ", recent.Message)
	}

	// Get Campaign Data

	url = "https://campaigns.zoho.com/api/getcampaignrecipientsdata"
	log.WithFields(log.Fields{
		"Url": url,
	}).Info("Fetching response from ")

	for _, camp := range recent.RecentCampaigns {
		var campData CampData
		// Skip the campaign with campaign status == draft
		if strings.Compare(camp.CampaignStatus, "Draft") == 0 {
			continue
		}
		// Fetching from response
		resp, err := http.Get(url +
			"?authtoken=" + authtoken +
			"&campaignkey=" + camp.CampaignKey +
			"&action=" + *action +
			"&scope=CampaignsAPI&resfmt=JSON")

		if err != nil {
			log.Fatal("Unable to fetch from /api/getcampaignrecipientsdata, error", err)
			continue
		}
		defer resp.Body.Close()

		// Decoding response
		if err = json.NewDecoder(resp.Body).Decode(&campData); err != nil {
			log.Fatal("Unable to decode from /api/getcampaignrecipientsdata", err)
			continue
		}

		if campData.ListOfDetails == nil {
			log.WithFields(log.Fields{
				"CampaingnID": camp.CampaignKey,
			}).Warn("No Records Found")
			continue
		}
		// Appending data
		var records [][]string
		for _, data := range campData.ListOfDetails {
			record := []string{camp.CampaignKey, camp.CampaignStatus, data.ContactEmailAddress, data.Sentdate, data.LastSeen, data.Title}
			records = append(records, record)

		}

		// Encoding to CSV file
		if err = writer.WriteAll(records); err != nil {
			log.Println("Unable to encode to csv, with err ", err)
			continue
		}

	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal("Unable to write to csv file ", err)
	}

	log.Info("Response successfully encoded to file ", fname)
}
