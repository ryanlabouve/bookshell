# Bookshell

Proof of concept to quickly deploy a static site as an API.

Currently powering https://security-grimoire.vercel.app/ / https://github.com/ryanlabouve/security-grimoire


* On initialization clones repo
* Loads posts that match a pattern to memdb
* Serves
* Webhook to update git repo
* TODO: re populate db when webhook runs