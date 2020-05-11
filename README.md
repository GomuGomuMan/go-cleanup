# go-cleanup

## What it is
An utility to clean up files/directories not used after certain time

## Why
Because I'm too lazy to clean up myself

## A little more details
* Default TTL is 30 days

## Install
```bash
go get -u github.com/GomuGomuMan/go-cleanup
```

## Example
```bash
go-cleanup -p ~/Downloads
```

If you want to set up to run in a schedule, simply register it with `crontab`.
My setup is every system start:
```bash
@daily go-cleanup -p ~/Downloads
```