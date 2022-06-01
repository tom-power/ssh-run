sshfs -p $portSsh $user@$ip:/home ~/mount/sshfs && \
cd ~/mount/sshfs/$user && \
exec $SHELL