__author__ = 'wali'
#filters output
import subprocess
import os
import signal

while True:
  proc = subprocess.Popen(['java','test'],stdout=subprocess.PIPE)
  while True:
    line = proc.stdout.readline()
    if line != '':
      #the real code does filtering here
      print "test:", line.rstrip()
      if ("End" in line):
          proc.kill()
          print("Killed it... Restarting")
          break
