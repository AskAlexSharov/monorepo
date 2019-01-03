#!/bin/sh
echo "prepare environment"

chown -R app:app /srv/var 2>/dev/null

echo "running app"

/sbin/su-exec app /app $@
