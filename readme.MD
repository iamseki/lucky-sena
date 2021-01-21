# Lucky Sena

<p align="center">Lucky Sena is a random number generator between 01 and 60 <i>closed interval</i></p>

## Usage with cli

Run with go e.g `go run cmd/cli/main.go --b=10 --e=1,2,3,4` or execute the properly binary of your *OS*:

Linux

        ./generate-bet --b=7 --e=2,16,19,31,43,60,27,35,39,50,59,5,10,29,34,41 --c=2336 --p

Windows

        generate-bet --b=7 --e=2,16,19,31,43,60,27,35,39,50,59,5,10,29,34,41 --c=2336 --p


### Defaults

*`--b`* How many sequence to generate, default value is ***7***

*`--e`* Numbers that must to be exclude of random numbers, default value is just the number ***0***

---

## Compiling

- There are some examples [here](run-example.sh)
