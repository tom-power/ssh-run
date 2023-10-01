# read/write/**_RUN_** scripts over ssh

You have some **hosts**, and some **scripts** you want to run against them without remembering many/any details..

..me too, so wrote this nonsense and use it most days now, **ymmv**!

![demo](https://github.com/tom-power/ssh-run/blob/master/assets/demo.gif)

## Installation

```bash
wget https://github.com/tom-power/ssh-run/releases/download/latest/release_linuxamd64.tar.gz
tar -xvzf release_linuxamd64.tar.gz
cd ./release_linuxamd64

# executables
cp sshRun /path/to/bin/
cp _sshRun /path/to/bin/

# shell integration, first install oh-my-zsh with fzf-tab
cat .bash_aliases | tee -a ~/.bash_aliases
cp _sshRunCompletion ~/.oh-my-zsh/custom/completions/_sshRun
```

## Usage

### Configuration

Hosts will be read from `~/.ssh/config` and all `*.yml` files in `~/.config/ssh-run/`, which have sshRun specific config. 

Use `localhostIs` to give access `localhost` access to the scripts of a `host` in `hosts`.

```yaml
includeSshConfigHosts: true
localhostIs: example
hosts:
  - host: 192.168.0.1
    user: exampleUser
    name: example
    port: 22
    portTunnel: 1080
    checkRemote: true
```

### Running scripts

Run a script against a configured host like:

`sshRun <hostName> <scriptName>`

to discover/explain hosts/scripts:

```
sshRun hosts
sshRun <hostName> explain
sshRun <hostName> scripts
sshRun <hostName> <scriptName> explain
```

also `sshRun <hostName>` or `sshRun <hostName> ssh` to start an ssh session.

drop `sshRun` from the above using the provided `bash_aliases` + your hosts.

### Writing scripts

Write scripts as you would normally for the target shell.

`$host`, `$user`, `$hostName`, `$port`, `$portTunnel` will be replaced with host configuration values.

use the following file name convention to indicate how your scripts should be run:

- `<scriptName>.sh` _with no flags_
- `<scriptName>.pty.sh` _with -t_
- `<scriptName>.x11.sh` _with -Y_
- `<scriptName>.local.sh` _locally_

save to the following directories under `.config/ssh-run/scripts` so the script can be found:

- `common`
- `host/<hostName>` also checked on the remote host if `Config.checkRemote` is `true`
- `shared/<sharedDir>` where `sharedDir` is also found in `host/<hostName>`

see scripts [here](https:#github.com/tom-power/ssh-run/tree/master/config/.config/ssh-run/scripts) for examples.

## Install from source

Install [go](https:#golang.org/), oh-my-zsh with fzf-tab.

```bash
# get repo
git clone https://github.com/tom-power/ssh-run.git
cd ./ssh-run

# build
cp ./example.install.config.sh ./install.config.sh # update config first
./install.sh

# shell integration
cat ./config/.bash_aliases | tee -a ~/.bash_aliases
cp ./config/_sshRunCompletion ~/.oh-my-zsh/custom/completions/_sshRun
```
