import subprocess
import sys

# with urllib.request.urlopen('http://localhost:8081/bubble') as response:
#    html = response.read()

process = subprocess.call(['F:\\Repositories\\github\\invowk\\invowk-cli\\invowk-cli.exe', 'tui'],
                          stdin=sys.stdin,
                          stdout=sys.stdout,  # needed for the next line to be sensible
                          stderr=sys.stderr)