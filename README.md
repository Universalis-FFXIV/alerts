# alerts

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9ab1040a9f9c43faaed6cc46925db82d)](https://app.codacy.com/gh/Universalis-FFXIV/alerts?utm_source=github.com&utm_medium=referral&utm_content=Universalis-FFXIV/alerts&utm_campaign=Badge_Grade_Settings)

Alerts service for mogboard.

## Notification methods
*   Email
*   Discord

## Setup
Set the environment variable `UNIVERSALIS_MAILGUN_KEY` to your Mailgun API key, and the environment variable `UNIVERSALIS_ALERTS_DISCORD_TOKEN` to your Discord bot token.

## Endpoints

### POST `/discord/send`
Sends a notification over Discord.

### POST `/email/send`
Sends a notification over email.

#### Notification payload
```
{
    "targetUser": string,
    "notification": {
        "itemName": string,
        "pageUrl": string,
        "body": string
    }
}
```
