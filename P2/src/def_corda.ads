package def_corda is

   protected type cordaMonitor is
      entry babuiSudLock;
      procedure babuiSudUnlock;
      entry babuiNordLock;
      procedure babuiNordUnlock;


   private
      babuinsSud: integer := 0;
      babuinsNord: integer := 0;
   end cordaMonitor;

end def_corda;
