# mfa
Generate TOTP(Time-based One-time Password) token with CLI.

## Usage
### Configuration
Create config at `$HOME/.mfa/secrets`.
```toml
[[service]]
Name = "amazon"
secret = "your secret key"

[[service]]
Name = "google"
secret = "your secret key"

[[service]]
Name = "github"
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