# Notify Telegram

Get workflow status notifications to Telegram chat or channel.

## Usage

First of all, you need to create a Telegram bot by talking to (@BotFather)[https://t.me/botfather] bot. See official guide here: https://core.telegram.org/bots#6-botfather

If you want to get notifications to personal chat with bot, find your user id, for example by talking to (@jsondumpbot)[https://t.me/jsondumpbot].

Also you can use channel for notifications, in this case just get your channel name in format `@channelname`.

Then add your bot token and user id (or channel name) to repository Secrets.

Add following step to the end of your workflow:

```yaml
    - uses: yanzay/notify-telegram@v0.1.0
      if: always()
      with:
        chat: ${{ secrets.chat }} # user id or channel name secret
        token: ${{ secrets.token }} # token secret
        status: ${{ job.status }} # do not modify this line
```
