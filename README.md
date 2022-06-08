# Marmota

## Introduction
Marmota is an open-source software to help people resolve domain names to their private IP.
Since the private IP address for person modifications frequently, once a domain name has been resolved for a while, the IP would change to another one, causing the previous resolution to expire. This program can help you by automatically updating the latest IP to your domain name after your resolution has failed. In this way, all you need to do is to keep your domain name in mind so that you can  access your home network or any other computer equiped this program.

## Mechanism
Marmota will access the echo service which shows your IP back with a fixed frequency, in this way you can know the latest IP. Once your IP is found to have changed by comparing it with the local record, it will use the ability provided by the appropriate provider according to your configuration to update your domain name

## notice
One disadvantage of this mechanism is that the service that returns your IP must be available. So this project introduces the mechanism of group verification. It is important to note that once you use the software, it indicates that you are willing to provide the IP echo service to other people who also use this software. So the software the more people use, the more services that can echo IP back, and the ability of this service will be more and more reliable.

## thanks
Generally， automatic domain name resolution is depending on the capabilities offered by your domain name provider. We have already made adaptations for some operators, and the more adaptations of the left operators will coming soon.

