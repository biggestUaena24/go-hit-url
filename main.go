package main

import (
	"fmt"
	"log"
	"net/http"
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
	totalPages := getPages()

	for i:=0; i<totalPages; i++{
		getPage(i)
	}
}

func getPage(page int) []extractedJob{
	var jobResults []extractedJob
	pageURL := baseURL + "&recruitPage="+ strconv.Itoa(page + 1)
	
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res.StatusCode)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCards := doc.Find(".item_recruit")

	jobCards.Each(func (i int, s *goquery.Selection){
		id,_ := s.Attr("value")
		title, _ := s.Find(".area_job>.job_tit>a").Attr("title")
		location := ""
		location += s.Find(".area_job>.job_condition>span>a").Text()
		
		job := extractedJob{id: id, title: title, location:location}
		jobResults = append(jobResults, job)
		fmt.Println(job.id, job.title, job.location)
	})

	return jobResults
	

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
