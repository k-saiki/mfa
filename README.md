# mfa

Generate TOTP(Time-based One-time Password) token with CLI.

## Installation

- Download from [Release](https://github.com/k-saiki/mfa/releases/latest).
- Unzip and move to the directory of PATH.

## Configuration

Default config file is `$HOME/.mfa/secrets.yml`  or `$HOME/.mfa/secrets.yaml` .

```yaml
service:
  - name: amazon
    secret: "your secret key"
  - name: google
    secret: "your secret key"
  - name: github
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

[aws-vault](https://github.com/99designs/aws-vault) with [direnv](https://github.com/direnv/direnv).

```bash
$ echo 'export AWS_REGION=<region code>' > .envrc
$ echo 'export AWS_DEFAULT_REGION=${AWS_REGION}' >> .envrc
$ echo 'eval $(aws-vault export --mfa-token=$(mfa gen aws) --format=export-env <profile>)' >> .envrc
$ direnv allow
direnv: loading .envrc
direnv: export +AWS_ACCESS_KEY_ID +AWS_CREDENTIAL_EXPIRATION +AWS_DEFAULT_REGION +AWS_REGION +AWS_SECRET_ACCESS_KEY +AWS_SESSION_TOKEN
```
