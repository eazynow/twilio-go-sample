package main

import (
	//"encoding/json"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/eazynow/twilio-go/nouns"
	"github.com/eazynow/twilio-go/rest"
	"github.com/eazynow/twilio-go/verbs"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	accountsid = flag.String("accountsid", "", "The account sid to use")
	authtoken  = flag.String("authtoken", "", "The auth token to use")
	country    = flag.String("country", "", "The ISO3166 Country Code (GB, US, CA etc)")
)

func main() {
	flag.Parse()

	//primary()

	//secondary()

	restTest()
	fmt.Println("")
}

func restTest() {

	apibase := "https://api.twilio.com"
	apiversion := "2010-04-01"

	availuri := fmt.Sprintf("Accounts/%s/AvailablePhoneNumbers/%s/Local.json", url.QueryEscape(*accountsid), url.QueryEscape(*country))

	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", apibase, apiversion, availuri))
	q := u.Query()
	q.Set("Contains", "+441702")
	u.RawQuery = q.Encode()
	fmt.Println(u)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatalf("error building request: %s", err)
	}
	req.SetBasicAuth(*accountsid, *authtoken)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("wrong status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	numResponse := new(rest.AvailableNumbersResponse)
	err = decoder.Decode(numResponse)

	fmt.Println(numResponse)

	//-d "AreaCode=510" \
	//-u 'AC1ab10b47c2bb2d804d3dcee408ddf8ce:{AuthToken}'
}

func primary() {
	fmt.Println("\nPrimary verbs")

	say := verbs.Say{
		Voice:    "man",
		Language: "en-gb",
		Loop:     1,
		Text:     "Hello world!"}

	xmlout, err := xml.MarshalIndent(say, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	/* Add Play */
	play := verbs.Play{}

	gather := verbs.Gather{}

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.AddSay(say)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.AddPause(10)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.RemoveSay()

	gather.AddPlay(play)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	sip := nouns.Sip{}
	sip.Username = "user1"
	sip.Uri = "Test"
	sip.Password = "pass1"
	sip.Url = "http://www.test.com"
	sip.Method = "POST"

	xmlout, err = xml.MarshalIndent(sip, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")
}

func secondary() {
	fmt.Println("\nSecondary verbs")

	v := verbs.Enqueue{
		WaitUrl:       "http://www.wait.com",
		WaitUrlMethod: "POST",
		QueueName:     "newqueue"}
	v.Action = "http://www.something.com"
	v.Method = "POST"

	xmlout, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Leave{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Hangup{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Redirect{Url: "../nextInstructions", Method: "POST"}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Reject{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Reject{Reason: "not sure"}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Pause{Length: 30}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

}
