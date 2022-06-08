ssh-keygen -t rsa -f ~/.ssh/id_rsa_$hostName &&
ssh-add ~/.ssh/id_rsa_$hostName &&
ssh-copy-id -i ~/.ssh/id_rsa_$hostName.pub $user@$host