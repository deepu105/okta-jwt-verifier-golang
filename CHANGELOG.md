# Changelog

## v1.2.1 (February 16, 2022)

### Updates

* Update JWX package. Thanks, [@thomassampson](https://github.com/thomassampson)!

## v1.2.0 (February 16, 2022)

### Updates

* Customizable resource cache. Thanks, [@tschaub](https://github.com/tschaub)!


## v1.1.3

### Updates

- Fixed edge cause with `aud` claim that would not find Auth0 being JWTs valid. Thanks [@awrenn](https://github.com/awrenn)!
- Updated readme with testing notes.
- Ran `gofumpt` on code for clean up.

## v1.1.2

### Updates

- Only `alg` and `kid` claims in a JWT header are considered during verification.

