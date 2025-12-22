package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// var enemyStore = make([]*Enemy, 0)
// var playerStore = make([]*Player, 0)
// var enemyMap = make(map[string]*Enemy)
// var playerMap = make(map[string]*Player)
var c = counter()

var playerInterface = make([]Character, 0)
var enemyInterface = make([]Character, 0)

type Character interface {
	addPlayer() error
	Attack() int
	Hpreduced(int) int
	showname() string
}

type Player struct {
	Name string
	Hp   int
	Hit  int
}

type Enemy struct {
	Name string
	Hp   int
	Hit  int
}

func (p *Player) addPlayer() error {
	var playerName string
	fmt.Printf("Enter the name of the player: ")
	fmt.Scanf("%s", &playerName)
	p.Name = playerName
	p.Hp = 100

	playerInterface = append(playerInterface, p)
	// playerMap[p.Name] = p
	// playerStore = append(playerStore, p)
	return nil
}

func (e *Enemy) addPlayer() error {
	enemyName := fmt.Sprintf("Enemy%d", c())
	e.Name = enemyName
	e.Hp = 100
	enemyInterface = append(enemyInterface, e)
	// enemyMap[e.Name] = e
	// enemyStore = append(enemyStore, e)
	return nil
}

func (playerGotDamage *Player) Hpreduced(dmg int) int {
	playerGotDamage.Hp -= dmg
	return playerGotDamage.Hp
}

func (playername *Player) showname() string {
	return playername.Name
}

func (enemyname *Enemy) showname() string {
	return enemyname.Name
}

func (enemyGotDamage *Enemy) Hpreduced(dmg int) int {
	enemyGotDamage.Hp -= dmg
	return enemyGotDamage.Hp
}

var dmg int

func (playerAttacked *Player) Attack() int {

	fmt.Printf("\n%s's turn to attack:\nEnter the hit (10-20): ", playerAttacked.Name)
	fmt.Scanf("%d", &dmg)

	if dmg > 20 || dmg < 10 {
		for {
			fmt.Printf("Enter in valid range (10-20): ")
			fmt.Scanf("%d", &dmg)

			if dmg <= 20 && dmg >= 10 {
				break
			}
		}
	}

	var prob = r.Intn(8)
	if prob%2 == 0 {
		// fmt.Println("Damage Blocked! +")
		playerAttacked.Hit = dmg - prob
		fmt.Printf("Damage Blocked! Damage received: %d-%d = %d\n", dmg, dmg-playerAttacked.Hit, playerAttacked.Hit)
	} else {
		// fmt.Println("Extra Damage!")
		playerAttacked.Hit = dmg + prob
		fmt.Printf("Extra Damage! Damage received: %d+%d = %d\n", dmg, dmg+playerAttacked.Hit, playerAttacked.Hit)
	}

	return playerAttacked.Hit
}

func (enemyAttacked *Enemy) Attack() int {
	fmt.Printf("\n%s attacked:\n", enemyAttacked.Name)
	var eprob = r.Intn(10)

	min := 10
	max := 20
	dmg = rand.Intn(max-min+1) + min

	if eprob%2 == 0 {
		// fmt.Println("Damage Blocked!")
		enemyAttacked.Hit = dmg - eprob
		fmt.Printf("Damage Blocked! Damage received: %d-%d = %d\n", dmg, dmg-enemyAttacked.Hit, enemyAttacked.Hit)
	} else {
		enemyAttacked.Hit = dmg + eprob
		fmt.Printf("Extra Damage! Damage received: %d+%d = %d\n", dmg, dmg+enemyAttacked.Hit, enemyAttacked.Hit)
	}
	// enemyAttacked.Hit = 10 + eprob
	return enemyAttacked.Hit
}

func addplayer(character Character) error {
	return character.addPlayer()
}

func hitDamage(character Character, reducedHp int) int {
	return character.Hpreduced(reducedHp)
}

func attackOpp(character Character) int {
	return character.Attack()
}

func namedisplay(character Character) string {
	return character.showname()
}

func main() {

	var n int
	// var dmg int
	var win = 2
	// fmt.Println(win)
	fmt.Printf("Enter number of players: ")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		p := &Player{}
		p.addPlayer() // stores in global slice/map directly
		e := &Enemy{}
		e.addPlayer()
	}

	// for i, v := range playerInterface {
	// 	fmt.Println(i, v)
	// }

	// for i, v := range enemyInterface {
	// 	fmt.Println(i, v)
	// }

	for {

		for i := 0; i < n; i++ {
			// fmt.Printf("\n%s's turn to attack:\nEnter the hit (10-20): ", namedisplay(playerInterface[i]))
			// fmt.Scanf("%d", &dmg)

			// if dmg > 20 || dmg < 10 {
			// 	for {
			// 		fmt.Printf("Enter in valid range (10-20): ")
			// 		fmt.Scanf("%d", &dmg)

			// 		if dmg <= 20 && dmg >= 10 {
			// 			break
			// 		}
			// 	}
			// }

			ehealthReduced := attackOpp(playerInterface[i])

			fmt.Printf("%s's remaining health = %d \n", namedisplay(enemyInterface[i]), hitDamage(enemyInterface[i], ehealthReduced))
			phealthReduced := attackOpp(enemyInterface[i])
			fmt.Printf("%s's remaining health = %d \n", namedisplay(playerInterface[i]), hitDamage(playerInterface[i], phealthReduced))

			// 	dmg := addplayer(playerInterface[i])
			// 	enemyHealth, enemyAttack := enemyInterface.Attack(dmg, 0)
			// 	fmt.Printf("%s's HP reduced by %d\nTotal remaining health = %d.\n", enemyStore[i].Name, dmg, enemyHealth)

			// 	_, playerStore[i].Hp = playerInterface.Attack(0, enemyAttack)

			// 	fmt.Printf("\n%s Attacks!\n%d Damage dealt.\nTotal remaining health of %s = %d\n", enemyStore[i].Name, enemyAttack, playerStore[i].Name, playerStore[i].Hp)

			// 	if playerStore[i].Hp <= 0 {
			// 		win = 0
			// 		fmt.Printf("%s won!", enemyStore[i].Name)
			// 		return
			// 	} else if enemyStore[i].Hp <= 0 {
			// 		win = 1
			// 		fmt.Printf("%s won!", playerStore[i].Name)
			// 		return
			// 	}
			// }
			if win == 1 || win == 0 {
				break
			}
		}
	}
}
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
