# Run scripts against hosts over ssh

[![ci status][badge]][workflow]

[workflow]: https://github.com/tom-power/ssh-run/actions/workflows/go.yml
[badge]: https://img.shields.io/github/actions/workflow/status/tom-power/ssh-run/go.yml?style=flat-round&logo=github&label=CI%20status

Binary and zsh completions to help running **scripts** against **hosts** without remembering many/any details, I use it most days, **ymmv**!

![demo](https://github.com/tom-power/ssh-run/blob/main/assets/demo.gif)

## Usage

```
Run scripts against hosts over ssh

Usage:
  sshRun --hosts                
  sshRun <host> [--ssh|--scripts|--explain]
  sshRun <host> <script> [--explain]
  sshRun --help

Options:
  --hosts         List hosts
  --ssh           Ssh to the <host>, also on no options
  --scripts       List scripts for <host>
  --explain       Explain the <host> or <script>
  --help          Show this screen
```

## Configuration

### Hosts

From any `*.yml` file in `~/.config/ssh-run/`.

```yaml
includeSshConfigHosts: true # include hosts in ~/.ssh/config, without sshRun specific config
localhostIs: example # alias localhost to this host

# hosts with sshRun specific config
hosts:
  - ip: 192.168.0.1
    user: exampleUser
    name: example
    port: 22
    portTunnel: 1080
    checkRemote: true      
```

### Scripts

Scripts saved in `.config/ssh-run/scripts` will be associated with hosts according to subdirectory:

- `common` script is associated with all hosts
- `host/$hostName` script is associated with `$hostName`, also checked on the remote host if `Config.checkRemote` is `true`
- `shared/$sharedDir` script is associated with `$hostName` when `$sharedDir` is also found in `host/$hostName`

Filename suffixes control how the scripts are run:

- `$scriptName.sh` _with no flags_
- `$scriptName.pty.sh` _with -t_
- `$scriptName.x11.sh` _with -Y_
- `$scriptName.local.sh` _locally_


In a given script `$ip`, `$userName`, `$hostName`, `$port`, `$portTunnel` are replaced with values from the associated host.

The script is passed eval, so your shell will also [expand parameters](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html) using passed arguments and environment variables etc from the client.


Use `sshRun $hostName $scriptName --explain` to see what will be run.

Example scripts can be found [here](https://github.com/tom-power/ssh-run/tree/main/config/.config/ssh-run/scripts).

## Installation

### From sources

Install [golang](https://go.dev/), then to build.

```shell
git clone https://github.com/tom-power/ssh-run &&
cd ./ssh-run &&
sh/setup.sh &&
sh/build.sh
```
to copy scripts to your path.

```shell
cp sh/.install.env.example sh/.install.env &&
nano sh/.install.env && 
sh/install.sh
```

### From release

```shell
wget https://github.com/tom-power/ssh-run/releases/download/latest/release_linuxamd64.tar.gz &&
tar -xvf release_linuxamd64.tar.gz -C release &&
cp ./release/{sshRun,_sshRun} /your/path/
```

### Local requirements

All the commands `sshRun` generate use [ssh](https://www.openssh.com/), so it should be available locally, and you should be able to [access](https://www.ssh.com/academy/ssh/public-key-authentication) any hosts you configure.

### Zsh completions

- `release/completions/_sshRun` completions for hosts/scripts.
- `release/completions/settings.aliases.sshRun.zsh` adding aliases for your hosts.
- [fzf-tab](https://github.com/Aloxaf/fzf-tab) is awesome!

