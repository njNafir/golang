package greetings

import (
	"fmt"
	"errors"

	"math/rand"
	"time"
)

// Hellos returns a map that associates each of the named people
// with a greeting message.

func Hellos(names []string) (map[string]string, error) {

	// A map to associate names with messages.

	fmt.Println(names)

	messages := make(map[string]string)

	fmt.Println(messages)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with 
		// the name.
		
		messages[name] = message
	}

	fmt.Println(messages)

	return messages, nil

}

// Hello returns a greeting for a name

func Hello(name string) (string, error) {
	// Hello is getting a name parameter, which can be attached with greeting message

	// Check if name is empty, then return error if name is empty
	if name == "" {
		return "", errors.New("name empty")
	}

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}


// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message format

	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v",
		"Hail, %v Well met",
	}

	fmt.Println(formats)

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	
	return formats[rand.Intn(len(formats))]
}