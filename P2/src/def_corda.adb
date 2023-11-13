with Text_IO; use Text_IO;

package body def_corda is

   protected body cordaMonitor is                                   -- Objecte protegit d'ADA

      entry babuiSudLock when (babuinsNord = 0) and (babuinsSud /= 3) is  -- Si hi ha algun babui del nord a la corda o ja hi ha tres babuins del sud , hem d'esperar
      begin                                                         -- Comença el mètode
         babuinsSud:= babuinsSud + 1;                               -- Sino, llavors poden entrar els babuins del sud
         Put_Line("***** A la corda n'hi ha "& babuinsSud'Image &" direcció Sud *****");
      end babuiSudLock;                                             -- fi del mètode

      procedure babuiSudUnlock is                                   -- Procediment per permetre la continuació dels babuins sobre la corda
      begin                                                         -- Comença el procedure
         babuinsSud:= babuinsSud - 1;                               -- Un babui s'en va de la corda
         if(babuinsSud = 0) then
            Put_Line("***** A la corda n'hi ha "& babuinsSud'Image &" direcció Sud *****");
         end if;
      end babuiSudUnlock;                                           -- Fi del procedure

      entry babuiNordLock when (babuinsSud = 0) and (babuinsNord /= 3) is  -- Si hi ha algun babui del nord a la corda o ja hi ha tres babuins del sud , hem d'esperar
      begin                                                         -- Comença el mètode
         babuinsNord:= babuinsNord + 1;                               -- Sino, llavors poden entrar els babuins del sud
         Put_Line("+++++ A la corda n'hi ha "& babuinsNord'Image &" direcció Nord +++++");
      end babuiNordLock;                                             -- fi del mètode

      procedure babuiNordUnlock is                                  -- Procediment per permetre la continuació dels babuins sobre la corda
      begin                                                         -- Comença el procedure
         babuinsNord:= babuinsNord - 1;                             -- Un babui s'en va de la corda
         if(babuinsNord = 0) then
            Put_Line("+++++ A la corda n'hi ha "& babuinsNord'Image &" direcció Nord +++++");
         end if;
      end babuiNordUnlock;                                           -- Fi del procedure

   end cordaMonitor;

end def_corda;
