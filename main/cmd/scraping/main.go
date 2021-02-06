package main

import (
	"log"
	"lucky-sena/main/factories"
)

func main() {
	scrapper := factories.NewScrappingLastBetUseCase()
	bet := scrapper.Scrap("http://www1.caixa.gov.br/loterias/loterias/megasena/megasena_pesquisa_new.asp")
	log.Println("Last Winner bet:", bet)
}
