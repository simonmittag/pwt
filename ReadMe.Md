![](pwt.png)

`pwt` ("port wait") is a commandline utility that blocks until a tcp socket, aka port becomes available. Designed to 
replace the bash sleep command in your CI/CD scripts, this process will exit with `0` once the socket connects, or `-1`
in the event of a timeout. 

[![Circleci Builds](https://circleci.com/gh/simonmittag/pwt.svg?style=shield)](https://circleci.com/gh/simonmittag/pwt)
[![Github Workflows](https://github.com/simonmittag/mse6/workflows/Go/badge.svg)](https://github.com/simonmittag/pwt/actions)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/pwt)](https://github.com/simonmittag/pwt/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/pwt)](https://img.shields.io/github/commit-activity/m/simonmittag/pwt)  
[![CodeClimate Maintainability](https://api.codeclimate.com/v1/badges/70cd59e4dfd2801f8661/maintainability)](https://codeclimate.com/github/simonmittag/pwt/maintainability)
[![CodeClimate Test Coverage](https://api.codeclimate.com/v1/badges/70cd59e4dfd2801f8661/test_coverage)](https://codeclimate.com/github/simonmittag/pwt/test_coverage)
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/pwt)](https://img.shields.io/github/go-mod/go-version/simonmittag/pwt)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.7-orange)](https://github.com/simonmittag/pwt/releases/tag/v0.1.6)

## What's New
### v0.1.7
* -h -v flags for cli
* bumped to go 1.20

## Up and running
### Homebrew
```
brew tap simonmittag/cli &&
  brew install pwt &&
  pwt 
```

### Golang
```bash
git clone https://github.com/simonmittag/pwt && cd pwt && 
go install github.com/simonmittag/pwt/cmd/pwt && 
pwt 
```

## Usage
```
λ pwt[v0.1.7]
Usage: pwt [-v]|[-w n] host[:port]
  -h    print usage instructions
  -v    print pwt version
  -w int
        time wait in seconds (default 10)
```

## Examples

Wait for ipv4 localhost on default port
```
λ pwt 127.0.0.1
  pwt waiting for 127.0.0.1:80 .......... aborted after 10 seconds  
```

Wait for ipv6 localhost on port 8080 for 3 seconds
```
λ pwt -w 3 [::1]:8080
  pwt waiting for [::1]:8080 .. connected in 2 seconds
```



Wait for google on port 443
```
λ pwt -w 60 google.com:443
  pwt waiting for google.com:443 connected in <1 second  
```



## Contributions
The pwt team welcomes all [contributors](https://github.com/simonmittag/pwt/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/pwt/blob/master/CODE_OF_CONDUCT.md)
