- [Type of RateLimit](#type_of_rate_limit)
- [Best Practice Choose](#best_practice_chose)

## Type of RateLimit <a name="type_of_rate_limit"></a>

The simple algorithm rate limit is to count the number of hits per unit of time and control it within a set threshold.
There are two most common types of rate limits: rate limit in software application and rate limit in network. </br>

With RateLimit Software, the software running on your server will directly count and control the limit. With ratelimit
layer network, the service network will control the number of requests to your servcie, ensuring it is always within the
allowed threshold. </br>

## Best Practice Choose <a name="best_practice_chose"></a>

Rate limit in software application can work with low request threshold, but with high request threshold, all will be a
storm to your server. The amount of requests is so large that your system will be overloaded as soon as it only does the
job of rate limit. In this case, use a third party service like cloud
flare: https://developers.cloudflare.com/fundamentals/api/reference/limits/. It handles from DNS, network, and ensures
the total number of requests to the server is always within the safe threshold of that service. </br>

The disadvantages of rate limit software application are described in detail with
laravel: https://github.com/Nghiait123456/DissectLaravel#why_do_not_use_rate_limit_laravel_for_attack_ddos. </br>