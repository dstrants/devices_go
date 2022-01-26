# Devices

This is a simple to get notified about battery related events on iOS devices using the powerful shortcuts tool. More specifically when event like running out of battery, plugged in and out of charger, or when fully charged, a webhook is posted to this app and a notification is delivered to slack regarding this.


## Environment Configuration

Variable | Default | Description
---------|---------|-------------
MONGO_CONNECTION_STRING | **required** | The connection string to the mongo database. Used for caching requests.
MONGO_DATABASE | `devices` | The database to be used by the tool
SLACK_TOKEN | **required** | Slack application token to be used for the notification delivery
SLACK_CHANNEL | `devices` | The slack channel to deliver the notification. Make sure the app is a member fo that channel
API_KEY | **required** | Secret token used for authentication.

### iOS Shortcuts Configuration.

To be added soon