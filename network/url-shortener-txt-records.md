## URL shortener as TXT records

You need to take it with a grain of salt ;)

> What’s your favorite database?
>
> Mine is Route 53.
> <!--more-->
>
> It has a 100% SLA, you can query it, and DNS is fundamentally a large key-value store.
>
> Listen to @AWSMorningBrief here: https://t.co/6OlOVFakko
>
> — Corey Quinn (@QuinnyPig) November 13, 2019

> So how does this URL shortener work? Well like most it has a base URL and then a shortcode at the end which correspods to a URL to redirect to. We also have a base domain which may or may not be the same as the base URL, and the flow happens thusly:
>
> ```
>  HTTP request to https://[base URL]/[shortcode] (e.g. https://ols.wtf/_orgy)
>  DNS TXT record lookup against [shortcode].[base domain] (e.g. orgy.short.ols.wtf)
>  Contents of the TXT record are used to perform a 302 redirect (e.g. https://ols.wtf/2021/05/10/orgy-personal-tech-stack.html)
> ```
>
> You want to create a new redirect? Create a TXT record with the URL to redirect to on your dedicated URL shortening subdomain. With these being individual DNS records, you can chop and change the TTL on each of them as you wish, and now you can run as many instances of this as you want with their distributed yet authoritative record of your data. Fantastic!

_source: https://ols.wtf/2021/05/17/url-shortener.html_

---

> So, wait… this is actually a good idea?

> I think it’s safe to say that the proof of concept works pretty nicely, and it’s a pretty creative solution. But despite the glowing benefits as to why you might do this, there are a few drawbacks to consider:
>
> Write lag: Due to the caching behavior of DNS resolvers, it can take anywhere from 60s to 48hrs for DNS changes to propagate. Even though most public DNS resolvers will respect the 60s TTL, some override it or ignore it entirely.
>
> Size limitations: Strings in TXT records have a max length of 255 characters. The FQDN is also restricted to 255 characters.
>
> Row limits: Route53 has a limit of 10,000 resource record sets per hosted zone. You could request a higher limit, but good luck explaining the reason.

_source: https://betterprogramming.pub/apparently-you-can-use-route53-as-a-blazingly-fast-database-dd416b56b005_
