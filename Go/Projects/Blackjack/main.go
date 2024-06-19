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
		// Convert to decimal rep. for other values (2-10)
		value = fmt.Sprintf("%d", card.value)
	}

	return value + suit
}

type Deck struct {
	cards []Card
}

func (d *Deck) dealCard() Card {
	index := rand.Intn(len(d.cards))
	card := d.cards[index]
	d.cards = append(d.cards[:index], d.cards[index+1:]...)
	return card
}

func (d *Deck) create() {
	for s := 0; s < 4; s++ {
		for c := 1; c <= 13; c++ {
			d.cards = append(d.cards, Card{value: c, suit: s})
		}
	}
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d Deck) getCards(playerCards []Card) string {
	cards := ""
	for _, card := range playerCards {
		cards += card.getString() + " "
	}
	return cards
}

type Game struct {
	deck        Deck
	playerCards []Card
	dealerCards []Card
}

func (game *Game) dealStartingCards() {
	game.playerCards = append(game.playerCards, game.deck.dealCard())
	game.playerCards = append(game.playerCards, game.deck.dealCard())
	game.dealerCards = append(game.dealerCards, game.deck.dealCard())
	game.dealerCards = append(game.dealerCards, game.deck.dealCard())
}

func (d Deck) getHandCount(playerCards []Card) int {
	count := 0
	aces := 0

	for _, card := range playerCards {
		if card.value == 1 {
			aces++
			count++
		} else if card.value > 10 {
			count += 10
		} else {
			count += card.value
		}
	}

	for aces > 0 && count+10 <= 21 {
		count += 10
		aces--
	}

	return count
}

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
		return 2, "Dealer got Blackjack!"
	}

	return 0, ""
}

func (game *Game) play(bet float64) float64 {
	game.deck.create()
	game.deck.shuffle()
	game.dealStartingCards()

	fmt.Println("Dealing the starting cards:")
	fmt.Println("---------------------------")
	fmt.Println("Player's Cards: ", game.deck.getCards(game.playerCards))
	fmt.Println("Player's Count: ", game.deck.getHandCount(game.playerCards))
	fmt.Println("Dealer's Cards: ", game.deck.getCards(game.dealerCards[:1]), "Hidden")
	fmt.Println("---------------------------")

	winner, reason := game.checkGame()
	if winner != 0 {
		fmt.Println(reason)
		if winner == 1 {
			return bet
		} else {
			return -bet
		}
	}

	for {
		game.playerTurn()
		winner, reason = game.checkGame()
		if winner != 0 {
			fmt.Println(reason)
			if winner == 1 {
				return bet
			} else {
				return -bet
			}
		}

		game.dealerTurn()
		winner, reason = game.checkGame()
		if winner != 0 {
			fmt.Println(reason)
			if winner == 1 {
				return bet
			} else {
				return -bet
			}
		}

		playerCount := game.deck.getHandCount(game.playerCards)
		dealerCount := game.deck.getHandCount(game.dealerCards)
		fmt.Println("Player's Cards: ", game.deck.getCards(game.playerCards))
		fmt.Println("Player's Count: ", playerCount)
		fmt.Println("Dealer's Cards: ", game.deck.getCards(game.dealerCards))
		fmt.Println("Dealer's Count: ", dealerCount)

		if playerCount > dealerCount {
			fmt.Println("Player wins!")
			return bet
		} else {
			fmt.Println("Dealer wins!")
			return -bet
		}
	}
}

func (game *Game) playerTurn() {
	for {
		fmt.Println("Your turn: (h)it or (s)tand?")
		action := enterString()
		if action == "h" {
			game.playerCards = append(game.playerCards, game.deck.dealCard())
			fmt.Println("Player's Cards: ", game.deck.getCards(game.playerCards))
			fmt.Println("Player's Count: ", game.deck.getHandCount(game.playerCards))
			if game.deck.getHandCount(game.playerCards) > 21 {
				break
			}
		} else if action == "s" {
			break
		} else {
			fmt.Println("Invalid action. Please enter 'h' or 's'.")
		}
	}
}

func (game *Game) dealerTurn() {
	for game.deck.getHandCount(game.dealerCards) < 17 {
		game.dealerCards = append(game.dealerCards, game.deck.dealCard())
	}
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
