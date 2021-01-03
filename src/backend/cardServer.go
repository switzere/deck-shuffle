package main

import (
  "fmt"
  "net/http"
  //"io/ioutil"
  "github.com/gorilla/mux"
  "math/rand"
  "time"
  "strconv"
  "encoding/json"
)

type Deck struct {
  cards []Card
  drawnCards []Card
  inDeckCards []Card
}

type Card struct {
  suit int
  value int
}

type ResponseCard struct {
  suit string
  value string
}


const (
  Spade = iota
  Heart
  Club
  Diamond
)

var deck Deck

func draw(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")

  card := drawCard(&deck)
  if card.suit == -1 && card.value == -1 {
    deck = *makeDeck("euchre")
    card = drawCard(&deck)
  }

  //fmt.Fprintf(w, "%s", numToFace(card.value))
  //fmt.Fprintf(w, "%s", numToSuit(card.suit))
  //responseCard := ResponseCard{suit: numToSuit(card.suit), value: numToFace(card.value)}
  responseCard := map[string]string{"suit": numToSuit(card.suit), "value": numToFace(card.value)}

  //str := `{"suit": `+numToSuit(card.suit)+`, "value": `+numToFace(card.value)+`}`
  jsonResp, _ := json.Marshal(responseCard)

  fmt.Printf("%s\n",jsonResp)
  fmt.Fprintf(w, "%s", jsonResp)

  //json.NewEncoder(w).Encode(responseCard)
}


func main() {
  router := mux.NewRouter()

  deck = *makeDeck("euchre")

  //router.HandleFunc("/", home)
  router.HandleFunc("/draw", draw).Methods("GET", "OPTIONS")
  http.ListenAndServe(":3001", router)

}




func drawCard(deck *Deck) Card {
  if(len(deck.inDeckCards) == 0) {
    return Card{-1, -1}
  }
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  r := random.Intn(len(deck.inDeckCards))
  retCard := deck.inDeckCards[r]
  deck.inDeckCards = append(deck.inDeckCards[:r], deck.inDeckCards[r+1:]...)
  return retCard
}


func makeDeck(deckType string) *Deck {
  var deck = new(Deck)

  if deckType == "euchre" {

    for i := 0; i < 4; i++ {
      for j := 0; j < 6; j++ {

        newCard := Card{i, j + 9}
        deck.cards = append(deck.cards, newCard)
        deck.inDeckCards = append(deck.inDeckCards, newCard)

      }
    }
  } else if deckType != "euchre" {

    for i := 0; i < 4; i++ {
      for j := 0; j < 13; j++ {

        newCard := Card{i, j}
        deck.cards = append(deck.cards, newCard)
        deck.inDeckCards = append(deck.inDeckCards, newCard)

      }
    }
  }

  return deck

}


func numToSuit(suit int) string {

  if suit == 0 {
    return "Spade"
  } else if suit == 1 {
    return "Heart"
  } else if suit == 2 {
    return "Club"
  } else if suit == 3 {
    return "Diamond"
  }

  return "error"
}

func numToFace(value int) string {

  if value == 11 {
    return "Jack"
  } else if value == 12 {
    return "Queen"
  } else if value == 13 {
    return "King"
  } else if value == 14 {
    return "Ace"
  }

  return  strconv.Itoa(value)
}
