---
title: SIDUPASTE.CONF
header: Sidupaste Manual
footer: 1.0.1
date: Jan 20, 2022
section: 1
---
# NAME
**sidupaste.conf** – sidupaste CLI configuration file.

# DESCRIPTION
The sidupaste(1) command line interface create paste to an PrivateBin
instance configured in the **config.json**.

# FORMAT
## Top level object keys:
**open_discussion** *bool* (default: false)
: The default value of open discussion for a paste.

**burn_after_reading** *bool* (default: false)
: The default value of burn after reading for a paste.

**formatter** *string* (default: "plaintext")
: The default formatter for a paste.

**expire** *string* (default: "1day")
: The default time to live for a paste.

**bin** *array\<bin\>*
: The list of bin instances.

## The bin object format:
**name** *string*
: The name of the bin instance.

**host** *string*
: The url of the bin instance.

**auth** *auth*
: The basic auth configuration of the bin instance.

**expire** *string*
: The default time to live for a paste.

**open_discussion** *bool*
: The default value of open discussion for a paste.

**burn_after_reading** *bool*
: The default value of burn after reading for a paste.

**formatter** *string*
: The formatter for the paste.

## The auth object format:
**username** *string*
: The basic auth username.

**password** *string*
: The basic auth password.

# EXAMPLES
Minimal sidupaste configuration file:

    {
        "bin": [
            {
                "name": "", // default
                "host": "https://paste.siduction.org"
            }
        ]
    }

A bit more complete configuration file:

    {
        "bin": [
            {
                "name": "example",
                "host": "bin.example.com",
                "auth": {
                    "username": "john.doe",
                    "password": "s$cr$t"
                },
                "formatter": "markdown",
                "burn_after_reading": false
            },
            {
                "name": "",
                "host": "https://privatebin.net"
            },
        ],
        "burn_after_reading": true
    }

# FILES
*~/.config/sidupaste/config.json*
: Default location of the sidupaste configuration. The file has to be
  created manually as it is not installed with a standard installation.

# AUTHORS
Bryan Frimin.
