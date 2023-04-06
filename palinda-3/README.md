# **Övning palinda-3**
Mathias Grindsäter (grin@kth.se)

### **Idag**
* Gå igenom veckans task + presentation
* Kort om Mutex + uppgift
* Git

#### 💬**Veckans task**

Diskutera först i grupper av två eller tre i 20 minuter:
* Gå igenom uppgifterna.
* Vad gjorde ni lika/olika, vad var svårt?
* Bestäm själva i gruppen hur ni vill dela upp vem som presenterar vad.
* Presentera som grupp.

### **Mutex**
> Låser en variabel så att olika trådar får tillgång en i taget.

##### **Bankkonto**
> Öppna filen `mutexExerciseSkal.go` och fyll i det som saknas.\
> Använd inget Mutex-objekt (än). Tanken är att minska
> saldot från 1000 till 0 med 1000 gorutiner.


<details>
<summary>Kontroll</summary>
<br>

```go
func incrementBalance(s *bankAccount, wg *sync.WaitGroup) {
    (*s).balance = (*s).balance - 1
    wg.Done()
}

/*
 * For each Goroutine we want to decrement the balance
 * by 1.
 */
func main() {
    numOfGoroutines := 1000
    myAccount := bankAccount{"Handelsbanken", 1000}
    var w sync.WaitGroup
    for i := 0; i < numOfGoroutines; i++ {
        w.Add(1)
        go incrementBalance(&myAccount, &w)
    }
    w.Wait()
    fmt.Println(myAccount.balance)
}
```
</details>
<br>

>Testa att köra programmet några gånger (OBS! Ej i Go Playground). Vad händer och varför?

<details>
<summary>Förklaring</summary>
<br>

>Vi får olika slutsaldon vid varje körning. Detta eftersom
> gorutinerna inte väntar på varandra.

</details>
<br>

> Låt oss nu använda en Mutex.
> Vi deklarerar en Mutex likt en WorkGroup:
>```go
>var mtx sync.Mutex
>```
>Vi är främst intresserade av två funktioner, vilka låser och låser upp
> en variabel.
> ```go
> mtx.Lock()
> // Mellan har vi det vi vill låsa.
> mtx.Unlock()
>```
>Implementera nu programmet med hjälp av en Mutex.

<details>
<summary>Facit</summary>
<br>

```go
func incrementBalance(s *bankAccount, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	(*s).balance = (*s).balance - 1
	mtx.Unlock()
	wg.Done()
}

/*
* For each Goroutine we want to decrement the balance
* by 1.
 */
func main() {
	numOfGoroutines := 1000
	myAccount := bankAccount{"Handelsbanken", 1000}
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < numOfGoroutines; i++ {
		w.Add(1)
		go incrementBalance(&myAccount, &w, &m)
	}
	w.Wait()
	fmt.Println(myAccount.balance)
}

```
</details>
<br>

>Kör programmet några gånger och se att det verkligen funkar.
>Kontemplera sedan i par: Vad tror ni är mest effektivt, denna lösning eller 
>en sekventiell implementation?

<details>
<summary>OBS! En elegant lösning med en kanal</summary>
<br>

```go
func incrementBalance(s *bankAccount, wg *sync.WaitGroup, blockingCh chan bool) {
	blockingCh <- true
	(*s).balance = (*s).balance - 1
	<- blockingCh
	wg.Done()
}

func main() {
	numOfGoroutines := 1000
	myAccount := bankAccount{"Handelsbanken", 1000}
	var w sync.WaitGroup
	blockingCh := make(chan bool, 1)
	for i := 0; i < numOfGoroutines; i++ {
		w.Add(1)
		go incrementBalance(&myAccount, &w, blockingCh)
	}
	w.Wait()
	fmt.Println(myAccount.balance)

}
```
</details>
<br>

### **Git**
[Git Bootcamp](https://github.com/eeegl/inda22-uppgifter/tree/main/palinda-3)



