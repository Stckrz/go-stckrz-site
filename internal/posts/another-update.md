---
title:  "Another Dev Update"
date:   "2025-08-08"
slug:   "another-dev-update"
tag:    "dev-log"
summary: "Continued development for the site"
---

So today was a lot of work on the overall aesthetic of the site, and kind of setting up some groundwork for consisntent styling.. I like where some of the stuff is going, but not everything. I'm also working toward having a config file in a way that makes sense and isn't frustrating, as well as starting work on building some custom stuff for the site.

Today, the one I started working on, was a custom guest book. I still havent decided if I want the guestbook to be an internal database, or to make a separate microservice for the guestbook that I could maybe re-use later if I needed.. Also, I could do it internally with something like just md files or something, or have an actual DB. My heart is telling me it would be better to make a separate service, run it on the cluster, and just call that information overe using JS or maybe HTMX to populate.. Modularity, separation of concerns, etc.. 

I really need to get the kubernetes up, though, because once I start dockerizing this and having multiple services running, it's going to be more frustrating to develop, and the whole point of this site is for it to be kind of more laid back.
