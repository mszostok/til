# Email alias

If you have a domain on Cloudflare, there is a beta feature that allows you to set up [Email Routing](https://developers.cloudflare.com/email-routing/). However, Cloudflare provides email forwarding, not hosting. You cannot send your emails through Cloudflare.

While Cloudflare’s email routing service doesn't provide a way to send emails, you can use a service which provides SMTP login credentials. Next, you need to update the SPF record injected by Email Routing to add a given provider’s SPF. For example, for Google:

  > Create a custom TXT record for SPF: It should look like this: v=spf1 [IP ADDRESS] include:_spf.google.com include:[DOMAIN] ~all

As a result, your email will be trusted by the recipient and won't land in spam.

[Here](https://improvmx.com/guides/send-emails-using-gmail/) is a good tutorial on how to set up Gmail.

Important note is that this is a deprecated Google feature. At some point, they could disable this functionality.

Other services:
- https://www.mailgun.com/
- https://simplelogin.io/
