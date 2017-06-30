package main

import(
	"log",
	"os"	
	
)
 func init(){
	 // Muda o dispositivo fr logging para o stdout
	 log.SetOutput(os.Stdout)
 }

 func main (){
	 // Pesquisa o termo específico
	 search.run("presidente")
 }