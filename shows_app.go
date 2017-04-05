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

func newHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>New Show</h1><form action='/shows' method='POST'><input type='text' name='artist' placeholder='artist'><br>year<input type='number' name='year' min='1950' max='2020'><br>month<input type='number' name='month' min='1' max='12'><br>day<input type='number' name='day' min='1' max='31'><br><input type='text' name='city' placeholder='city'><br><input type='text' name='venue' placeholder='venue'><br><input type='submit'></form>")
}

func saveHandler(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println()

  var s Show
  s.artist = r.Form["artist"][0]
  s.year, _ = strconv.Atoi(r.Form["year"][0])
  s.month, _ = strconv.Atoi(r.Form["month"][0])
  s.day, _ = strconv.Atoi(r.Form["day"][0])
  s.city = r.Form["city"][0]
  s.venue = r.Form["venue"][0]

  shows = append(shows, s)

  http.Redirect(w, r, "/shows", 301)
}

func main() {
  r := mux.NewRouter()
  shows = append(shows, Show{artist: "Phish", year: 2017, month: 4, day: 20, city: "Guelph", venue: "Up Your Friggin Dick"})
  shows = append(shows, Show{artist: "The Disco Biscuits", year: 2017, month: 5, day: 21, city: "Guelph", venue: "Up Your Friggin Balls"})
  shows = append(shows, Show{artist: "Creed", year: 2017, month: 6, day: 22, city: "Guelph", venue: "Up Your Friggin Ass"})

  r.HandleFunc("/shows", showsHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}", showsArtistHandler).Methods("GET")
  r.HandleFunc("/shows/{artist}/{year}", showsArtistYearHandler).Methods("GET")
  r.HandleFunc("/new", newHandler).Methods("GET")
  r.HandleFunc("/shows", saveHandler).Methods("POST")

  log.Fatal(http.ListenAndServe(":8000", r))
}
