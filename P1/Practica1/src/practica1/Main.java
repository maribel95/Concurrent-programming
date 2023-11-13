/*
        COMPONENTS DEL GRUP:
        - Odilo Fortes Domínguez
        - María Isabel Crespí Valero
        VIDEO EXPLICATIU:
        https://drive.google.com/file/d/16pzi7yNdTT7nNtET_5JeEhkXOyOnvJf5/view
*/

package practica1;

import static java.lang.Thread.sleep;
import java.util.Random;
import java.util.concurrent.Semaphore;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Main implements Runnable {

    final static int DONES = 6; // Hi ha 6 dones al despatx
    final static int HOMES = 6; // Hi ha 6 homes al despatx
    final static int NUMPERSBANY = 3; // Número de persones que poden entrar al bany
    // Semàfors utilitzats
    static Semaphore capacitatHomes = new Semaphore(NUMPERSBANY);  // Només poden entrar tres homes al bany
    static Semaphore capacitatDones = new Semaphore(NUMPERSBANY);  // Només poden entrar tres dones al bany
    static Semaphore mutexDona = new Semaphore(1);                 // Semàfor per controlar l'exclusió mútua de les dones
    static Semaphore mutexHome = new Semaphore(1);                 // Semàfor per controlar l'exclusió mútua dels homes
    static Semaphore mutexBany = new Semaphore(1);                 // Semàfor per controlar la síncronia dels sexes

    int id;                          // Id de la persona, servirà per saber quin codi executarà el procés
    String nom;                      // Nom de la persona
    static volatile int homes = 0;   // Contador dels homes al bany
    static volatile int dones = 0;   // Contador de les dones al bany
    Random rnd = new Random();       // Random per generar nombres aleatoris per l'sleep.
    final static int NUMRANDOM = 50; // Límit dels nombres random

    public Main(int id, String nom) {
        this.id = id;
        this.nom = nom;
    }

    public static void main(String[] args) throws InterruptedException {
        // Noms dels processos
        String[] noms = {"CATALINA", "SEBASTIA", "CARME", "ONOFRE", "CONXA", "GUILLEM", "XESCA", "PAU", "ANTONIA", "JAUME", "MARGA", "TONI"};
        // Array dels processos, tant de dones com dels homes
        Thread[] threads = new Thread[DONES + HOMES];
        // Inicialitzam tots els processos
        for (int i = 0; i < DONES + HOMES; i++) {
            threads[i] = new Thread(new Main(i, noms[i]));
            threads[i].start();
        }
        // Esperam a que acabin tots els processos
        for (int i = 0; i < DONES + HOMES; i++) {
            threads[i].join();
        }
    }

    @Override
    public void run() {
        arribaDespatx();                // Cada persona arribarà al despatx només una vegada
        for (int i = 1; i <= 2; i++) {  // Cada persona ha d'anar al bany o treballar almenys dues vegades
            treballa();                 // La persona treballa
            if (id % 2 == 0) {          // DONES
                try {
                    banyDones(i);       // Executam el codi per a que les dones puguin aanar al bany
                } catch (InterruptedException ex) {
                    Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
                }
            } else {                    // HOMES                
                try {
                    banyHomes(i);       // Executam el codi per a que els homes puguin aanar al bany
                } catch (InterruptedException ex) {
                    Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
                }
            }
        }

    }

    // Mètode que simplement fa un sleep d'un temps aleatori
    public void dormir() throws InterruptedException {
        sleep(rnd.nextInt(NUMRANDOM));
    }

    // Mètode que simplement imprimeix l'arribada al despatx
    public void arribaDespatx() {
        System.out.println(this.nom + " arriba al despatx.");
    }

    // Mètode que simplement imprimex que la persona està treballant
    public void treballa() {
        System.out.println(this.nom + " treballa.");
    }

    // Mètode que controla quan les dones volen entrar al bany
    public void banyDones(int i) throws InterruptedException {
        mutexBany.acquire();            // Mutex general per ambdós sexes i que només pot executar una sola persona
        capacitatDones.acquire();       // Controlam la capacitat de les dones entrant al bany
        mutexDona.acquire();            // Mutex de les dones que controla que el seu contador s'incrementi correctament  
        dones++;                        // Hi ha una dona més al bany     
        System.out.println(this.nom + " entra " + i + "/2. Dones al bany: " + dones);
        mutexDona.release();            // També alliberam el mutex de la dona
        if (dones == 1) {               // En cas de que sigui la primera dóna que ha entrat al bany
            mutexHome.acquire();        // Llavors podem bloquejar als homes que vulguin entrar, hauran d'esperar afora
        }       
        mutexBany.release();            // Alliberam el mutex genèric
        dormir();                       // La dona està una estona dins el bany abans de sortir
        mutexDona.acquire();            // Tornam a demanar permís al mutex de la dona per fer el decrement
        dones--;                        // Una dona surt del bany
        System.out.println(this.nom + " surt del bany.");
        mutexDona.release();            // Alliberam el mutex de la dona
        if (dones == 0) {               // Si no queda cap dona dins el bany
            System.out.println("*** El bany està buit ***");
            mutexHome.release();        // Alliberam als homes que esperaven per entrar dins el bany
        }        
        capacitatDones.release();        // Alliberam un espai més al bany per a que entri una altra dona si fos el cas
        dormir();                       // Feim un petit sleep
    }

    // Mètode que controla quan els homes volen entrar al bany
    // El codi no s'ha comentat ja que és simètric respecte al de la dona
    public void banyHomes(int i) throws InterruptedException {
        mutexBany.acquire();
        capacitatHomes.acquire();
        mutexHome.acquire();            // Mutex de les dones que controla que el seu contador s'incrementi correctament  
        homes++;        
        System.out.println(this.nom + " entra " + i + "/2. Homes al bany: " + homes);
        mutexHome.release();
        if (homes == 1) {
            mutexDona.acquire();
        }
        mutexBany.release();
        dormir();                   // L'home està una estoneta dins el bany abans de sortir
        mutexHome.acquire();
        homes--;
        System.out.println(this.nom + " surt del bany.");
        mutexHome.release();
        if (homes == 0) {
            System.out.println("*** El bany està buit ***");
            mutexDona.release();
        }        
        capacitatHomes.release();
        dormir();
    }

}
