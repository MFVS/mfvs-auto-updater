# MVFS docker-compose auto-updater

We need to have a zero-touch containerized deployment for a series of apps we make for our school that updates periodically. There's 2 ways we could do this:

## Cron

We can run a cronjob that stops the container, pulls and rebuilds.

`crontab -e`
`0 0 * * * /bin/bash -c "docker-compose down --volumes && cd /path/to/your/folder && git pull && docker-compose up --build --force-recreate"`
Replace for `* * * * *` for running it every minute.

## Auto-updater container

We can also have a container that has a volume bound to the local filesystem that periodically does a `git pull` and `docker-compose up --build --force-recreate` on all folders in the directory. This is more complicated, but easier to manage.
