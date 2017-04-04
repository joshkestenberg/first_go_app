package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

type Show struct {
  artist string
  year int
  month int
  day int
  city string
  venue string
}

var shows []Show

func showsHandler(w http.ResponseWriter, r *http.Request){
  for _, item := range shows{
    fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
  }
}

func showsArtistHandler(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  for _, item := range shows{
    if item.artist == params["artist"] {
      fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
    }
  }
}

func showsArtistYearHandler(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  yearParam, _ := strconv.Atoi(params["year"])
  for _, item := range shows{
    if item.artist == params["artist"] && item.year == yearParam {
      fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
    }
  }
}

func showsCreateHandler(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  yearParam, _ := strconv.Atoi(params["year"])
  monthParam, _ := strconv.Atoi(params["month"])
  dayParam, _ := strconv.Atoi(params["day"])

  var s Show
  s.artist = params["artist"]
  s.year = yearParam
  s.month = monthParam
  s.day = dayParam
  s.city = params["city"]
  s.venue = params["venue"]

  shows = append(shows, s)

  http.Redirect(w, r, "localhost:8000/shows", 301)
}

func main() {
  r := mux.NewRouter()
  shows = append(shows, Show{artist: "Phish", year: 2017, month: 4, day: 20, city: "Guelph", venue: "Up Your Friggin Dick"})
  shows = append(shows, Show{artist: "The Disco Biscuits", year: 2017, month: 5, day: 21, city: "Guelph", venue: "Up Your Friggin Balls"})
  shows = append(shows, Show{artist: "Creed", year: 2017, month: 6, day: 22, city: "Guelph", venue: "Up Your Friggin Ass"})

  r.HandleFunc("/shows", showsHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}", showsArtistHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}/{year}", showsArtistYearHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}/{year}/{month}/{day}/{city}/{venue}", showsCreateHandler).Methods("POST")

  log.Fatal(http.ListenAndServe(":8000", r))
}
