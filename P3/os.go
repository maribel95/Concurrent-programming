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

//---------------------------------      		 VARIABLES				------------------------------------------------------
var (
	repeticionsOs = 0
	//contador = 0
	body = "0"
)

//--------------------------------------	   FUNCIÓ GESTIÓ ERRORS	    	--------------------------------------------------
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//--------------------------------------				MAIN		   		----------------------------------------------------
func main() {
	//*********************************** ESTABLIM LA CONEXIÓ AMB RABBITMQ ******************************************************
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Ens conectam al servidor RabbitMQ
	failOnError(err, "Failed to connect to RabbitMQ")            // Recollim possible error
	defer conn.Close()                                           // Quan el programi acabi tancam la conexió
	//****************************************************************************************************************************

	//*************************************** ESTABLIM UN CANAL DE CONEXIÓ *******************************************************
	ch, err := conn.Channel()                    // Definim canal de conexió anomenat ch
	failOnError(err, "Failed to open a channel") // Recollim possible error
	defer ch.Close()                             // Quan el programi acabi tancam el canal
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

	//********************************************* CONSUMIM MISSATGE DE LA COA **************************************************
	msgs, err := ch.Consume(
		avisosAbella.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
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

	// Esta función se ejecuta nada más recibir un mensaje de permiso de la abeja
	forever := make(chan bool)

	//********************************************** L'OS ENVIA 10 PERMISOS **************************************************
	log.Printf("[*] L'os dorm si no li donen menjar")

	for contador := 1; contador <= 10; contador++ { // L'os envia 10 missatges
		err = ch.Publish(
			"",               // exchange
			permisosOso.Name, // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,                // Per preservar la persistència dels missatges
				ContentType:  "text/plain",                   // Publicam un text
				Body:         []byte(strconv.Itoa(contador)), // El cos sempre serà un array de bytes
			})
		failOnError(err, "Failed to publish a message") // Agafam possible error
		time.Sleep(3000 * time.Millisecond)

		// HAY QUE PONER ALGO PARA QUE DESPUÉS DE HABER ENVIADO CADA MENSAJE SE QUEDE ESPERANDO A LA RESPUESTA DE LAS ABEJAS
	}

	go func() {

		for d := range msgs {
			repeticionsOs++
			log.Printf("L'ha despertat l'abella: %s, i menja %d/3", d.Body, repeticionsOs)
			imprimirPuntos(5)
			//****************************************************************************************************************************
			if repeticionsOs == 3 { // FIN PROGRAMA HAY QUE CAMBIARLO A == 3
				err = ch.Publish(
					"logs", // exchange
					"",     // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,  // Per preservar la persistència dels missatges
						ContentType:  "text/plain",     // Publicam un text
						Body:         []byte("ROMPUT"), // El cos sempre serà un array de bytes
					})
				failOnError(err, "Failed to publish a message")       // Agafam possible error
				log.Printf("L'os està ple, ha romput el pot!!!")      // Missatge de pot romput
				imprimirPuntos(5)                                     // Espera 5 segons a que acabin tots els procesos
				ch.QueueDelete(permisosOso.Name, false, false, true)  // Eliminam cua dels permisos
				ch.QueueDelete(avisosAbella.Name, false, false, true) // Eliminam cua dels avisos
				ch.ExchangeDelete("logs", false, false)               // També eliminam s'exchange
				log.Printf("Eliminant cues")                          // Missatge confirmació eliminació cues
				log.Printf("SIMULACIÓ ACABADA")                       // Missatge de simulació acabada
				ch.Close()                                            // Tancam el canal
				conn.Close()                                          // Tancam la conexió
				os.Exit(0)                                            // Fi de la simulació
			}
			// UNA VEGADA ENVIATS ELS PERMISOS, L'OS JA HA ACABAT
			log.Printf("L'os se'n va a dormir")
			//********************************************** L'OS ENVIA 10 PERMISOS **************************************************
			for contador := 1; contador <= 10; contador++ { // L'os envia 10 missatges
				err = ch.Publish(
					"",               // exchange
					permisosOso.Name, // routing key
					false,            // mandatory
					false,            // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,                // Per preservar la persistència dels missatges
						ContentType:  "text/plain",                   // Publicam un text
						Body:         []byte(strconv.Itoa(contador)), // El cos sempre serà un array de bytes
					})
				failOnError(err, "Failed to publish a message") // Agafam possible error
				time.Sleep(3000 * time.Millisecond)
			}

		}
	}()

	<-forever

}

// Funció auxiliar que el que farà és imprimir els punts d'espera de les abelles o l'os
func imprimirPuntos(n int) { // Li passam el número de punts(segons) que esperarà
	for i := 0; i < n; i++ { // Bucle for
		fmt.Print(".")                      // Imprimim el punt
		time.Sleep(1000 * time.Millisecond) // L'abella tarda una estoneta en produir
	} // Fi del bucle
	fmt.Print("\n") // Finalmente imprimim un salt de línea
}
