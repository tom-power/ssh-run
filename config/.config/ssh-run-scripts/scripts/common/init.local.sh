ssh-keygen -t rsa -f ~/.ssh/id_rsa_$host &&
ssh-add ~/.ssh/id_rsa_$host &&
ssh-copy-id -i ~/.ssh/id_rsa_$host.pub $user@$ip