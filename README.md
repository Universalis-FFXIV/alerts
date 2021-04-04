# alerts
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
