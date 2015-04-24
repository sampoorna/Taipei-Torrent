__author__ = 'wali'
#filters output
import subprocess
import os
import signal
import sys
import shutil
import time
from subprocess import call

while True:
    time.sleep(2)
    shutil.rmtree(sys.argv[1])
    dir='-fileDir=%s' % str(sys.argv[1])
    proc = subprocess.Popen(['go','run','main.go',dir,"testData/main.torrent"],stdout=subprocess.PIPE)
    while True:
        line = proc.stdout.readline()
        if line != '':
            #the real code does filtering here
            print line.rstrip()
            if ("good, total 122 122" in line):
                os.system("kill `lsof -t -i:7777`")
                #os.kill(proc.pid, signal.SIGINT)
                print("Killed it... Restarting")
                break
