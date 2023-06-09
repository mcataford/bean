<p align="center"><img src="./assets/bean.png" alt="Bean!"/></p>

# bean
Fancy Discord bot for personal use


## Setup

A few environment variables should be set through `.env` or directly in the `docker-compose.yml` service definition:

|Variable|Description|
|---|---|
|`DISCORD_APP_TOKEN`|Discord application auth token|
|`DISCORD_APP_ID`|Application ID from the bot configuration|
|`DISCORD_APP_PUBLIC_KEY`|Public key assigned to the application for signature verification|

### Local development

[Ngrok](https://ngrok.com/docs/getting-started/) can be used to expose the application port (`8080`) to the web so that it can be used as a webhook endpoint by
Discord.

### System stats reporting

The `beancounter` executable should be executed on the system that the bot will report on to expose summary statistics.

It can be build via `go build cmd/beancounter`.

At runtime, the following guidance on arguments is printed if not satisfied:

```
beancounter <outputPath> <intervalSeconds>

This tool will collect system statistics and create/update a report at <outputPath> every <intervalSeconds>.
```

## Commands

The bot installs a number of [slash commands](https://discord.com/developers/docs/interactions/application-commands) on the Discord server it's installed on. Available commands can be found
under `internal/commands`.

|Command|Description|
|---|---|
|`/ruok`|Bot healthcheck - if healthy, the bot responds to the ping.|

## Contributing

Outside contributions or issue reports aren't a particular focus right now as this is main focused on my own use
case. Feel free to fork this or take inspiration from it if you want to put together your own bots! If you do do that,
feel free to include a backlink here. :)

## Additional readings

- [Discord developer docs](https://discord.com/developers/docs)
- [Ngrok](https://ngrok.com/)
