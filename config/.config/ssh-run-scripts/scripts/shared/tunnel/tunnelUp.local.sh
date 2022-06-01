autossh -f -M 0 -o \
\"ServerAliveInterval 10\" -o \
\"ServerAliveCountMax 3\" \
-D localhost:$sshTunnel -N $user@$ip
