# mfa
Generate TOTP(Time-based One-time Password) token with CLI.

## Installation
```bash
$ go get -u github.com/k-saiki/mfa
```

## Configuration
Default config is `$HOME/.mfa/secrets` in yaml format.
```yaml
service:
  - name: "amazon"
    secret: "your secret key"
  - name: "google"
    secret: "your secret key"
  - name: "github"
    secret: "your secret key"
```

You can change config file path to use environment variable.
```bash
$ export MFA_CONFIG=/path/to/file
```

## Usage
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
[assume-role](https://github.com/remind101/assume-role) with [direnv](https://github.com/direnv/direnv).
```bash
$ echo 'eval $(mfa gen aws | assume-role <profile name>)' > .envrc
$ direnv allow
direnv: loading .envrc
MFA code:
direnv: export +ASSUMED_ROLE +AWS_ACCESS_KEY_ID +AWS_SECRET_ACCESS_KEY +AWS_SECURITY_TOKEN +AWS_SESSION_TOKEN
```

## TODO
- Add test
