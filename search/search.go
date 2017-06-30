package search

import (
	"log"
	"os"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// Obtém a lista de feeds para pesquisa
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Cria um canal sem buffer parar receber os resultados das correspondências
	result := make(chan *Result)

	// Cria um grupo de espera para que possamos processsasr todos os feeds
	var waitGroup sync.waitGroup

	//Define o número de goroutines que precisamos esperar euqnato 
	// elas processam os feeds individuais
	waitGroup.Add(len(feeds))

	// Inicia uma gorountine para cada feed a fim de obter os resultados
	for	_, feed := range feeds{
		// Obtém um matcher para a pesquisa
		matcher, exists := matchers[feed.type]
		if !exists{
			matcher = matchers["default"]
		}
	}
}
