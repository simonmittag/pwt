![](pwt.png)

pwt ("port wait") waits for you until a tcp socket, aka port becomes available and exits cleanly once it does. Designed for 
continuous integration scripts, never call sleep inside your pipeline again.

[![Circleci Builds](https://circleci.com/gh/simonmittag/pwt.svg?style=shield)](https://circleci.com/gh/simonmittag/pwt)
[![Github Workflows](https://github.com/simonmittag/mse6/workflows/Go/badge.svg)](https://github.com/simonmittag/pwt/actions)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/pwt)](https://github.com/simonmittag/pwt/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/pwt)](https://img.shields.io/github/commit-activity/m/simonmittag/pwt)  
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/pwt)](https://img.shields.io/github/go-mod/go-version/simonmittag/pwt)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Up and running
### Osx
```
brew tap simonmittag/cli &&
brew install pwt &&
pwt -t localhost -p 80
```

### Golang
```bash
go get github.com/simonmittag/pwt && 
go install github.com/simonmittag/pwt/cmd/pwt && 
pwt -t localhost -p 80
```

## Usage
```
λ pwt -h
  Usage of pwt:
    -p uint
          the destination port (default 80)
    -t string
          the destination host name (default "localhost")
    -w int
          time wait in seconds (default 10)
```

## Contributions
The pwt team welcomes all [contributors](https://github.com/simonmittag/pwt/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/pwt/blob/master/CODE_OF_CONDUCT.md)
