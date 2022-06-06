# Marmota

## Introduction
Marmota is an open source software to help people resolve domain names to his own private ip.
Since private ip's change frequently, once a domain name has been resolved for a while, the ip will change, causing the previous resolution to fail. This software can help you by automatically updating the latest ip to your domain name after your resolution has failed.

## Mechanism
Marmota will access the service that can show back your ip with a fixed frequency, so you can know the latest ip when it has been invalidated for some time.

## notice
One disadvantage of this mechanism is that the service that returns your ip must be available. So this project introduces the mechanism of group verification. It is important to note that when you use the software, you indicate that you are willing to provide the ip replay service to other people who use the software. When more people use the software online, then the more services that can provide ip back, so the ability of ip back will become more and more reliable.

## thanks
Automatic domain name resolution is often required depending on the service capabilities offered by your domain name provider. We have already done adaptations for some carriers, and we will gradually do so for more carriers.

