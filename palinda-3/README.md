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

##### **Del 1: Bankkonto**
> Öppna filen `mutexExerciseSkal.go` och fyll i det som saknas.\
> Använd inget Mutex-objekt (än). Tanken är att minska
> saldot från 1000 till 0 med 1000 gorutiner.
> Lös uppgiften själv innan du trycker på `kontroll`.


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


##### **Del 2: Lösning med Mutex**
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

