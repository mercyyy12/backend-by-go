# Turn-Based Battle Game

A simple **turn-based battle game** where players and enemies take turns attacking each other until one side wins.

---

## Go concepts

- Structs with methods
- Interfaces
- Pointer receivers for mutation
- Slices for dynamic data storage
- Maps for optional lookup
- Random numbers with math/rand
- Closures for unique IDs
- Loops and conditionals
- User input with fmt.Scanf
- defer for cleanup message

---

## Features

- Player and Enemy structs with health points (HP) and attack logic.
- Character interface to unify Player and Enemy behaviors.
- Turn-based combat with random variations:
  - Extra damage
  - Damage blocked
- Input validation for player attack selection and damage values.
- Removal of defeated characters from the battle.
- Winner detection for both players and enemies.
- Uses pointer receivers for HP mutation to reflect real changes.

---


