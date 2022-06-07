# Marmota

## Introduction
Marmota is an open source software to help people resolve domain names with his private ip.
Since private ip's change frequently, once a domain name has been resolved for a while, the ip will change, causing the previous resolution to expired. This progran can help you by automatically updating the latest ip to your domain name after your resolution has failed.

## Mechanism
Marmota will access the service which can show back your ip with a fixed frequency, so you can know the latest ip when it has been invalidated for some time. Once your ip is found to have changed, it will use the ability provided by the appropriate provider according to your configuration to update your domain name

## notice
One disadvantage of this mechanism is that the service that returns your ip must be available. So this project introduces the mechanism of group verification. It is important to note that when you use the software, it indicate that you are willing to provide the ip replay service to other people who also use the software. So the software the more people use, the more services that can provide ip back, and the ability of ip back will become more and more reliable.

## thanks
Generally， automatic domain name resolution is depending on the capabilities offered by your domain name provider. We have already done adaptations for some operators, and we will gradually do so for more operators.

