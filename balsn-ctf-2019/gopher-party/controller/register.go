package controller

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode"

	"balsnctf/gopherparty/config"
	"balsnctf/gopherparty/event"
	"balsnctf/gopherparty/solver"
	"balsnctf/gopherparty/store"

	"github.com/go-redis/redis"
	"golang.org/x/text/language"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Register should have comment or be unexported
func Register(sp *store.Store, c *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("True-Client-IP")
		gopher := "(O..O)!!!\n"
		if event.IsDown() {
			gogopher(w, "(O口o)!!!\n")
			return
		}
		if r.Method != "POST" {
			return
		}
		t := r.FormValue("access_token")
		if t == "" {
			gogopher(w, "(o口o)!!!\n")
			return
		}
		user, err := getUser(t)
		if err != nil {
			log.Printf("IP: %v, Err: %v\n", ip, err)
			gogopher(w, "(O口O)!!!\n")
			return
		}
		chosen := fmt.Sprintf("%v-%v", user.Email, time.Now())
		sp.Set("chosen", chosen)
		log.Printf("Start: IP: %v, Name: %v, Email: %v, Name2: %v, Age: %v, Praise: %v, Prove: %v, Language: %v\n",
			ip,
			user.Name,
			user.Email,
			r.FormValue("name"),
			r.FormValue("age"),
			r.FormValue("praise"),
			r.FormValue("prove"),
			r.Header.Get("Accept-Language"),
		)

		check := "true"
		emailKey := config.Redis.UserEmailKey + user.Email
		if os.Getenv("APP_ENV") == config.App.Production {
			if val, ok := sp.Get(emailKey); ok || val == check {
				log.Printf("IP: %v, Msg: Throttled\n", ip)
				gogopher(w, gopher)
				return
			}
		}
		sp.Set(emailKey, check)

		interest := r.FormValue("interest")
		if interest != "AH!" {
			gogopher(w, gopher)
			return
		}

		if aa, bb := parseName(user.Name, c, sp); aa == 0 && bb == 0 {
			log.Printf("IP: %v, Name: %v\n", ip, user.Name)
		}

		name := r.FormValue("name")
		if name == user.Name || name == "" {
			pic, err := getUserPicture(t)
			if err != nil {
				log.Printf("IP: %v, Err: %v, Pic: %v\n", ip, err, pic)
			}
		}

		prove := r.FormValue("prove")
		sc := make(chan string)
		var b [1 << 10]byte
		if prove == "" {
			go func() {
				copy(b[:], prove)
				sc <- base64.StdEncoding.EncodeToString(b[:])
			}()
		} else {
			copy(b[:], prove)
			sc <- base64.StdEncoding.EncodeToString(b[:])
		}
		select {
		case prove = <-sc:
			log.Printf("IP: %v, Err: %v, Prove: %v\n", ip, err, prove[:10])
		case <-time.After(time.Millisecond * 1):
			log.Printf("IP: %v, Err: %v, Prove: %v\n", ip, err, prove[:10])
		}

		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			gogopher(w, gopher)
			return
		}
		if age >= 1 && age < 4 {
			locale, err := getUserLocale(t)
			if err != nil {
				log.Printf("IP: %v, Err: %v\n", ip, err)
				gogopher(w, gopher)
				return
			}
			log.Printf("IP: %v, Name: %v, Email: %v, Locale %v\n", ip,
				user.Name,
				user.Email,
				locale,
			)
		} else {
			if p := parsePraise(r.FormValue("praise")); p == "" {
				gogopher(w, gopher)
				return
			}
		}

		langs, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
		if err != nil {
			log.Printf("IP: %v, Err: %v\n", ip, err)
			gogopher(w, gopher)
			return
		}

		if len(langs) == 0 {
			gogopher(w, gopher)
			return
		}

		for _, l := range langs {
			prefix := fmt.Sprintf("%v_gopher", l)
			s, _ := solver.New(solver.SolverConfig{
				Prefix:     prefix,
				Difficulty: 16,
			})
			s.Solve()
		}

		if v, ok := sp.Get("chosen"); ok {
			if v == chosen && v != "" && user.Email != "" {
				fmt.Fprintf(w, "%s\n", os.Getenv("FLAG"))
				log.Printf("End: IP: %v, Name: %v, Email: %v, Name2: %v, Age: %v, Praise: %v, Prove: %v, Language: %v\n",
					ip,
					user.Name,
					user.Email,
					r.FormValue("name"),
					r.FormValue("age"),
					r.FormValue("praise"),
					r.FormValue("prove"),
					r.Header.Get("Accept-Language"),
				)
				sp.Delete(emailKey)
				return
			}
		}
		gogopher(w, gopher)
	}
}

func parseName(name string,
	c *redis.Client,
	sp *store.Store,
) (int64, int64) {
	nameRunes := []rune(name)
	var (
		a int64
		b int64
	)
	for i := 0; i < len(nameRunes); i++ {
		for j, v := range config.Scripts {
			if unicode.In(nameRunes[i], v) {
				a, _ = c.Incr(config.Redis.UserKey + j).Result()
				b, _ = c.Incr(config.Redis.UserCountKey + j).Result()
			}
		}
	}
	sp.Set(config.Redis.UserKey, string(a))
	sp.Set(config.Redis.UserCountKey+name, string(b))
	return a, b
}

func parsePraise(praise string) string {
	var b [1 << 10]byte
	copy(b[:], praise)
	return base64.StdEncoding.EncodeToString(b[:])
}
