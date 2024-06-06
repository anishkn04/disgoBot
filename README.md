**This CLI application**, fetches data from publicly available website or a file in the same directory. It regularly checks for new events and posts updates when they are available. These updates are posted on both Discord and Facebook. The bot ensures that community members are always informed about the latest events. It streamlines event notifications across multiple platforms.

---

**The first initialization** will ask you to provide `pageId`, `botToken`, `page_access_token`, and `channel_id`.

---

### To set up your bot for Discord and Facebook, follow these steps:

**For Discord**

1. **Create a Bot in Discord Developer Portal:**
    - Go to the [Discord Developer Portal](https://discord.com/developers/applications).
    - Click on "New Application" and give your application a name.
    - Navigate to the "Bot" section and click "Add Bot".
    - Under the "TOKEN" section, click "Copy" to copy your bot token. *Keep this token secure and do not share it publicly*.

2. **Add the Bot to Your Server:**
    - Go to the "OAuth2" section.
    - In "OAuth2 URL Generator", under "SCOPES", select `bot`.
    - In "BOT PERMISSIONS", select the permissions your bot needs.
    - Copy the generated URL and open it in a new browser tab.
    - Select the server you want to add the bot to and authorize the bot.

3. **For Channel Id**

    To obtain the channel ID for Discord, follow these steps:

    1. **Open Discord:** Launch the Discord application or visit the Discord website and log in to your account.
    2. **Navigate to Server Settings:** Go to the server where you want your bot to post messages.
    3. **Enable Developer Mode:** If you haven't already, enable Developer Mode in Discord. Go to User Settings > Appearance > Advanced, then toggle on Developer Mode.
    4. **Right-click the Channel:** In the channel where you want your bot to post messages, right-click the channel name.
    5. **Copy Channel ID:** From the context menu, select "Copy ID". This will copy the unique channel ID to your clipboard.

**For Facebook**

1. **Get Your Page ID and Page Access Token:**
    - Follow the video tutorial provided: [YouTube Guide](https://www.youtube.com/watch?v=s8c2SMpWDOo).
    - Go to the [Facebook for Developers](https://developers.facebook.com/).
    - Create a new app or select an existing app.
    - Navigate to the "Graph API Explorer".
    - Get a User Access Token, then use it to get a Page Access Token.
    - Copy the Page ID and Page Access Token.

---

By following these steps, you should be able to set up and run your bot on both Discord and Facebook platforms. If you have any specific code issues or need further guidance, feel free to ask!

Additionally, ensure that the events on the website are formatted as follows for fetching:

```json
{
    "events": [
        {
            "banner": "",
            "title": "",
            "start_date": "",
            "end_date": "",
            "location": "",
            "description": "",
            "link": ""
        },
        {
            ...
        },
    ]
}
