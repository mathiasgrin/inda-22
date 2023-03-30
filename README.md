# **Övning palinda-2**
Mathias Grindsäter (grin@kth.se)

### **Idag**
* Gå igenom veckans task + presentation
* Kort om Structs
* Kort om obuffrade och buffrade kanaler + uppgift
* Bygg en Thread Pool för att hantera DB-requests.
* Övningsuppgifter, om vi hinner.

#### 💬**Veckans task**

Diskutera först i grupper av två eller tre i 15 minuter:
* Gå igenom uppgifterna.
* Vad gjorde ni lika/olika, vad var svårt?
* Bestäm själva i gruppen hur ni vill dela upp vem som presenterar vad.
* Gruppresentation.

### **Structs**
> Väldigt lika klasser och objekt i Java.

##### **Skapa en Struct**
<details>
<summary>Exempel</summary>
<br>

```go
type car struct {
	company string
	model string
	yearMade int
}

type driver struct {
	name string
	car car
}
```
</details>
<br>

##### **Skapa en Struct-variabel**

<details>
<summary>Exempel</summary>
<br>

```go
myCar := car{"Volvo", "XC90", 2006}
johnTheDriver := driver{"John Johnson", myCar}
```

</details>
<br>

##### **Hämta variabler från en struct**
<details>
<summary>Exempel</summary>
<br>

```go
carCompany := myCar.company
modelOfCarOfJohn := johnTheDriver.car.company
```
</details>
<br>




### **Unbuffered vs Buffered Channels**
#### **Send**
> En send till en obuffrad kanal **blockeras** tills en receiver
> är redo att ta emot från kanalen.

<details>
<summary>Exempel</summary>
<br>

```go
func main() {
	unbufferedCh := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		received := <-unbufferedCh
	}()

	// main routine blockerad i 3 sek.
	unbufferedCh <- 1
	// mer kod ...
```
</details>
<br>


> Vi kan skicka till en buffrad kanal tills den blir full, även om
> ingen receiver står redo att ta emot.

<details>
<summary>Exempel</summary>
<br>

```go
func main() {
    bufferedCh := make(chan int, 2)

    go func() {
        time.Sleep(5 * time.Second)
        for v := range bufferedCh {
            // Gör något med v
        }
	}()

    // Kan skicka trots att receiver inte är redo
    bufferedCh <- 1
    // Kan skicka trots att receiver inte är redo
    bufferedCh <- 2
    // Blockeras ca 5 sek eftersom kanalen är full
    bufferedCh <- 3
    // mer kod ...
}
```
</details>
<br>

####**Receive**

> Receive från en kanal blockerar tills något skickas 
> till kanalen. Exemplet nedan gäller även om kanalen
> är buffrad.


<details>
<summary>Exempel</summary>
<br>

```go
func main() {
    ch := make(chan int)
		
    go func() {
         time.Sleep(5 * time.Second)
        ch <- 1
	}()

    // Blockerar i 5 sek
    <- ch
}
```
</details>
<br>

### **Buffrade kanaler - uppgift**
Koden nedan ger sju utskriftsrader. Analysera koden och ange vad som skrivs ut på varje rad.

<details>
<summary>Uppgift</summary>
<br>

```go
func main() {
    ch := make(chan int, 2)
    iterations := 3

    go func() { // Run anonymous function as a goroutine.
        for i := 1; i <= iterations; i++ {
            ch <- i
            fmt.Printf("Sent %d to the channel.\n", i)
        }
        fmt.Printf("All %d numbers sent!\n", iterations)
        close(ch) // Close the channel when done sending.
	}()

    time.Sleep(3 * time.Second) // Give the goroutine time to run.

    for chVal := range ch {
        fmt.Printf("%d received!\n", chVal)
        time.Sleep(200 * time.Millisecond)
    }
}
```

</details>
<br>

<details>
<summary>Lösning</summary>
<br>

>Sent 1 to the channel.\
>Sent 2 to the channel.\
>1 received!\
>Sent 3 to the channel.\
>All 3 numbers sent!\
>2 received!\
>3 received!
</details>
<br>

Fundera även på vad som händer om vi
tar bort `close(ch)`. Tänk först ut ett svar och testkör sedan med `close(ch)` 
utkommenterad. Varför blir det som det blir?
 

<details>
<summary>Förklaring</summary>
<br>

>Om vi inte stänger kanalen
>så kommer `for chVal := range ch` att fortsätta vänta
>på att hämta data från kanalen. Dock finns ingen
>goroutine som längre skriver data till kanalen 
>(vår anonyma sender-goroutine har ju kört färdigt efter sin
>3:e iteration). Vi får således Deadlock.
</details>
<br>

## **Thread Pool**
> Vi ska nu implementera en Thread Pool / Worker Pool steg för steg.\
> Programmet simulerar en ström av inkommande DB-requests i form av SQL-queries.
> De ser ut såhär:\
> `"SELECT name FROM user WHERE user.id=abc123"`\
> Ett sådant SQL query hämtar namnet på en användare (user) kopplad till det givna id:et.
> 
> Kodskelettet hittar du i filen `threadPoolTask.go` i github-repot.

