---
title: SIDUPASTE
header: Sidupaste Manual
footer: 1.0.1
date: Jan 20, 2022
section: 1
---
# NAME
**sidupaste** – create sidupaste paste with simple shell command

# SYNOPSIS
**sidupaste** [-help] [-bin=\<name\>] [-cfg-file=\<filename\>]\
\ \ \ \ \ \ \ \ \ \ \ \[-burn-after-reading] [-expire=\<time\>] [-formatter=\<format\>]\
\ \ \ \ \ \ \ \ \ \ \ \[-open-discussion] *STDIN*

# DESCRIPTION
A minimalist, open source command line interface for **Siduction Paste**
and other PrivateBin instances.

# OPTIONS
**-help**
: Show help message.

**-bin** \<name\>
: The sidupaste instance name.

**-burn-after-reading**
: Delete the paste after reading.

**-cfg-file** \<path\>
: The path of the configuration file (default
  "~/.config/sidupaste/config.json").

**-expire** \<time\>
: The time to live of the paste.

**-formatter** \<format\>
: The text formatter to use, can be plaintext, markdown or
  syntaxhighlighting.

**-open-discussion**
: Enable discussion on the paste.

# EXIT STATUS
The **sidupaste** utility exits 0 on success, and >0 if an error
occurs.

# EXAMPLES
Create a paste on the default sidupaste instance:

    $ cat example.txt | siduction

# SEE ALSO
**sidupaste.conf**(5)

# AUTHORS
Bryan Frimin.
