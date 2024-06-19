package main

import (
	"bufio" // provides buffered I/O operations (NewReader and ReadString)
	"fmt"
	"math/rand"
	"os"      // used to create a new buffered reader that reads - (os.Stdin)
	"strconv" // Used to ParseFloat
	"strings" // Used for string manipulation - TrimSuffix
)

type Card struct {
	value int
	suit  int // 0 - spades, 1 - hearts, 2 - diamonds, 3 - clubs
}

func (card Card) getString() string {
	var suit string
	var value string

	switch card.suit {
	case 0:
		suit = "♠"
	case 1:
		suit = "♥"
	case 2:
		suit = "♦"
	case 3:
		suit = "♣"
	}

	switch card.value {
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	case 1:
		value = "A"
	default:
		//Convert to decimal rep. for other values (2-10)
		value = fmt.Sprintf("%d", card.value)
	}

	return value + suit
}

type Deck struct {
	cards []Card
}

func (d *Deck) dealCard() Card {
	// Get card
	index := uint(rand.Intn(len(d.cards)))
	card := d.cards[index]
	// Remove the card
	d.cards = append(d.cards[:index], d.cards[index+1:]...)
	return card
}

func (d *Deck) create() {
	// Create the new card
	var card Card

	for s := 0; s < 4; s++ {
		card.suit = s
		for c := 1; c <= 13; c++ {
			card.value = c
			d.cards = append(d.cards, card)
		}
	}
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d Deck) getCards(playersCards []Card) (cards string) {
	for c := range playersCards {
		cards += playersCards[c].getString() + " "
	}
	return
}

type Game struct {
	deck        Deck
	playerCards []Card
	dealerCards []Card
}

func (game *Game) dealStartingCards() {
	// Deal 2 cards to each (player 1st!)
	// Adds card to players hand & Removes it from deck
	game.playerCards = append(game.playerCards, game.deck.dealCard())
	game.playerCards = append(game.playerCards, game.deck.dealCard())

	// Do the same for the dealer
	game.dealerCards = append(game.dealerCards, game.deck.dealCard())
	game.dealerCards = append(game.dealerCards, game.deck.dealCard())
}

func (d Deck) getHandCount(playerCards []Card) (count int) {
	for c := range playerCards {
		if playerCards[c].value > 10 {
			count += 10
		} else {
			count += playerCards[c].value
		}
	}
	return
}

// Winner will be set as 0-none, 1-Player, 2-Dealer
func (game *Game) checkGame() (int, string) {
	playerCount := game.deck.getHandCount(game.playerCards)
	dealerCount := game.deck.getHandCount(game.dealerCards)
	if playerCount > 21 {
		return 2, "Player Busted."
	} else if dealerCount > 21 {
		return 1, "Dealer Busted."
	} else if playerCount == 21 {
		return 1, "Player got Blackjack!"
	} else if dealerCount == 21 {
		return 1, "Dealer got Blackjack!"
	}

	return 0, ""
}

func (game *Game) play(bet float64) (pot float64) {
	winner := false
	pot = bet

	// Make & Shuffle Deck
	game.deck.create()
	game.deck.shuffle()

	// Deal the starting cards
	game.dealStartingCards()

	playerCards := game.deck.getCards(game.playerCards)
	playerCount := game.deck.getHandCount(game.playerCards)
	dealerCards := game.deck.getCards(game.dealerCards)
	dealerCount := game.deck.getHandCount(game.dealerCards)

	fmt.Println("Dealing the starting cards:")
	fmt.Println("---------------------------")
	fmt.Println("Players Cards: ", playerCards)
	fmt.Println("Players Count: ", playerCount)
	fmt.Println("---------------------------")

	winner = game.checkGame()
	for winner != true {

		winner = game.checkGame()
	}

	fmt.Println("Dealers Cards: ", dealerCards)
	fmt.Println("Dealers count: ", dealerCount)

	fmt.Println("Current pot: ", pot)

	return
}

// func (game *Game) playerTurn() bool {

// }

// func (game *Game) dealerTurn() {

// }

func (d Deck) print() {
	var deck string

	for c := 0; c < len(d.cards); c++ {
		deck += d.cards[c].getString() + " "
	}
	fmt.Println(deck)
}

func enterString() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return ""
	}

	// remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	balance := float64(100)

	for balance > 0 {
		fmt.Printf("Your balance is: $%.2f\n", balance)
		fmt.Printf("Enter your bet (q to quit): ")
		bet, err := strconv.ParseFloat(enterString(), 64)
		if err != nil {
			break
		}
		if bet > balance || bet <= 0 {
			fmt.Println("Invalid bet.")
			continue
		}

		game := Game{}
		balance += game.play(bet)
	}

	fmt.Printf("You left with: $%2.f\n", balance)
}
