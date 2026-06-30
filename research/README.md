# PackageCloud `/install` reference captures

Example responses from PackageCloud's per-repository install surface
(`https://packagecloud.io/install/repositories/<user>/<repo>/…`), kept for
posterity / as a reference if we ever reimplement package-manager setup in the
CLI instead of piping these scripts.

All token-shaped hex (read tokens, and the master token embedded in the
scripts' install URLs) is replaced with `REDACTED_TOKEN`.

These were captured from a single repo with representative `os`/`dist` values;
the repo name is replaced with `USER`/`REPOSITORY`.

## Setup scripts — `GET script.<type>.sh` (no query params)

Piped to a shell (`sudo bash` for the OS managers, `bash` for the language
ones). Each detects/derives `os`/`dist`/`unique_id` and configures the local
package manager (see each file).

| file | manager | calls |
|---|---|---|
| `script.deb.sh` | apt | `gpg_key_url.list`, `config_file.list`, `apt_auth_conf` |
| `script.rpm.sh` | yum/zypper | `config_file.repo` |
| `script.alpine.sh` | apk | `rsa_key_url.alpine`, `config_file.alpine` |
| `script.node.sh` | npm | `tokens.text` |
| `script.gem.sh` | rubygems | `tokens.text` |
| `script.python.sh` | pip | `tokens.text` |
| `script.helm.sh` | helm | `tokens.text` |

## Config / auth / token files

| file | endpoint | params used |
|---|---|---|
| `config_file.list` | `GET config_file.list` | `os=ubuntu&dist=jammy&source=script` (apt sources; token-less + `signed-by=`) |
| `config_file.repo` | `GET config_file.repo` | `os=el&dist=9&source=script` (yum/zypper) |
| `config_file.alpine` | `GET config_file.alpine` | `os=alpine&dist=v3.18&source=script` (apk repositories line) |
| `apt_auth_conf` | `GET apt_auth_conf` | `os=ubuntu&dist=jammy` (`/etc/apt/auth.conf.d/` line) |
| `gpg_key_url.list` | `GET gpg_key_url.list` | `os=ubuntu&dist=jammy` (URL to the deb/rpm GPG signing key) |
| `rsa_key_url.alpine` | `GET rsa_key_url.alpine` | `os=alpine&dist=v3.18` (URL to the apk RSA signing key) |
| `tokens.text` | `POST tokens.text` | `name=<id>` (mints/returns a read token, idempotent by name) |

All are authenticated with the repository's **master token** value as the
basic-auth username (the `default` master token is what the install page uses).
