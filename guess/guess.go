package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Guess game
func Guess() {

	fmt.Println()

	time := time.Now().Unix()
	rand.Seed(time)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it ?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		yourNum, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			log.Fatal(err)
		}

		if target > yourNum {
			fmt.Println("Oops your guess was low")
		} else if target < yourNum {
			fmt.Println("Oops your guess was high")
		} else {
			fmt.Println("Good job! You guessed it!", target)
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you did't guess my number. It was", target)
	}

}
