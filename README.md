therealplato/shortlink
======================

This is a shortlink host.

Configure it with environments:
HEALTHCHECK_LISTEN_ADDR
SHORTLINK_LISTEN_ADDR
BASE_URL (this is the prefix to all short links, with a trailing slash)
POSTGRES_URI

Setup
-----
Getting Postgres into the correct state and keeping it there is out of scope of this project.
The most basic manual setup:
```
psql -hmydb.foo -Umyuser -c "CREATE DATABASE shortlink" template1
psql -hmydb.foo -Umyuser -f $GOPATH/src/therealplato/shortlink/schema.sql shortlink
```

Usage
-----
Request BASE_URL. You'll get a landing page with a form where you submit the link you want to shorten. After submission, you'll land on
BASE_URL/preview/slug, displaying both the long and short links.

Request BASE_URL/slug. If slug is a shortened link, you'll be redirected to the long link.

Advanced Usage
--------------
Request BASE_URL/https://imgflip.com. A shortlink will be looked up, or created, for https://imgflip.com. You'll land on a
BASE_URL/preview/slug, displaying both the long and short links.
