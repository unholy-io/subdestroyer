# SUBDESTORYER

Need some subdomains? I got you. Need some subdirectories? I got you. `subdestroyer` by [unholy.io](www.unholy.io) collects subdomain and subdirectory data.

## Inspiration

This project was originally inspired by what I learned watching [Nahamsec's Twitch Stream](https://twitch.tv/nahamsec).

He has a few scripts available over at [his Github](https://github.com/nahamsec).

## Background
What makes this project different is that I wanted to explore the impact of adding enhancements such as recursion when querying crt.sh. Performance became a huge bottleneck when attempting to do this with Bash scripting. 

My original Bash prototype first made a call to crt.sh for %.playstation.net and then a new call for each subdomain the original request returned. The script took approximately 27 minutes. The code you'll find in the initial commit of this repository was written the same day and took about 15 seconds.

I'm also trying to use this as an opportunity to learn Go. I've dabbled for work but haven't made time to really dive in. I'm hoping this project, and in fact the rest of the code under [github.com/unholy-io](https://github.com/unholy-io), will be the catalyst I need to really learn the language.

## Current State of Development
This project is an infant right now but I hope to grow it. If you find bugs please open an issue. If you want to collab, hit me up.