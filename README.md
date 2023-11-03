# MVFS docker-compose auto-updater

## Why

We need to have a zero-touch containerized deployment for a series of apps we make for our school that updates periodically. There's 2 ways we could do this:

### Cron

We can run a cronjob that stops the container, pulls and rebuilds.

`crontab -e`
`0 0 * * * /bin/bash -c "docker-compose down --volumes && cd /path/to/your/folder && git pull && docker-compose up --build --force-recreate"`
Replace for `* * * * *` for running it every minute.

### Auto-updater container

We can also have a container that has a volume bound to the local filesystem that periodically does a `git pull` and `docker-compose up --build --force-recreate` on all folders in the directory. This is more complicated, but easier to manage.

## Current progress

CIT UJEP has created one VM for the purposes of hosting these apps and @nexovec has been given vpn and ssh accesses to this machine. We are currently waiting for a subdomain to be assigned for us(`stagapps.ki.ujep.cz`).
