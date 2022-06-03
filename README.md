# write, read and **_RUN_** scripts over ssh

![demo]()


## Installation

```bash
wget https://github.com/tom-power/ssh-run-scripts/releases/download/latest/release_linuxamd64.tar.gz &&  
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

`sshRun` can be dropped from the above using provided `bash_aliases` + your hosts.

### Writing scripts

Write scripts as you would normally for the target shell.

`$host`, `$ip`, `$user`, `$portSsh`, `$portTunnel` will be replaced with host configuration values.

Use the following file name convention to indicate how your scripts should be run:

- `<scriptName>.sh` _with -f_
- `<scriptName>.sudo.sh` _with -t_
- `<scriptName>.x11.sh` _with -Y_
- `<scriptName>.local.sh` _locally_

Save to the following directories in `.config/ssh-run-scripts/scripts` so the script can be found:

- `common`
- `host/<hostName>`
- `shared/<sharedDir>` where `sharedDir` is also found in `host/<hostName>`

See scripts [here](https:#github.com/tom-power/ssh-run-scripts/tree/master/config/.config/ssh-run-scripts/scripts) for examples.

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
