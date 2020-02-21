# Github Commit Reput

> Improve your github footprint by committing every day 

Github Commit Reput stream a specific keyword from Twitter, encrypt the tweets, and push them to your private repository.

The idea is to be as close to a normal commit as possible. Streaming Twitter allows to not follow any logical pattern.

## Requirements

- `go > 1.13`
- Create a [Twitter App](https://docs.inboundnow.com/guide/create-twitter-application/) and note the twitter consumer key/secret and twitter access token/secret.  
- Create a brand new private github repository and generate a [Deploy Key](https://developer.github.com/v3/guides/managing-deploy-keys/#deploy-keys). Using a deployment key instead of a token is more secure. Github Commit Reput will only be able to access this specific repository.

That's it.

## How to run locally

Clone the repository

```shell script
git clone git@github.com:bappr/github-commit-reput.git
```

Get the dependencies

```shell script
go get ./...
```

Copy the `.env.sample` file into `.env` and fill in the environment variables (see below).

Run the app

```shell script
go run cmd/main.go
```

## How to run it with docker

_Note: docker needs to be installed_

Build the image

```shell script
docker build -t github-commit-repo .
```

Run the image

```shell script
docker run --env $ENVIRONMENT_VARIABLES github-commit-repo
```

## Environment variables

| *Variable*                | *Description*                                            | *Default*          |
|:--------------------------|:---------------------------------------------------------|:-------------------|
| `LOG_LEVEL`               | Logging level                                            | `DEBUG`            |
| `TIMEOUT`                 | Time To Live - in seconds                                | `3600`             |
| `GIT_USERNAME`            | Github Username                                          |                    |
| `GIT_EMAIL`               | Github Email                                             |                    |
| `GIT_COMMIT_QUEUE_MIN`    | Minimum number of commits that will committed each round | `2`                |
| `GIT_COMMIT_QUEUE_MAX`    | Maximum number of commits that will committed each round | `5`                |
| `GIT_REPO`                | Name of the private github repository                    | `commit-reput-bot` |
| `GIT_DEPLOY_KEY`          | Github Deployment Key encoded in Base64                  |                    |
| `TWITTER_KEYWORD`         | Twitter keyword to follow                                | `COVID2019`        |
| `TWITTER_CONSUMER_KEY`    | Twitter Consumer Key                                     |                    |
| `TWITTER_CONSUMER_SECRET` | Twitter Consumer Secret                                  |                    |
| `TWITTER_ACCESS_TOKEN`    | Twitter Access Token                                     |                    |
| `TWITTER_ACCESS_SECRET`   | Twitter Access Secret                                    |                    |
| `REPO_PATH`               | Local path where the repo will be created                | `/tmp`             |




## Support

ben@codereput.com