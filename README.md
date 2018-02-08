[![Build Status](https://travis-ci.org/malud/temgo.svg?branch=master)](https://travis-ci.org/malud/temgo)
# temgo
Environment variable based template engine like Jinja2 for your container config files. Lightweight, fast, go cli

**Note:** wip

## Usage

Define variables with the format `{{ VAR_IABLE }}` inside your template files and set them in your environment before **tg** execution.

Currently with stdin/stdout and inline support.
 
 ```bash
# redirection
TESTVAR=foo  tg < /opt/templates/file1.ext > /dest/config/file1.ext
# or inline
TESTVAR=foo tg -i /dest/config/file1.ext
 ```
 
However, using the inline option on your templates will overwrite them.
It is recommended to use this option on resettable files.

Placeholders which have no corresponding environment variable gets not replaced.
`echo 'foo {{ NOT_SET }} bar' | ./tg` results in `foo {{ NOT_SET }} bar`

### TODO
* flags for file in/out
* prefix flag for env var e.g. SERVICE_XXX

### License

MIT
