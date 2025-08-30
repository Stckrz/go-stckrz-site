---
title:  "Traefik and Metallb"
date:   "2025-08-30"
slug:   "traefik-and-metallb"
tag:    "dev-update"
summary: "An update about the job interview situation, and Kubernetes status."
---

I had my final interview, and I'm waiting to hear back.. I think it went more.. bad than it went good haha. I am not really holding my breath, but hey, maybe they saw something that I didn't. To be honest, I thought I was doing okay in the first bit, but I feel like I got fried on the hackerrank section, despite my ceaseless studying.. Either way, I really felt at odds, like maybe they sort of expected me to not do great from the beginning or something? It's possible I still get offered the position, but like I said, I'm not really holding my breath.

In other news, since that is sort of out of the way for now, it means that I can get back to focusing on all of those things that I have been putting off! This means learning Go, playing with Blender, building THIS website, and most of all: the cluster! So, finally yesterday, I installed Metallb as a load balancer and traefik for ingress: a pretty standar combo from what I understand. Thing is, i was not exactly able to get access to the services.. I ended up having to go to bed, but it seems like it is either on the incorrect NIC(unlikely), or there is some other issue with the ARP or something.. I was able to ping the internal ports, but not the external ip's, so that will still take some time to sort out.. Once I get that, though, it'll be straight to some kind of workflow for deploying and updating THIS site, as well as moving things over from the pi homelab to Kubernetes. So, job stuff be what it may, at least we are back on track for THIS stuff :) 
