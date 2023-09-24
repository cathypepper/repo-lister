# repo-lister

List all my repos for easy cloning.

# Usage

```bash
for rep in `./repo-lister | tr -d '"'`; do git clone git@github.com:cathypepper/$rep; done
```
