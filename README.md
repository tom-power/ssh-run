# write/read/**_RUN_** scripts over ssh

This must exist elsewhere, or is easy enough to configure so isn't needed..

but I couldn't find either, so started writing this nonsense and use it everyday now, ymmv!

![demo]()


## Installation

```bash
wget https://github.com/tom-power/ssh-run-scripts/releases/download/latest/release_linuxamd64.tar.gz&&  
tar -xvzf release_linuxamd64.tar.gz &&
cd ./release_linuxamd64

# executables
cp sshRun /path/to/bin/ &&
cp _sshRun /path/to/bin/

# shell integration, first install oh-my-zsh with fzf-tab
cat .bash_aliases | tee -a ~/.bash_aliases &&
cp _sshRunCompletion ~/.oh-my-zsh/custom/completions/_sshRun
```

## Usage

### Configuring hosts

Add hosts to `~/.config/ssh-run-scripts/config.yaml`:

```yaml
hosts:
  - name: example
    user: exampleUser
    ip: 192.168.0.1
    portSsh: 22
    portTunnel: 1080
    checkForScripts: true
```

### Running scripts

Run a script against a configured host like:

`sshRun <hostName> <scriptName>`

to discover/explain hosts/scripts or start an ssh session:

```
sshRun hosts
sshRun <hostName> explain
sshRun <hostName> scripts
sshRun <hostName> <scriptName> explain
sshRun <hostName> ssh
```

drop `sshRun` from the above using the provided `bash_aliases` + your hosts.

### Writing scripts

Write scripts as you would normally for the target shell.

`$host`, `$ip`, `$user`, `$portSsh`, `$portTunnel` will be replaced with host configuration values.

use the following file name convention to indicate how your scripts should be run:

- `<scriptName>.sh` _with no flags_
- `<scriptName>.pty.sh` _with -t_
- `<scriptName>.x11.sh` _with -Y_
- `<scriptName>.local.sh` _locally_

save to the following directories under `.config/ssh-run-scripts/scripts` so the script can be found:

- `common`
- `host/<hostName>` also checked on the host if `Host.checkForScripts` is `true` 
- `shared/<sharedDir>` where `sharedDir` is also found in `host/<hostName>`

see scripts [here](https:#github.com/tom-power/ssh-run-scripts/tree/master/config/.config/ssh-run-scripts/scripts) for examples.

## Install from source

Install [go](https:#golang.org/), oh-my-zsh with fzf-tab.

```bash
# get repo
git clone https://github.com/tom-power/ssh-run-scripts.git &&
cd ./ssh-run-scripts &&

# build
cp ./example.install.config.sh ./install.config.sh && # update config first
./install.sh

# shell integration
cat ./config/.bash_aliases | tee -a ~/.bash_aliases
cp ./config/_sshRunCompletion ~/.oh-my-zsh/custom/completions/_sshRun
```
