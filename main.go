package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
}
var baseURL = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"


func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()

	for i:=0; i<totalPages; i++{
		go getPage(i, c)
	}

	for i:=0; i<totalPages;i++{
		jobs = append(jobs, <-c...)
	}

	writeJobs(jobs)
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	defer w.Flush()
	headers := []string{"job link", "title", "location"}
	checkErr(w.Write(headers))

	for _, job := range(jobs){
		info := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?rec_idx=" + job.id, " " + job.title, job.location}
		checkErr(w.Write(info))
	}
}

func getPage(page int, mainChannel chan[]extractedJob){
	var jobResults []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage="+ strconv.Itoa(page + 1)
	fmt.Println("Extracting from URL: ", pageURL)
	
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res.StatusCode)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCards := doc.Find(".item_recruit")

	jobCards.Each(func (i int, s *goquery.Selection){
		go extractJob(s, c)
	})

	for i:=0; i < jobCards.Length();i++{
		jobResults = append(jobResults, <-c)
	}

	mainChannel <- jobResults
}

func extractJob (s *goquery.Selection, c chan extractedJob){

	id,_ :=  s.Attr("value")
	title, _ := s.Find(".area_job>.job_tit>a").Attr("title")
	location := ""
	location += s.Find(".area_job>.job_condition>span>a").Text()
	
	c <- extractedJob{id: id, title: title, location:location}


}

func getPages() int {
	var pages = 0
	res, err := http.Get(baseURL);
	checkErr(err)
	checkCode(res.StatusCode)
	
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(code int){
	if code != 200 {
		log.Fatal("Request failed due to status code: ", code)
	}
}
