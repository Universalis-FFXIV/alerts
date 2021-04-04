# alerts

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9ab1040a9f9c43faaed6cc46925db82d)](https://app.codacy.com/gh/Universalis-FFXIV/alerts?utm_source=github.com&utm_medium=referral&utm_content=Universalis-FFXIV/alerts&utm_campaign=Badge_Grade_Settings)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Universalis-FFXIV/alerts/Go?label=build%20%26%20test)

Alerts service for mogboard.

## Notification methods
*   Email
*   Discord

## Setup

### Environment variables
*   `UNIVERSALIS_MAILGUN_KEY` - Your Mailgun API key
*   `UNIVERSALIS_ALERTS_DISCORD_TOKEN` - Your Discord bot token
*   `UNIVERSALIS_ALERTS_EMAIL_ADDRESS` - Your notifications email address

## Endpoints

### POST `/discord/send`
Sends a notification over Discord.

### POST `/email/send`
Sends a notification over email.

#### Notification payload
```ts
{
    targetUser: string,
    notification: {
        itemName: string,
        pageUrl: string,
        reasons: string[]
    }
}
```

## Testing environment variables
*   `UNIALERTS_TEST_DISCORD_ID` - The Discord account to which the test notification will be sent.