```bash
  _    _ _____  _     __   ________ ______ ______ 
 | |  | |  __ \| |    \ \ / /  ____|  ____|  ____|
 | |  | | |__) | |     \ V /| |__  | |__  | |__   
 | |  | |  _  /| |      > < |  __| |  __| |  __|  
 | |__| | | \ \| |____ / . \| |____| |    | |     
  \____/|_|  \_\______/_/ \_\______|_|    |_|     
```
Builds urls for directory and subdomain enumeration.

# Requirements
- Golang

# Installation

```bash
git clone <your link>
cd urlxeff
go build -o urlxeff main.go
```

# Usage

```bash
urlxeff
```
This shows the help message.

```bash
urlxeff -t target -wl wl.txt -m single-directory -n single-d
```
This builds all directory combinations (according with the wordlist) for a single target.

```bash
urlxeff -tw twl.txt -wl wl.txt -m multiple-directory -n multiple-d
```
This builds all directory combinations (according with the wordlist) for all targets on the list.

```bash
urlxeff -t target -wl wl.txt -m single-sub -n single-sub
```
This builds all subdomain combinations (according with the wordlist) for a single target.

```bash
urlxeff -tw twl.txt -wl wl.txt -m multiple-sub -n multiple-sub
```
This builds all subdomain combinations (according with the wordlist) for all targets on the list.

You can always change the -n flag for what you want...

If you use ```bash -p```, it don't put http:// as a prefix.
