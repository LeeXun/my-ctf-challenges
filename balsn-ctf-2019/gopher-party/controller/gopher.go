package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var gophers = []string{
	"https://www.youtube.com/watch?v=OOGTd1vejKw",
	"https://youtu.be/OOGTd1vejKw?t=1",
	"https://youtu.be/OOGTd1vejKw?t=7",
	"https://www.youtube.com/watch?v=IN5xcEbJftg",
}

func init() {
	rand.Seed(time.Now().Unix())
}

func gogopher(w http.ResponseWriter, m string) {
	w.Header().Set("REFRESH", fmt.Sprintf("1;URL=%v", gophers[rand.Intn(len(gophers))]))
	fmt.Fprintf(w, m)
}
