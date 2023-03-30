/*
Author: Mathias Grindsäter (grin@kth.se) - 2023-03-30
THREAD POOL DB SIMULATION for DD1396
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// ----------CONSTANTS----------//
const numOfDbUsers = 45

// ----------STRUCTS----------//
// TODO struct1
// TODO struct2
// TODO struct3

// ----------ARRAYS----------//
// The array represents the users held in the DB
// TODO

// ----------CHANNELS----------//
// TODO

// ----------SIMULATION FUNCTIONS----------//

/*
* The DB takes a unique ID as argument and returns a user
* from the users array.
* This is a simulated task made by the DB.
 */
func getUserFromDB(id string) user {
	time.Sleep(100 * time.Millisecond) // Simulate DB request time.
	// TODO
}

/*
* Creates SQL DB requests and
* sends to the requestsCh
 */
func requestFactory(numOfRequests int) {
	for i := 0; i < numOfRequests; i++ {
		sqlQuery := "SELECT name FROM user WHERE user.id=" // SQL query to get a name.
		id := getRandomIdFromUsers()                       // Generate a random ID.
		// TODO
		requestsCh <- req // Send the request to the requestsCh.
	}
	close(requestsCh) // Close the channel when numOfRequests requests have been created.
}

// ----------THREAD POOL FUNCTIONS----------//

/*
* Receives requests, fetches the user from the DB,
* creates a result and sends the result to the ResultsCh.
* This is the most central part of the Thread Pool.
* The function represents the work of each
* individual worker/goroutine, or rather
* can be seen as the threads themselves.
 */
func taskExecutor(wg *sync.WaitGroup) {
	for request := range requestsCh {
		user := getUserFromDB(request.id)
		// TODO
		resultsCh <- res
	}
	wg.Done()
}

/*
* Receives the results from the resultCh
* and prints out the SQL query with the ID and
* the corresponding user's name.
 */
func resultReceiver(done chan<- bool) {
	for result := range resultsCh {
		// TODO
		fmt.Printf("Query: %s%s ==> %s\n", query, id, name)
	}
	done <- true // All work is done. Tell the main routine to continue.
}

/*
* Initialize the Threads/Goroutines/Workers,
* thus putting the Thread Pool in action.
* Each Goroutine executes taskExecutor().
* NumOfWorkers is the amount of Goroutines
* we want.
 */
func createThreadPool(numOfWorkers int) {
	// TODO Create a WaitGroup here
	for i := 0; i < numOfWorkers; i++ {
		// TODO
	}
	wg.Wait()
	close(resultsCh)
}

func main() {
	numOfRequests := 200
	createUsers()

	//Read number of workers from terminal
	var inputNumOfWorkers string
	fmt.Print("Num of workers: ")
	fmt.Scanf("%s", &inputNumOfWorkers)
	NumOfWorkers, _ := strconv.Atoi(inputNumOfWorkers)

	// Start
	start := time.Now()                //Start the clock.
	go requestFactory(numOfRequests)   // Create tasks and start filling up the requestsCh.
	time.Sleep(200 * time.Millisecond) // Give some time to fill upp the requestsCh.

	done := make(chan bool)
	go resultReceiver(done)        // Start goroutine that constantly takes in the results from the results channel.
	createThreadPool(NumOfWorkers) // Get the workers started (start the Thread Pool).
	<-done                         // Block until receiveResults reports done.

	diff := time.Since(start)
	fmt.Println("Time taken: ", diff.Seconds())
}

// ----------HELPER FUNCTIONS----------//

func createUsers() {
	for i, name := range names {
		usr := user{name, idGenerator()}
		users[i] = usr
	}
}

func idGenerator() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	strLength := 10
	randStr := make([]byte, strLength)
	for j := range randStr {
		randStr[j] = chars[rand.Intn(len(chars))]
	}
	return string(randStr)
}

func getRandomIdFromUsers() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(users))
	randomIdInUsers := users[randomIndex].id
	return randomIdInUsers
}

var names = []string{"Lena Finnegan", "Joan Lange", "Octavia Ashworth",
	"Abel Guillen", "Uriel Bourne", "Omari Gooch",
	"Brady Catalano", "Devon Stern", "Jarret Landrum",
	"Keshon Breeden", "Lexie McCartney", "Miriam Trevino",
	"Daija Dickerson", "Marlee Odell", "Jordy Quiroz",
	"Curtis Loy", "Devan Hamblin", "Patricia Register",
	"Jakob Grabowski", "Tehya Qualls", "Lilian Conn",
	"Johnpaul Duarte", "Bayleigh Fogle", "Kerrigan Dasilva",
	"Jailyn Grogan", "Josh Halverson", "Dominick Lomeli",
	"Cori Campbell", "Kristina Bright", "Morris Wilson",
	"Sydni Yazzie", "Priscilla Briggs", "Francis Ingle",
	"Johnpaul Bower", "Brenda Brito", "Stephon Keys",
	"Hayden Fleming", "Aysia Rand", "Noor Pettit",
	"Jodi Levesque", "Salvatore Serna", "Darwin Marcum",
	"Jevon Rasmussen", "Colby Danner", "Marissa Stringer"}
