/*
   COMPONENTS DEL GRUP:
   - Odilo Fortes Domínguez
   - María Isabel Crespí Valero
   VIDEO EXPLICATIU:
   https://drive.google.com/file/d/13eOFBH0nqzbpsvTb7JZm85L3moJpHO-b/view
*/

package main

//---------------------------------      		 IMPORTS				------------------------------------------------------
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	amqp "github.com/streadway/amqp"
)

//----------------------------------      		 VARIABLES				------------------------------------------------------
var (
	contador = 0   // Valor numèric del contador
	body     = "0" // Valor en string del contador(degut a que amb la coa intercanviam arrays de bytes)
)

//----------------------------------      		 CONSTANTS				------------------------------------------------------
const (
	BUFSIZE = 10 // En el pot només caben 10 unitats de mel
)

//--------------------------------------	   FUNCIÓ GESTIÓ ERRORS	    	--------------------------------------------------
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//--------------------------------------				MAIN		   		----------------------------------------------------
func main() {
	// Recollim de la linea de comandes el nom de l'abella
	args := os.Args[1:]                                 // Guardam un array a partir el primer argument
	nomAbella := args[0]                                // El primer element d'aquest array és el nom de l'abella
	log.Printf("[*] Aquesta és l'abella %s", nomAbella) // Imprimim el nom de l'abella per pantalla

	//*********************************** ESTABLIM LA CONEXIÓ AMB RABBITMQ ******************************************************
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Ens conectam al servidor RabbitMQ
	failOnError(err, "Failed to connect to RabbitMQ")            // Recollim possible error
	defer conn.Close()                                           // Quan el programi acabi tancam la conexió
	//****************************************************************************************************************************

	//*************************************** ESTABLIM UN CANAL DE CONEXIÓ *******************************************************
	ch, err := conn.Channel()                    // Establim canal de conexió
	failOnError(err, "Failed to open a channel") // Agafam possible error
	defer ch.Close()                             // Quan el programi acabi tancam la sessió
	//****************************************************************************************************************************

	//********************************************** ESTABLIM UNA COA ************************************************************
	permisosOso, err := ch.QueueDeclare(
		"permisosOso", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")
	//****************************************************************************************************************************

	//*************************************** ESTABLIM UNA SEGONA COA ************************************************************
	avisosAbella, err := ch.QueueDeclare(
		"avisosAbella", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")
	//****************************************************************************************************************************

	//********************************* CONSUMIM MISSATGE DE LA COA DELS PERMISOS DE L'ÓS ****************************************
	msgs, err := ch.Consume(
		permisosOso.Name, // queue
		"",               // consumer
		false,            // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	failOnError(err, "Failed to register a consumer")
	//****************************************************************************************************************************

	//**************************************************  DECLARAM EXCHANGE  *****************************************************

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//****************************************************************************************************************************

	//*****************************************  DECLARAM UN BIND D'UNA CUA  *****************************************************
	err = ch.QueueBind(
		permisosOso.Name, // queue name
		"",               // routing key
		"logs",           // exchange
		false,
		nil,
	)
	//****************************************************************************************************************************

	//*********************************************  DECLARAM PREFETCH  **********************************************************

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	//****************************************************************************************************************************

	//****************************************** FUNCIÓ GO PER LLEGIR MISSATGES **************************************************
	forever := make(chan bool)
	// Funció anònima que s'executarà en la mateixa línea go
	go func() {
		// Això s'executarà cada vegada que rebem un missatge(permís per produir) de l'ós
		for d := range msgs { // Per cada un dels permisos farem lo següent
			// En cas de que el missatge sigui ROMPUT, significa que la simulació s'ha d'acabar
			if string(d.Body) == "ROMPUT" {
				// Missatges de finalització de la abella
				log.Printf("El pot està romput, no es pot omplir de mel!!!")
				log.Printf("L'abella %s se'n va", nomAbella)
				// Enviam missatges d'acabament a les altres abelles a través del fnaout
				err = ch.Publish(
					"logs", // exchange
					"",     // routing
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,  // Per preservar la persistència dels missatges
						ContentType:  "text/plain",     // Publicam un text
						Body:         []byte("ROMPUT"), // El cos sempre serà un array de bytes
					})
				failOnError(err, "Failed to publish a message") // Agafam possible error
				log.Printf("SIMULACIÓ ACABADA")                 // Aquesta abella ha acabat la seva simulació
				os.Exit(0)                                      // Feim un exit del procés
			} // Fi del condicional
			// En cas de que no hagi arribat el missatge ROMPUT, significa que la simulació ha de continuar
			log.Printf("L'abella %s produeix mel %s", nomAbella, d.Body) // L'abella produeix mel
			imprimirPuntos(3)                                            // Les abelles tarden una estoneta en produir la mel
			contador, err = strconv.Atoi(string(d.Body))                 // Feim la conversió del missatge del body a enter
			// Si es la dècima porció de mel, llavors l'abella ha de despertar l'ós per a que menji
			if contador == BUFSIZE {
				log.Printf("L'abella %s desperta l'ós", nomAbella) //Llegim el missatge i el treim de la coa
				err = ch.Publish(
					"",                // exchange
					avisosAbella.Name, // routing
					false,             // mandatory
					false,             // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,   // Per preservar la persistència dels missatges
						ContentType:  "text/plain",      // Publicam un text
						Body:         []byte(nomAbella), // El cos sempre serà un array de bytes
					},
				)
				failOnError(err, "Failed to publish a message") // Agafam possible error
			} // Fi condicional del contador
			d.Ack(false) // Para que nos se pierdan los mensajes
		} // Fi for

	}() // Fi funció anònima
	<-forever // Espera eternament
	//****************************************************************************************************************************
}

// Funció auxiliar que el que farà és imprimir els punts d'espera de les abelles o l'os
func imprimirPuntos(n int) { // Li passam el número de punts(segons) que esperarà
	for i := 0; i < n; i++ { // Bucle for
		fmt.Print(".")                      // Imprimim el punt
		time.Sleep(1000 * time.Millisecond) // L'abella tarda una estoneta en produir
	} // Fi del bucle
	fmt.Print("\n") // Finalmente imprimim un salt de línea
}