####**STEG 1**
>Skapa 3 struct.\
> Struct 1: Heter user och har variablerna name och id (båda strings).\
> struct 2: Heter request och har variablerna sqlQuery och id. (båda strings).\
> struct 3: Heter result och har variabeln request av typen request och name av typen string.
> 

<details>
<summary>Kontroll</summary>
<br>

```go
// ----------STRUCTS----------//
type user struct {
    name string
    id string
}

type request struct {
    sqlQuery string // Example: "SELECT name FROM user WHERE user.id="
    id string // Example ID: O533TUJgPb
}

type result struct {
    request request
    name string
}
```
</details>
<br>

####**STEG 2**

>Skapa en array (obs inte slice) av storlek `numOfDBUsers` som heter users och tar
> element av typer user.

<details>
<summary>Kontroll</summary>
<br>

```go
// ----------ARRAYS----------//
// The array represents the users held in the DB
var users [numOfDbUsers]user
```
</details>
<br>

####**STEG 3**
>Skapa två buffrade kanaler.\
> Kanal 1: Heter requestsCh och tar data av typen request. Kapacitet ska vara 10 element.\
> Kanal 2: Heter resultsCh och tar data av typen result. Kapacitet ska vara 10 element.\
> Notera att kanalerna är globala, vilket betyder att de kan nås av alla funktioner utan att behöva tas in som argument.

<details>
<summary>Kontroll</summary>
<br>

```go
// ----------CHANNELS----------//
var requestsCh = make(chan request, 10)
var resultsCh = make(chan result, 10)

```
</details>
<br>

####**STEG 4**
> Implementera funktionen `getUserFromDB(id string) user`.\
> Funktionen får in ett unikt id och ska se om id:et
> finns i DB:en, det vill säga i array:en `user`.\
> Om ID:et hittas i DB:en så ska funktionen returna user:n med detta id.\
> Om ingen user med det givna id:et hittas ska en tom user returnas.

<details>
<summary>Kontroll</summary>
<br>

```go
/*
* The DB takes a unique ID as argument and returns a user.
* This is a simulated task made by the DB.
 */
func getUserFromDB(id string) user {
    time.Sleep(100 * time.Millisecond) // Simulate DB request time.
    for _, user := range users {
        if id == user.id {
            return user
        }
    }
    return user{} // Return an empty user if we failed to find the id in the DB.
}
```
</details>
<br>

####**STEG 5**
> Fyll i den rad som saknas i funktionen `requestFactory`.\
> Funktionen skapar requests och skickar dessa till requestsCh.

<details>
<summary>Kontroll</summary>
<br>

```go
/*
* Creates SQL DB requests and
* sends to the requestsCh
*/
func requestFactory(numOfRequests int) {
    for i := 0; i < numOfRequests; i++ {
        sqlQuery := "SELECT name FROM user WHERE user.id=" // SQL query to get a name.
        id := getRandomIdFromUsers()                       // Generate a random ID.
        req := request{sqlQuery, id}                       // Create a request.
        requestsCh <- req                                  // Send the request to the requestsCh.
    }
    close(requestsCh) // Close the channel when numOfRequests requests have been created.
}
```
</details>
<br>

####**STEG 6**
> Skriv färdigt funktionen `taskExecutor`.\
> Funktionen representerar arbetet som varje individuell Goroutine/Worker/Tråd utför.\
> Funktionen tar emot requests från requestsCh och hämtar den user som är kopplad\
> till request-id:et genom att anropa `getUserFromDB` som vi implementerade tidigare.
> Se till att funktionen kan skicka ett `result` till `resultsCh`

<details>
<summary>Kontroll</summary>
<br>

```go
func taskExecutor(wg *sync.WaitGroup) {
    for request := range requestsCh {
        user := getUserFromDB(request.id)
        name := user.name
        res := result{request, name}
        resultsCh <- res
    }
    wg.Done()
}
```
</details>
<br>

####**STEG 7**
> Funktionen `resultReceiver` är det sista steget i vår Thread Pool.\
> Funktionen tar emot resultaten och skriver ut 3 variabler.
> Fyll i det som saknas.

<details>
<summary>Kontroll</summary>
<br>

```go
func resultReceiver(done chan<- bool) {
    for result := range resultsCh {
        query := result.request.sqlQuery
        name := result.name
        id := result.request.id
        fmt.Printf("Query: %s%s ==> %s\n", query, id, name)
    }
    done <- true
}
```
</details>
<br>

####**STEG 8**
> Slutligen behöver vi en funktion som startar våra workers/Goroutiner/trådar.\
> Detta görs i funktionen `createThreadPool`
> Funktionen ska anropa funktionen `taskExecutor` som n=numOfWorkers antal Goroutiner.
> Fyll i det som saknas.

<details>
<summary>Kontroll</summary>
<br>

```go
func createThreadPool(numOfWorkers int) {
    var wg sync.WaitGroup
    for i := 0; i < numOfWorkers; i++ {
        wg.Add(1)
        go taskExecutor(&wg)
    }
    wg.Wait()
    close(resultsCh)
}
```
</details>
<br>

> Nu är vår Thread Pool färdig. Bra jobbat! Testa att köra programmet\
> med olika antal workers/goroutiner och att
> programmet faktiskt hanterar DB-requests snabbare!

