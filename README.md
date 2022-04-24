# times

## Description
times is incredibly CLI that enables us to post message and see conversation history via terminal.

## Requirement
the times cli needs the authorization to execute some sub commands.
specifically, it needs `SLACK_API_TOKEN` and `TIMES`.

1. you'll create a Slack app that's just a container for your credentials and where put all the vital information about what your app is or does.
you cant' get a token without one.
[here](https://api.slack.com/apps) is the link

2. you'll set up permission scopes that gives you permissson to do things(for example, post messages)
you can select the scopes to add to your app by heading over to the OAuth & Permissions side bar.
Scroll down to the Scopes section and click to Add an OAuth Scope.
And one thing I want you to remember is that you'll have to add scope to `your User token` not `Bot token`
here's the scopes you'll have to request to use this CLI.
    - channels:history
    - channels:read
    - chat:write

3. now you must be done requesting scopes you need.
Next you'll have to install your App to workspace.
After clicking through one more green Install App to Workspace button, you'll get a token that enable us to executes subcommand on this CLI.

4. export below environment values by the use of `SLACK_API_TOKEN` and `TIMES`

```
export SLACK_API_TOKEN=${api token you got}
export TIMES=${channle id that you wanna post message to}
```
