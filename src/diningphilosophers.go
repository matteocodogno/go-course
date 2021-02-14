package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
Implement the dining philosopher’s problem with the following constraints/modifications.

    - There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    - Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
    - The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
    - In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    - The host allows no more than 2 philosophers to eat concurrently.
    - Each philosopher is numbered, 1 through 5.
    - When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a
      line by itself, where <number> is the number of the philosopher.
    - When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a
      line by itself, where <number> is the number of the philosopher.
 */
type ChopS struct { sync.Mutex }

type Philo struct {
	id int
	leftCS, rightCS *ChopS
	pc chan bool
	nMeals int
}

func (p Philo) eat() {
	for p.nMeals < 3 {
		p.leftCS.Lock()
		p.rightCS.Lock()

		p.pc <- true
		canIEat := <- p.pc
		if canIEat {
			fmt.Printf("start to eat <%d>\n", p.id)
			p.nMeals ++
			fmt.Printf("finishing eating <%d>\n", p.id)

			p.pc <- false
		}

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

	close(p.pc)
}

func host(philos []*Philo, wg *sync.WaitGroup) {
	concurrentEat := 0
	var mut sync.Mutex
	var wg2 sync.WaitGroup
	wg2.Add(5)
	for i := 0; i < 5; i ++ {
		go philos[i].eat()
		go func(philo *Philo) {
			for readToEat := range philo.pc {
				mut.Lock()
				if readToEat {
					canYouEat := false
					if concurrentEat <= 2 {
						concurrentEat ++
						canYouEat = true
					}
					philo.pc <- canYouEat
				} else {
					concurrentEat --
				}
				mut.Unlock()
			}
			wg2.Done()
		}(philos[i])
	}

	wg2.Wait()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i ++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i ++ {
		if rand.Intn(1) == 0 {
			philos[i] = &Philo{i+1, CSticks[i], CSticks[(i+1)%5], make(chan bool), 0}
		} else {
			philos[i] = &Philo{i+1, CSticks[(i+1)%5], CSticks[i], make(chan bool), 0}
		}
	}

	go host(philos, &wg)
	wg.Wait()
}
