package store

import (
	"time"

	"balsnctf/gopherparty/event"
)

type Store struct {
	data map[string]string
}

func (s *Store) Get(name string) (string, bool) {
	v, ok := s.data[name]
	return v, ok
}

func (s *Store) Set(name string, value string) {
	s.data[name] = value
}

func (s *Store) Delete(name string) {
	if _, ok := s.data[name]; ok {
		delete(s.data, name)
	}
}

func New() *Store {
	s := &Store{
		data: make(map[string]string),
	}
	// 孟子曰：「舜發於畎畝之中，傅說舉於版築之閒，膠鬲舉於魚鹽之中，
	// 管夷吾舉於士，孫叔敖舉於海，百里奚舉於市。故天將降大任於斯人也，
	// 必先苦其心志，勞其筋骨，餓其體膚，空乏其身，行拂亂其所為，
	// 所以動心忍性，曾益其所不能。人恒過，然後能改；困於心，衡於慮，而後作；
	// 徵於色，發於聲，而後喻。入則無法家拂士，出則無敵國外患者，國恒亡。
	// 然後知生於憂患而死於安樂也。」
	// "When Heaven is about to place a great responsibility on a man,
	// it always first frustrates his spirit and will, exhausts his
	// muscles and bones, exposes him to starvation and poverty, harasses
	// him by troubles and setbacks so as to stimulate his mind,
	// toughen his nature and enhance his abilities.",said Mencius.

	go func() {
		for {
			if event.IsDown() {
				for k := range s.data {
					if k != "chosen" {
						delete(s.data, k)
					}
				}
			}
			s.data["chosen"] = "WHO_IS_THE_CHOSEN_ONE"
			time.Sleep(2 * time.Millisecond)
		}
	}()

	go func() {
		for {
			for k := range s.data {
				if k != "chosen" {
					delete(s.data, k)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return s
}
