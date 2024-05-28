The events.json should be in this particular format:
<br>
```
{
    "events": [
        {
            "banner": "Image location (Must be in JPG format)",
            "title": "____",
            "start_date": "____",
            "end_date": "____",
            "location": "____",
            "description": "____",
            "link": "____"
        },
        ...
    ]
}
```

**Copy .env.example to a new .env file and change the data as needed**

**Most of the problems listed below are solvable and will be solved soon!** <br><br>
- If the server is restarted, the events will be sent from the very beginning because the events' title are stored in a global variable and not compared to the messages present in the channel.<br>
    - Due to the limit of reading just 100 messages from a channel, if the events exceed that, any event before the last 100 events will be re-sent as a message.<br>
    - After the message is sent, it can be read as a `Message` only but not as a `MessageEmbed` due to which comparison will require creating a different method to parse the data.