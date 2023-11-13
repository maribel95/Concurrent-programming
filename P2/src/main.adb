--        COMPONENTS DEL GRUP:
--        - Odilo Fortes Domínguez
--        - María Isabel Crespí Valero
--        VIDEO EXPLICATIU:
--        https://drive.google.com/file/d/1BLuUON2Yc8gV1DZqEeRO7xHCECvW2ke6/view


--------------------------------------------------------------------------------
---------------------        ÚTILS DEL PROGRAMA         ------------------------
--------------------------------------------------------------------------------

with Text_IO; use Text_IO;          -- Per entrada i sortida de dades
with def_corda; use def_corda;      -- El nostre monitor corda
with ada.numerics.discrete_random;  -- Per poder generar nombres aleatoris

procedure Main is

   -----------------------------------------------------------------------------
   ------------------------        VARIABLES         ---------------------------
   -----------------------------------------------------------------------------

   -------------------------       CONSTANTS         ---------------------------

   BABUINS_SUD : constant integer := 5;                -- Hi ha 5 babuins al sud
   BABUINS_NORD : constant integer := 5;               -- I 5 babuins al nord
   NUM_RECORREGUTS: constant integer := 3;             -- Nombre de recorreguts que ha de fer cada babui

   -------------------------       VARIABLES         ---------------------------

   MAX_TASQUES: Integer := BABUINS_NORD + BABUINS_SUD; -- El total de fils del programa serà la suma de tots els babuin
   type randRange is range 1..5;                       -- Com a màxim un babuí pot tardar 5 segons en donar la volta
   corda : cordaMonitor;                               -- Tipus protegit per a la SC
   Type sentido is (Nord,Sud);                         -- Enumerat dels possibles sentits del babuins

   -------------------------       PAQUETS         -----------------------------

   package Rand_Int is new ada.numerics.discrete_random(randRange); -- Cridam a la biblioteca per poder utilitzar randoms
   use Rand_Int;                                                    -- Habilitam per poder utilitzar la bilbioteca

   gen : Generator;                                                 -- Generador de nombres aleatoris


   -----------------------------------------------------------------------------
   --------------------         TASCA CONCURRENT         -----------------------
   -----------------------------------------------------------------------------
   task type congost_babuins is                                                      -- Definició de la tasca dels babuins al congost
      entry Start(Idx : in integer; procedencia: in sentido; direccion :in sentido); -- Assignam un identificador, una procedencia i una direcció
   end congost_babuins;                                                              -- Fi de la definició de la tasca

   task body congost_babuins is
      My_idx : Integer;          -- Declaram l'index
      My_procedencia : sentido;  -- Declaram la procedencia
      My_direccion : sentido;    -- Declaram la direcció
   begin
      accept Start(Idx : in integer; procedencia: in sentido; direccion :in sentido) do
         My_idx := Idx;                  -- Assignam l'index
         My_procedencia := procedencia;  -- Assignam la procedencia
         My_direccion := direccion;      -- Assignamaram la direcció
      end Start;
      -- Presentació del babuí
      Put_Line("BON DIA, som el babui "& My_procedencia'Image &""& My_Idx'img & " i vaig cap al "& My_direccion'Image);
      -- El babuins fan el recorregut 3 vegades
      for i in 1..NUM_RECORREGUTS loop
         if(My_procedencia = Nord) then  -- BABUI DEL NORD

            -- Lock per als babuins del nord a la corda
            corda.babuiNordLock;
            -- El babui ha pogut entrar a la corda
            Put_Line(My_procedencia'Image &""& My_Idx'img &": És a la corda i travessa cap al "& My_direccion'Image);
            -- El babui surt de la corda
            corda.babuiNordUnlock;
            -- Ha arribat al final de la corda
            Put_line(My_procedencia'Image &""& My_Idx'img &" ha arribat a la vorera");
            -- Fa la volta a la muntanya per tornar a pujar-se a la corda
            delay Duration(Rand_Int.Random(gen)); -- Lo que tarda el babui en tornar a donar la volta

         else                            -- BABUI DEL SUD, el codi es un mirall del babui del NORD

            corda.babuiSudLock;
            Put_Line(My_procedencia'Image &""& My_Idx'img &": És a la corda i travessa cap al "& My_direccion'Image);
            corda.babuiSudUnlock;
            Put_line(My_procedencia'Image &""& My_Idx'img &" ha arribat a la vorera");
            delay Duration(Rand_Int.Random(gen)*2); -- Lo que tarda el babui en tornar a donar la volta

         end if;
            -- Després de donar la volta a la muntanya, el babuí imprimeix quantes voltes duu
            if(i = NUM_RECORREGUTS) then -- Si es la última, fa un print més efusiu perquè ja acaba
               put_line(My_procedencia'Image &""& My_Idx'img &": Fa la volta "& i'Image  &" de "&NUM_RECORREGUTS'Image & " i acaba!!!!!!!");
            else
               Put_line(My_procedencia'Image &""& My_Idx'img &": Fa la volta "& (i'Image ) &" de "&NUM_RECORREGUTS'Image);
            end if;
      end loop;
   end congost_babuins;

   -----------------------------------------------------------------------------
   ---------------------       PROGRAMA PRINCIPAL         ----------------------
   -----------------------------------------------------------------------------


  type array_babuins is array (1..MAX_TASQUES) of congost_babuins; -- Definim Array on guardam la tasca del congost
   ab : array_babuins;                                             -- Declaram Array de les tasques
   i: Integer := 1;                                                -- Index dels babuins


begin

   for Idx in 1..MAX_TASQUES loop               -- Llançam les tasques dels 5 babuins del nord + 5 babuins del sud
      Rand_Int.Reset(gen);                      -- Reset del random
      if(Idx mod 2 = 1) then                    -- Començarem pels babuins del nord
         ab(Idx).Start(i, Nord, Sud);           -- Els babuins del nord es llançaran si el residu es 1
      else                                      -- Sino...
         ab(Idx).Start(i,Sud, Nord);            -- Els babuins del nord es llançaran si el residuo es 0
         i:= i + 1;                             -- L'índex dels babuins només s'incrementará cada dues pasades del bucle
      end if;

  end loop;
end Main;








