package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	start := time.Now()
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
	elapsed := time.Since(start)
	fmt.Println("total time", elapsed)
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := card.Find(".jobTitle").Text()
	location := card.Find(".companyLocation").Text()
	salary := card.Find(".salary-snippet").Text()
	summary := card.Find(".job-snippet").Text()
	if salary == "" {
		salary = "면접 후 협의"
	}

	c <- extractedJob{
		id:       cleanString(id),
		title:    cleanString(title),
		location: cleanString(location),
		salary:   cleanString(salary),
		summary:  cleanString(summary)}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
		pages = getPageNum(cleanString(s.Text()))
	})
	return pages
}

func getPageNum(str string) int {
	r, _ := regexp.Compile("[0-9]")
	pageList := r.FindAllString(str, -1)[1:]
	totalResults, _ := strconv.Atoi(strings.Join(pageList, ""))
	itemsPerPage := 50
	pageNum := totalResults / itemsPerPage
	if totalResults%itemsPerPage != 0 {
		pageNum += 1
	}
	return pageNum
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		joblink := ""
		if job.id != "" {
			joblink = "https://kr.indeed.com/viewjob?jk=" + job.id
		}
		jobSlice := []string{joblink, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.Status)
	}
}
