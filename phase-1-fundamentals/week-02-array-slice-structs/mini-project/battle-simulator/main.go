package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Random number generator seeded with current time
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Counter function to generate incremental numbers (used for enemy naming)
var c = counter()

// Defines the behavior all characters must have
type Character interface {
	Attack(int) int    // returns the damage dealt
	Hpreduced(int) int // returns remaining HP
	showname() string  // returns the character's name
	Health() int
}

// Player struct
type Player struct {
	Name string
	Hp   int
	Hit  int
}

// Enemy struct
type Enemy struct {
	Name string
	Hp   int
	Hit  int
}

func (p *Player) Health() int {
	return p.Hp
}

func (e *Enemy) Health() int {
	return e.Hp
}

// Function to add a new player and initialize HP
func addPlayer(p *Player) {
	var playerName string
	fmt.Printf("Enter the name of the player: ")
	fmt.Scanf("%s", &playerName) // Read player name from user
	p.Name = playerName
	p.Hp = 100 // Set starting HP
}

// Function to add enemy with a unique name and initialize HP
func addEnemy(e *Enemy) {
	enemyName := fmt.Sprintf("Enemy%d", c()) // Generate name like Enemy1, Enemy2, ...
	e.Name = enemyName
	e.Hp = 100
}

// Reduce the player's HP by the damage taken
func (p *Player) Hpreduced(dmg int) int {
	p.Hp -= dmg
	return p.Hp
}

// Reduce the enemy's HP by the damage taken
func (e *Enemy) Hpreduced(dmg int) int {
	e.Hp -= dmg
	return e.Hp
}

// Return the player's name
func (playername *Player) showname() string {
	return playername.Name
}

// Return the enemy's name
func (enemyname *Enemy) showname() string {
	return enemyname.Name
}

// Player attack method with block or extra damage
func (p *Player) Attack(dmg int) int {
	var prob = r.Intn(8) // Random factor for variability
	if prob%2 == 0 {
		p.Hit = dmg - prob
		fmt.Printf("Damage Blocked! Damage received: %d-%d = %d\n", dmg, dmg-p.Hit, p.Hit)
	} else {
		p.Hit = dmg + prob
		fmt.Printf("Extra Damage! Damage received: %d+%d = %d\n", dmg, p.Hit-dmg, p.Hit)
	}

	return p.Hit
}

// Enemy attack method with random variation
func (e *Enemy) Attack(dmg int) int {
	var eprob = r.Intn(10) // Random factor for variability

	if eprob%2 == 0 {
		e.Hit = dmg - eprob
		fmt.Printf("Damage Blocked! Damage received: %d-%d = %d\n", dmg, dmg-e.Hit, e.Hit)
	} else {
		e.Hit = dmg + eprob
		fmt.Printf("Extra Damage! Damage received: %d+%d = %d\n", dmg, e.Hit-dmg, e.Hit)
	}

	return e.Hit
}

func main() {
	defer fmt.Println("Thanks for playing the game!")
	var n int
	var dmg int

	fmt.Printf("Enter number of players: ")
	fmt.Scanf("%d", &n) // Read number of players

	playerInterface := make([]Character, 0) // Slice of Character interfaces for players
	enemyInterface := make([]Character, 0)  // Slice of Character interfaces for enemies

	// Create players and enemies
	for i := 0; i < n; i++ {
		p := &Player{}
		addPlayer(p) // Get player info from user
		playerInterface = append(playerInterface, p)
		// playerMap = append(playerMap, p)

		e := &Enemy{}
		addEnemy(e) // Create enemy with unique name
		enemyInterface = append(enemyInterface, e)
	}

	// Battle loop
	for {
		for i := 0; i < len(playerInterface); i++ {

			var x int

			for k, v := range enemyInterface {
				fmt.Printf("\n%s's remaining HP = %d\tID = %d", v.showname(), v.Health(), k)
			}

			// Player's turn to attack
			fmt.Printf("\n\n%s's turn to attack:\n", playerInterface[i].showname())
			fmt.Printf("Enter the ID of an enemy you want to attack: ")
			fmt.Scanf("%d", &x)
			for x < 0 || x >= len(enemyInterface) {
				fmt.Printf("Invalid ID. Enter again: ")
				fmt.Scanf("%d", &x)
			}

			fmt.Printf("Enter the hit (10-20): ")
			fmt.Scanf("%d", &dmg)

			// Validate damage input
			for dmg < 10 || dmg > 20 {
				fmt.Printf("Enter in valid range (10-20): ")
				fmt.Scanf("%d", &dmg)
			}

			// Player attacks enemy
			ehealthReduced := playerInterface[i].Attack(dmg)
			enemyHpAfter := enemyInterface[x].Hpreduced(ehealthReduced)
			fmt.Printf("%s's remaining health = %d \n", enemyInterface[x].showname(), enemyHpAfter)

			if enemyHpAfter <= 0 {
				enemyInterface = append(enemyInterface[:x], enemyInterface[x+1:]...)
				i--
				continue
			}
			if len(enemyInterface) == 0 {
				fmt.Println("Players win!")
				return
			}

			// Enemy's turn to attack
			fmt.Printf("\n%s attacked:\n", enemyInterface[x].showname())
			min, max := 15, 20
			damage := r.Intn(max-min+1) + min // Random damage between 15-20

			phealthReduced := enemyInterface[x].Attack(damage)
			playerHealthAfter := playerInterface[i].Hpreduced(phealthReduced)

			fmt.Printf("%s's remaining health = %d \n", playerInterface[i].showname(), playerHealthAfter)
			if playerHealthAfter <= 0 {
				playerInterface = append(playerInterface[:i], playerInterface[i+1:]...)
				i--
				continue

			}

			if len(playerInterface) == 0 {
				fmt.Println("Enemies win!")
				return
			}

		}
	}
}

// Counter function returns a closure that increments count each call
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
