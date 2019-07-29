[![CircleCI](https://circleci.com/gh/k-saiki/mfa.svg?style=svg)](https://circleci.com/gh/k-saiki/mfa)

# mfa
Generate TOTP(Time-based One-time Password) token with CLI.

## Usage
### Configuration
Create config at `$HOME/.mfa/secrets` in YAML.
```yaml
service:
  - name = "amazon"
    secret = "your secret key"
  - name = "google"
    secret = "your secret key"
  - name = "github"
    secret = "your secret key"
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

## Install
```bash
go get -u github.com/k-saiki/mfa
```