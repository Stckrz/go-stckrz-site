---
title:  "Guestbook Update"
date:   "2025-08-12"
slug:   "guestbook-update"
tag:    "regular-posts"
summary: "Another update on the guesbook functionality"
---
For a couple days there, I was sort of focused on work, and this potential job application, but that may have just fell through.. oh well. Anyways, I have begun working on the guestbook stuff again, and basically have an api set up to handle the guestbook entries. It's not complicated, but it has validation for created entries, as well as the ability to get all of the guestbook entries. Should be more than good for what I'm doing with it.. Now, I just have to hook everythig up to the front-end and wala! Of course, it will have to run within the Kubernetes cluster, that way this api can just access it internally, and I wont even have to worry about exposing it to the internet or anything. 

Still, some of this stuff seems too simple.. I need examples of more complex api's in go in order to understand more complex work with the language. I still have to figure out images too, but a lot of this is waiting on the Kube to be up.. and it WILL be up this week, come hell or high water.
