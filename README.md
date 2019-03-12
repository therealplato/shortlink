therealplato/shortlink
======================

This is a shortlink host.

Configure it with environments:
HEALTHCHECK_LISTEN_ADDR
SHORTLINK_LISTEN_ADDR
BASE_URL (this is the prefix to all short links, with a trailing slash)
POSTGRES_URI

Usage
-----
Request BASE_URL. You'll get a landing page with a form where you submit the link you want to shorten. After submission, you'll land on a
page displaying both the long and short links.

Request BASE_URL/slug. If slug is a shortened link, you'll be redirected to the long link.

Advanced Usage
--------------
Request BASE_URL/https://imgflip.com. A shortlink will be looked up, or created, for https://imgflip.com. You'll land on a page displaying
both the long and short links.
