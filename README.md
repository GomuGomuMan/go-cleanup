# go-cleanup

## What it is
An utility to clean up files not used after certain time

## Why
Because I'm too lazy to clean up myself

## Example
```bash
go-cleanup -p ~/Downloads
```

If you want to set up to run in a schedule, simply register it with `crontab`.
My setup is every system start:
```bash
@reboot go-cleanup -p ~/Downloads
```