package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

// RecentCamps represents JSON data for /api/recentcampaigns
type RecentCamps struct {
	Fromindex        string `json:"fromindex"`
	TotalRecordCount string `json:"total_record_count"`
	Code             string `json:"code"`
	RecentCampaigns  []struct {
		CampaignKey string `json:"campaign_key"`
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
		LastSeen            string `json:"last seen yymmdd"`
		Title               string `json:"title"`
	} `json:"list_of_details"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type parser func(i interface{}) interface{}

func getResponse(url string) parser {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Unable to fetch from"+url+" with error :", err)
	}

	// Decode Response
	return func(i interface{}) interface{} {
		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&i); err != nil {
			log.Fatal("Unable to Decode data "+"with error :", err)
		}
		return i
	}
}

func main() {
	var err error
	t := time.Now()
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
	header := []string{"CampaignID", "ContactEmailAddress", "LastSeen", "Title"}
	if err = writer.Write(header); err != nil {
		log.Warn("unable to write header", err)
	}

	// Url params
	url := "https://campaigns.zoho.com/api/recentcampaigns"
	authtoken := "9ea6d981d30cd3880d72f618bf475a05"
	status := "all"
	fromIndex := "1"
	upTo := "10"
	action := "openedcontacts"

	log.WithFields(log.Fields{
		"Url":       url,
		"AuthToken": authtoken,
		"Status":    status,
		"FromIndex": fromIndex,
		"Range":     upTo,
		"Action":    action,
	}).Info("Fetching Response with following data;")

	req := url +
		"?authtoken=" + authtoken +
		"&scope=CampaignsAPI&status=" + status +
		"&fromindex=" + fromIndex +
		"&resfmt=JSON&range=" + upTo

	// Get Recent Campaigns
	p := getResponse(req)
	recent := p(&RecentCamps{}).(*RecentCamps)
	// var recent RecentCamps
	if recent.Status == "error" {
		log.Fatal("Invalid request " + "with error :" + recent.Message)
	}

	log.Println(recent)
	// Get Campaign Data
	for _, camp := range recent.RecentCampaigns {
		url = "https://campaigns.zoho.com/api/getcampaignrecipientsdata"

		req := url +
			"?authtoken=" + authtoken +
			"&campaignkey=" + camp.CampaignKey +
			"&action=" + action +
			"&scope=CampaignsAPI&resfmt=JSON"

		// Fetching from response
		parse := getResponse(req)
		campData := parse(&CampData{}).(*CampData)

		if campData.ListOfDetails == nil {
			log.WithFields(log.Fields{
				"CampaingnID": camp.CampaignKey,
			}).Warn("No Records Found")
			continue
		}

		// Appending data
		var records [][]string
		for _, data := range campData.ListOfDetails {
			record := []string{camp.CampaignKey, data.ContactEmailAddress, data.LastSeen, data.Title}
			records = append(records, record)

		}

		log.Info("Encoding data....")
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
	fmt.Println(time.Now().Sub(t).Seconds())
}
