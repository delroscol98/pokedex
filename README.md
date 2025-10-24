# Pokedex REPL

This project was made on the [Boot.dev](https://www.boot.dev/courses/build-pokedex-cli-golang) learning platform to build a REPL pokedex on the command line using Go, practicing skills in JSON, network requests and caching.

## How to use

1. Make a clone of this repo using the following command

```
git clone git@github.com:delroscol98/pokedex.git
```

2. Move into the pokedex directory

```
cd pokedex
```

3. Run the pokedex REPL program

```
./pokedex
```

For a short description of available commands enter the `help` then press `<ENTER>` to execute the command and you will see the following output:

```
Pokedex > help
Welcome to the Pokedex!
Usage:

pokedex: displays all pokemon in pokedex
exit: Exit the Pokedex
help: displays a help message
map: lists the next 20 locations
mapb: lists the previous 20 locations
explore: lists all the pokemon in the specified location
catch: catches a pokemon in the explored location
inspect: displays information about pokemon in the poked
ex
```

### Available Commands

Let me elaborate further on what each command does. NOTE: I recommend following along in the given order!

#### `map` and `mapb`

The `map` command fetches 20 location areas from the [PokeAPI](https://pokeapi.co/).

> From the PokeAPI Docs: Location areas are sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokemon encounters

Entering the `map` command subsequent times produces a new set of 20 location areas.

For example:

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

If you want to see the previous 20 location areas then use the `mapb` command. NOTE: If you use the `mapb` command on the first set of 20 locations, the following message will output: `"You're on the first page"`.

For example:

```
Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
Pokedex > mapb
You're on the first page
```

#### `explore`

Once you have a place in mind you would like to "explore" then use the `explore` command followed by the name of the location area you would like to "explore". The output will be the list of pokemon that can be found in that location area.

For example:

```
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
- tentacool
- tentacruel
- magikarp
- gyarados
- remoraid
- octillery
- wingull
- pelipper
- shellos
- gastrodon
```

#### `catch`

At this point you might want to catch some pokemon! Use the `catch` command followed by the name of the pokemon you want to catch. Be careful though... there's a chance the pokemon might escape.

```
Pokedex > catch gastrodon
Throwing a Pokeball at gastrodon...
gastrodon was caught!
You may now inspect it with the inspect command
Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool escaped!
```

#### `inspect`

As the output of the above example says, once a Pokemon has been caught it can be inspected in the Pokedex using the `inspect` command followed by the name of the Pokemon that exists in the Pokedex.

```
Pokedex > inspect gastrodon
Name: gastrodon
Height: 9
Stats:
- hp: 40
- attack: 40
- defense: 35
- special-attack: 50
- special-defense: 100
- speed: 70
Types:
- water
- poison
```

#### `pokedex`

At any moment you can use the `pokedex` command to see the names of the pokemon you have caught and thus inspect

```
Pokedex > pokedex
Your Pokedex:
- gastrodon
```

#### `exit`

The `exit` command simply ends the program.

```
Pokedex > exit
Closing the Pokedex... Goodbye!
```

## What did I learn

This project allowed me to practice my newly acquired skills with Go programming language. It was my first time creating a REPL (Read, Evaluate, Print Loop) that takes user inputs which was super fun.
In my programming journey working with APIs and HTTP requests have always been a point of struggle for me, but building this project greatly improved my confidence in working with third-party APIs. Sometimes I just got to RTFM and all will be well.

## Future improvements

Some time in the future I'd like to come back to this project and extend it even further:

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
  Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

If you'd like to collaborate with one of the above ideas, I'd love to work with you!

Thanks for reading!
