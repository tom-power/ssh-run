sshfs -p $port $user@$ip:/home ~/mount/sshfs &&
cd ~/mount/sshfs/$user &&
exec $SHELL
