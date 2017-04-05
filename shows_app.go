package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "strings"
)

type Show struct {
  artist string
  year int
  month int
  day int
  city string
  venue string
}

var allShows = make(map[string][]Show)

func showsHandler(w http.ResponseWriter, r *http.Request){
  for _, item := range allShows{
    for _, item := range item{
      fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
    }
  }
}

func showsArtistHandler(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  for _, item := range allShows[strings.Title(params["artist"])]{
    fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
  }
}

func showsArtistYearHandler(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  yearParam, _ := strconv.Atoi(params["year"])
  for _, item := range allShows[strings.Title(params["artist"])]{
    if item.year == yearParam {
      fmt.Fprintf(w, "<h1>%s</h1><div>%v/%v/%v</div><div>%s</div><div>%s</div>",item.artist, item.year, item.month, item.day, item.city, item.venue)
    }
  }
}

func newHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>New Show</h1><form action='/shows' method='POST'><input type='text' name='artist' placeholder='artist'><br>year<input type='number' name='year' min='1950' max='2020'><br>month<input type='number' name='month' min='1' max='12'><br>day<input type='number' name='day' min='1' max='31'><br><input type='text' name='city' placeholder='city'><br><input type='text' name='venue' placeholder='venue'><br><input type='submit'></form>")
}

func saveHandler(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println()

  var s Show
  s.artist = strings.Title(r.Form["artist"][0])
  s.year, _ = strconv.Atoi(r.Form["year"][0])
  s.month, _ = strconv.Atoi(r.Form["month"][0])
  s.day, _ = strconv.Atoi(r.Form["day"][0])
  s.city = strings.Title(r.Form["city"][0])
  s.venue = strings.Title(r.Form["venue"][0])

  allShows[s.artist] = append(allShows[s.artist], s)

  http.Redirect(w, r, "/shows", 301)
}

func main() {
  r := mux.NewRouter()
  allShows["Phish"] = append(allShows["Phish"], Show{artist: "Phish", year: 2017, month: 4, day: 20, city: "Guelph", venue: "Up Your Friggin Dick"})
  allShows["The Disco Biscuits"] = append(allShows["The Disco Biscuits"], Show{artist: "The Disco Biscuits", year: 2017, month: 5, day: 21, city: "Guelph", venue: "Up Your Friggin Balls"})
  allShows["Creed"] = append(allShows["Creed"], Show{artist: "Creed", year: 2017, month: 6, day: 22, city: "Guelph", venue: "Up Your Friggin Ass"})

  r.HandleFunc("/shows", showsHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}", showsArtistHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}/{year}", showsArtistYearHandler).Methods("GET")
  r.HandleFunc("/new", newHandler).Methods("GET")
  r.HandleFunc("/shows", saveHandler).Methods("POST")

  log.Fatal(http.ListenAndServe(":8000", r))
}
