cf-redirector
=============

`cf-redirector` is a small Cloud Foundry application for
redirecting traffic bound for one route over to another, via
standard HTTP 3xx redirection mechanisms.

Deployment
----------

Hey, it's Cloud Foundry!

    cf push point-a --no-start
    cf set-env point-a REDIRECT_TO other.domain.cf.tld
    cf set-env point-a STATUS_3XX 301
    cf start

Configuration
--------------

Configuration is done entirely through environment variables:

- `REDIRECT_TO` is the full hostname of the other CF application
  that you want to redirect your traffic _to_.  This value is
  **required**.
- `REDIRECT_SCHEME` is the URL scheme ('http' or 'https') to force
  redirection to.  By default, this is 'https'.
- `STATUS_3XX` is the HTTP status code to use when issuing
  redirect responses.  This value must be between 300 and 399.
  By default, HTTP 302's are sent.
- `DEBUG`, if set, will turn on debugging messages, to standard
  error, that can be viewed with `cf logs`.  This is useful for
  verifying that your configuration is correct, but is not
  advisable in production settings with high traffic.

To configure these, use `cf set-env` and then restage your app:

    cf set-env my-redirector REDIRECT_SCHEME=http
    cf restage my-redirector

