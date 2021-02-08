package main

import (
  "os"
  "flag"
  "net/http"
  "fmt"
  "log"
  "io/ioutil"
  "encoding/json"
  "time"
  "strings"
  "sync"
)

type Cert struct {
  IssuerCaId int
  IssuerName string `json:"issuer_name"`
  CommonName string `json:"common_name"`
  NameValue string `json:"name_value"`
  Id int
  EntryTimestamp time.Time
  NotBefore time.Time
  NotAfter time.Time
  SerialNumber string `json:"serial_number"`
}

var subdomains []string

func main() {
  var wg sync.WaitGroup
  wg.Add(1)

  targetFlag := flag.String("t", "", "the target")
  flag.Parse()

  if *targetFlag == "" {
    fmt.Printf("\nYou must provide a target e.g., -t http://example.com\n\n")
    os.Exit(1)
  }

  GetCerts(*targetFlag, &wg)

  for _, subdomain := range subdomains {
    wg.Add(1)
    go GetCerts(strings.Replace(subdomain, "*.", "", -1), &wg)
  }

  wg.Wait()

  Output(subdomains)
}

func GetCerts(target string, wg *sync.WaitGroup) {
  defer wg.Done()

  // make a sample HTTP GET request
  url := fmt.Sprintf("https://crt.sh/?Identity=%%.%s&output=json", target)
  res, err := http.Get(url)

  // check for response error
  if err != nil {
    log.Fatal( err )
  }

  // read all response body
  data, _ := ioutil.ReadAll( res.Body )

  // close response body
  res.Body.Close()

  // unmarshall and print data
  var certs []Cert

  json.Unmarshal([]byte(data), &certs)

  subdomains = append(subdomains, RemoveDuplicatesFromSlice(certs)...)

}

func Recurse() {
  // not implemented
}

func Output(results []string) {
  for _, element := range results {
    if len(element) != 0 {
      fmt.Println(strings.Replace(element, "*.", "", -1))
    }
  }
}

func RemoveDuplicatesFromSlice(c []Cert) []string {
  m := make(map[string]bool)
  for _, item := range c {
    if _, ok := m[item.CommonName]; ok {
    } else {
      m[item.CommonName] = true
    }
  }

  var result []string
  for item := range m {
    result = append(result, item)
  }
  return result
}