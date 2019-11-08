# Build
docker build -t fafeliscord .

# Run
docker run --name fafeliscord -h fafeliscord --network fafelnet -e DISCORD_TOKEN="xxx" -d fafeliscord

# Authorize Fafeliscord Bot :
https://discordapp.com/api/oauth2/authorize?client_id=642346661716295691&scope=bot&permissions=67584