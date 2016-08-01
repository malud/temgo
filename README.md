[![Build Status](https://travis-ci.org/malud/temgo.svg?branch=master)](https://travis-ci.org/malud/temgo)
# temgo
Environment variable based template engine like Jinja2 for your container config files. Lightweight, fast, go cli

**Note:** wip

## Usage

Define variables with the format `{{ VAR_IABLE }}` and set them in your environment before **tg** execution.

Currently with stdin/stdout support.
 ```
TESTVAR=foo cat /opt/templates/file1.ext | tg > /dest/config/file1.ext
 ```

### TODO
* flags for file in/out
* prefix flag for env var e.g. SERVICE_XXX

### Licence

MIT
