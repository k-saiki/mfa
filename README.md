# mfa
Generate TOTP(Time-based One-time Password) token with CLI.

## Usage
### Configuration
Create config at `$HOME/.mfa/secrets`.
```toml
[[service]]
name = "amazon"
secret = "your secret key"

[[service]]
name = "google"
secret = "your secret key"

[[service]]
name = "github"
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
go get github.com/k-saiki/mfa
```