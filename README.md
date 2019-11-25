[![CircleCI](https://circleci.com/gh/k-saiki/mfa.svg?style=svg)](https://circleci.com/gh/k-saiki/mfa)

# mfa
Generate TOTP(Time-based One-time Password) token with CLI.

## Installing
```bash
go get -u github.com/k-saiki/mfa
```

## Usage
### Configuration
Create config at `$HOME/.mfa/secrets` in YAML.
```yaml
service:
  - name: "amazon"
    secret: "your secret key"
  - name: "google"
    secret: "your secret key"
  - name: "github"
    secret: "your secret key"
```

### Generate token
```bash
$ mfa gen amazon
999999
```

### List services
```bash
$ mfa list
amazon
google
github
```

## Use case
[direnv](https://github.com/direnv/direnv) with [assume-role](https://github.com/remind101/assume-role)
```bash
$ echo 'eval $(mfa gen aws | assume-role ${PROFILE_NAME})' > .envrc
$ direnv allow
direnv: loading .envrc
MFA code:
direnv: export +ASSUMED_ROLE +AWS_ACCESS_KEY_ID +AWS_SECRET_ACCESS_KEY +AWS_SECURITY_TOKEN +AWS_SESSION_TOKEN
```
