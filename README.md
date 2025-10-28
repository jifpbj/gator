# gator

Boot.dev RSS aggreGATOR.

gator is a CLI tool that will follow RSS feeds automatically and aggregate them on your own postgresql database. You can then browse through the RSS feeds within your own terminal!

## Dependencies

You'll need postgresql and go installed to run the program.

```bash
brew install go
```

```bash
brew install postgresql
```

## Installation

```bash
go install "https://github.com/jifpbj/gator"
```

## Configuration

gator stores RSS posts based on users that are logged in. As this is a locally hosted sql server, there's no need to have a password. But use these commands to separate your feed from someone else if using the same computer.

register- register a new user

login - login an existing user.

reset - delete all users and feeds

users - show all users

## Commands

addfeed "Name" "URL"- add a feed and follow it

feeds - list all feeds

follow "URL"- follow an existing feed

following - show all feeds logged in user is following

unfollow "ID" - unfollow a feed

---

agg "time period"- aggregate feeds every 1 min, 30 min, hour, etc. This will run forever until you cancel using ctrl-C

browse "limit" - browse the posts aggregated, with a limit specified (defaults to showing 2)
