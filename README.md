# shrimporter

A very fragile CLI tool built on an extremely specific dataset, used to import messages exported by [Tyrrrz/DiscordChatExporter](https://github.com/Tyrrrz/DiscordChatExporter). Nothing of usefulness lives within these halls, there is no God here.

# Using

0. Create a webhook in the Discord channel of your choice; copy its URL
1. Build the binary or download a release
2. Open a terminal and navigate to shrimporter's directory
3. Run

    - Linux: `./shrimporter -w [webhook_url] -f [exported_messages.json]`
    - Windows: `shrimporter.exe -w [webhook_url] -f [exported_messages.json]`

# License
MIT