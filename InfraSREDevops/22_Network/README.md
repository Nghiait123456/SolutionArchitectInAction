- [All Of Network And Ten Minute](#all-of-network-in-ten-minute)
- [Facebook Smart Traffic And LB](#facebook-smart-traffic-and-lb)
- [How Facebook Disappeared From The Internet And Down](#how-facebook-disappeared-from-the-internet-and-down)
- [How To Reduce Network Delay In The Problem Of Users Lying Around The World](#how-to-reduce-network-delay-in-the-problem-of-users-lying-around-the-world)

## All Of Network And Ten Minute <a name="all-of-network-in-ten-minute"></a>

A panoramic, complete and detailed view of the network's components, how a computer in Vietnam connects to a computer in
the US. Here's a scripted talk with AI to create a relaxed, accessible atmosphere for network
newbies: https://www.youtube.com/watch?v=mPyqvuQ2GL0. <br>. The content of that conversation is also saved under the
video </br>

## Facebook Smart Traffic And LB <a name="facebook-smart-traffic-and-lb"></a>

Describe and explain how Facebook sets up smart traffic for its system: https://www.youtube.com/watch?v=mPyqvuQ2GL0.
Reference source: https://www.youtube.com/watch?v=akla5Ya6u2g </br>

## How Facebook Disappeared From The Internet And Down <a name="how-facebook-disappeared-from-the-internet-and-down"></a>

For a description of how a giant system like Facebook is completely down while the servers are still working,
link: https://www.youtube.com/watch?v=mPyqvuQ2GL0 . Reference
source: https://blog.cloudflare.com/october-2021-facebook-outage/ </br>

## How to reduce network delay in the problem of users lying around the world  <a name="how-to-reduce-network-delay-in-the-problem-of-users-lying-around-the-world"></a>

Here, I have a region in Vietnam but my user is located globally. The handshake time from the US user to the Vietnamese
server is about 600ms, I want to reduce this time. The two most common techniques: </br>

1) Add an intermediate proxy C, C located in the US and therefore, the handshake time between user vs C will be reduced,
   between C and the backend cluster in Vietnam set up TCP Fast Open. It will no longer take time to connect between the
   parties. </br>

2) Remove https between C and the backend cluster in Vietnam if the network has enough security factors. </br>

Best practices: </br>
If the user is globally but your system does not have the ability to invest in clustering in many regions, use smart
traffic. DNS will resolve users by region to LBs or proxies in that region, and that proxy will forward the request to
the backend cluster in Vietnam. Still based on the above prerequisites, you can reduce network latency in half. The
clouds all provide this service very well. </br>

A scripted conversation with AI will help anyone understand https
handshake:  https://github.com/Nghiait123456/InfraSREDevopsBackendForBigProject/blob/master/InfraSREDevops/22_Network/https_handshake.txt </br>