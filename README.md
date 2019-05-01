# Pmdr
A simple pomodoro CLI timer writed in go

## Install
Install Cobra and beep libs for go (https://github.com/spf13/cobra, https://github.com/faiface/beep)
Make sure gopath is set!
In repo folder run: 
```sh
$ go install
```
Now the binaries are in your Go/bin folder.
_____________________________________________________________________________________________________________
## Usage

### Launch Work Timer
```sh
pmdr work
```
#### Flags
- -t or --time ==> set the time it will run, default is 25min
- -s or --sound ==> set an alarm sound/song to play when time finish

### Launch Break Timer
```sh
pmdr break
```
#### Flags
- -t or --time ==> set the time it will run, default is 25min
- -s or --sound ==> set an alarm sound/song to play when time finish

> The sound path can only be absolute. If needed u can use a config file.
