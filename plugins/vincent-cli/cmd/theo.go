package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var quotes = []string{
	// Original quotes
	"The heart of man is very much like the sea, it has its storms, it has its tides and in its depths it has its pearls too.",
	"I also have nature and art and poetry, and if that isn't enough, what is?",
	"Find things beautiful as much as you can, most people find too little beautiful.",
	"Success is sometimes the outcome of a whole string of failures.",
	"Just as we take the train to go to Tarascon or Rouen, we take death to go to a star.",
	"I want to make drawings that touch some people... I want to get to the point where people say of my work: that man feels deeply, that man feels keenly.",
	"Painters understand nature and love it, and teach us to see.",
	"It often seems to me that the night is much more alive and richly colored than the day.",
	"I don't know if you'll understand that one can speak poetry just by arranging colours well, just as one can say comforting things in music.",
	"It is already late, but I felt like writing to you again anyway. You are not here - but I need you and sometimes feel that we are not far away from each other.",
	"If I was without your friendship I would be sent back without remorse to suicide, and however cowardly I am, I would end up going...",

	// New quotes
	"What is done in love is done well.",
	"I put my heart and my soul into my work, and have lost my mind in the process.",
	"There is nothing more truly artistic than to love people.",
	"If you hear a voice within you say 'you cannot paint,' then by all means paint, and that voice will be silenced.",
	"I wish they would take me as I am.",
	"Be clearly aware of the stars and infinity on high. Then life seems almost enchanted after all.",
	"The fishermen know that the sea is dangerous and the storm terrible, but they have never found these dangers sufficient reason for remaining ashore.",
	"It is good to love many things, for therein lies the true strength, and whosoever loves much performs much, and can accomplish much, and what is done in love is well done.",
	"I would rather die of passion than of boredom.",
	"Normality is a paved road: it’s comfortable to walk, but no flowers grow on it.",
	"I don't know anything with certainty, but seeing the stars makes me dream.",
}

var theoCmd = &cobra.Command{
	Use:   "theo",
	Short: "Commands related to Theo van Gogh.",
}

var lovedMeCmd = &cobra.Command{
	Use:   "loved-me",
	Short: "A reminder of the bond between brothers.",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		quote := quotes[rand.Intn(len(quotes))]
		fmt.Println(quote)
	},
}

func init() {
	rootCmd.AddCommand(theoCmd)
	theoCmd.AddCommand(lovedMeCmd)
}
