# gator

A simple CLI RSS feed reader.

Requires Postgres and Go.

Install using:

go install https://github.com/IgorP25/gator

Then create a .gatorconfig.json file in your home directory with the following contents:

{
        "db_url": "postgres://username:password@dbhost:dbport/gator?sslmode=disable",
}


Usage: gator <command> [args...]

login <name> - Set current user.
register <name> - Register new user.
reset - Reset database.
users - List users.
agg <time_between_requests> - Run continuous scan of feeds at intervals of time_between_requests.
addfeed <url> - Add RSS feed to database by URL.
feeds - List all feeds in database.
follow <url> - Follow feed at URL (must already be added to database).
following - List feeds current user is following.
unfollow <url> - Unfollow feed at URL.
browse [limit] - Browse posts from followed feeds. Optional limit number of posts returned. (Default: 2)
