/*
Copyright © 2022 NAME HERE philip.hamid@zettagenomics.com

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	xurls "github.com/mvdan/xurls"
	"github.com/phamidko/opencga-cli/pkg/util"
	"github.com/spf13/cobra"
)

var (
	URL                                string = "iva.mseqdr.org"
	DEFAULT_CELLBASE_SERVICE_BASE_PATH        = "cellbase/"
	DEFAULT_OPENCGA_SERVICE_BASE_PATH         = "opencga/"
)

const (
	HTTP_REQUEST_TIMEOUT        = 2   // second
	INDEX_EXTRACT_SUBSTRING int = 174 // Arbitrary number to achieve performance
	CELLBASE_VERSION            = `v5`
	OPENCGA_VERSION             = `v2`
	IVA_CONFIG_FILE_PATH        = "/iva/conf/config.js"
	HTTPS_DEFAULT_PROTOCOL      = "https://"
	HTTPS_SUFFIX                = "/"

	DEFAULT_RESTAPI_SERVICE = "webservices/rest/" // "http://usa.ws.zettagenomics.com/cellbase/webservices/rest/v5/meta/about"
	REST_API_ENDPOINT       = "/meta/about"
)

func fetch(url string) []byte {

	spaceClient := http.Client{
		Timeout: time.Second * HTTP_REQUEST_TIMEOUT, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "opencga-cli")
	res, err := spaceClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body

}
func parse_string(s []string) []string {
	for i := 0; i < len(s); i++ {
		if !(strings.HasSuffix(s[i], HTTPS_SUFFIX)) {
			var t strings.Builder
			t.WriteString(s[i])
			t.WriteString(HTTPS_SUFFIX)
			s[i] = t.String()
		}

	}
	return s
}

func parse_struct(s []string) (c *util.Cellbase, o *util.Opencga) {
	cellbase := &util.Cellbase{}
	opencga := &util.Opencga{}
	var body []byte
	for i := 0; i < len(s); i++ {
		if strings.HasSuffix(s[i], DEFAULT_CELLBASE_SERVICE_BASE_PATH) {
			var t strings.Builder
			t.WriteString(s[i])
			t.WriteString(DEFAULT_RESTAPI_SERVICE)
			t.WriteString(CELLBASE_VERSION)
			t.WriteString(REST_API_ENDPOINT)
			s[i] = t.String()

			body = fetch(s[i])
			err := json.Unmarshal([]byte(body), cellbase)
			if err != nil {
				log.Fatalf("unable to parse value: %q, error: %s",
					string(body), err.Error())
			}
		} else if strings.HasSuffix(s[i], DEFAULT_OPENCGA_SERVICE_BASE_PATH) {
			var t strings.Builder
			t.WriteString(s[i])
			t.WriteString(DEFAULT_RESTAPI_SERVICE)
			t.WriteString(OPENCGA_VERSION)
			t.WriteString(REST_API_ENDPOINT)
			s[i] = t.String()
			body = fetch(s[i])
			err := json.Unmarshal([]byte(body), opencga)
			if err != nil {
				log.Fatalf("unable to parse value: %q, error: %s",
					string(body), err.Error())
			}
		}

	}
	return cellbase, opencga
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// jsonResponseCellbase := `{"apiVersion":"v5","time":2,"params":{"species":"do not validate species","limit":"10"},"responses":[{"time":0,"numResults":0,"results":[{"Program":"CellBase (OpenCB)","Git commit":"eaae3a6f7b407c1eebdb1b4bfede941f4b506b30","Description":"High-Performance NoSQL database and RESTful web services to access the most relevant biological data","Version":"5.0.1","Git branch":"release-5.0.x"}],"numTotalResults":0,"numMatches":0,"numInserted":0,"numUpdated":0,"numDeleted":0,"id":"about"}]}`
		// jsonResponseOpencga := `{"apiVersion":"v2","time":1,"events":[],"params":{},"type":"QUERY","responses":[{"time":0,"numResults":1,"results":[{"Program":"OpenCGA (OpenCB)","Git commit":"27cf2ae4bb95596daf839f107dac3d8fb63e6715","Description":"Big Data platform for processing and analysing NGS data","Version":"2.2.1-SNAPSHOT","Git branch":"release-2.2.x"}],"numMatches":0,"numInserted":0,"numUpdated":0,"numDeleted":0,"numErrors":0,"federationNode":{"id":"primary","uri":"http://opencga.mseqdr.org/opencga/webservices/rest/","commit":"27cf2ae4bb95596daf839f107dac3d8fb63e6715","version":"2.2.1-SNAPSHOT"},"numTotalResults":0}]}`

		// build URL "https://iva.mseqdr.org/iva/conf/config.js"
		var i strings.Builder
		i.WriteString(HTTPS_DEFAULT_PROTOCOL)
		i.WriteString(URL)
		i.WriteString(IVA_CONFIG_FILE_PATH)

		var res_cellbase *util.Cellbase
		var res_opencga *util.Opencga

		fmt.Printf("Fetching from %s\n", i.String())
		body_IVA := string(fetch(i.String()))

		var end = len([]rune(body_IVA))
		var start = end - INDEX_EXTRACT_SUBSTRING
		body_IVA_substring := body_IVA[start:end]

		rxStrict := xurls.Strict
		output := rxStrict.FindAllString(body_IVA_substring, -1)

		output = parse_string(output)
		res_cellbase, res_opencga = parse_struct(output)

		fmt.Printf("\tCellbase Version: %s \t\tGit Commit: %s \tGit Branch: %s\n", res_cellbase.Responses[0].Results[0].Version, res_cellbase.Responses[0].Results[0].GitCommit, res_cellbase.Responses[0].Results[0].GitBranch)
		fmt.Printf("\tOpenCGA Version: %s \tGit Commit: %s \tGit Branch: %s\n", res_opencga.Responses[0].Results[0].Version, res_opencga.Responses[0].Results[0].GitCommit, res_opencga.Responses[0].Results[0].GitBranch)

	},
}

func init() {
	getCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")
	// versionCmd.PersistentFlags().StringVarP(&Url, "url", "u", "https://iva.zettagenomics.com/iva/", "site URL")
	versionCmd.PersistentFlags().StringVarP(&URL, "site", "s", "iva.mseqdr.org", "site URL")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
