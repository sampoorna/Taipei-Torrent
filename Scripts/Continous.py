__author__ = 'wali'
#filters output
import subprocess
import os
import signal
import sys
import time

for x in range(1,10):
    time.sleep(2)
    filename= str(sys.argv[1]) + `x` + "txt";
    text_file = open(filename, "w")
    proc = subprocess.Popen(['go','run','main.go',"testData/main.torrent"],stdout=subprocess.PIPE)
    while True:
        line = proc.stdout.readline();
        text_file.write(line)
        if line != '':
            #the real code does filtering here
            print "test:", line.rstrip()
            if ("Percent complete  100 %" in line):
                os.system("kill `lsof -t -i:7777`")
                os.system("rm American.Dad.S11E09.HDTV.x264-KILLERS.mp4")
                #os.kill(proc.pid, signal.SIGINT)
                print("Killed it... Restarting")
                break
    text_file.close()
