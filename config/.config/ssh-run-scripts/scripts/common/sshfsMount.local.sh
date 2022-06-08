sshfs -p $port $user@$host:/home ~/mount/sshfs &&
cd ~/mount/sshfs/$user &&
exec $SHELL