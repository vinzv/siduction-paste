# Preliminary note
This is a re-branded and slightly customized fork of Bryan Frimin's
"privatebin" which can be found [here](https://github.com/gearnode/privatebin).
We changed the name from privatebin to sidupaste and adapted file paths accordingly.

We want to thank Bryan for his great work on this handy tool!

# Introduction
This repository contains a CLI for privatebin server.

# Install
You can install the command line interface with:

## From source

    git clone https://github.com/siduction/sidupaste.git
    cd sidupaste
    make
    make install

# Usage
You can create paste from file with:

    cat resume.txt | sidupaste -bin demo

# Build
You can build the command line interface with:

    make build

# Documentation
The [handbook](doc/handbook.md) contains informations about various
aspects of the command line interface.

You can also use the standard Go documentation tool to read code
documentation, for example:

    go doc -all github.com/siduction/sidupaste


# Contact
If you find a bug or have any question, feel free to open a Github issue
at Brian's [upstream project](https://github.com/gearnode/privatebin) or to contact him [by email](mailto:bryan@frimin.fr).

Please note that he does not currently review or accept any contribution.

# Licence
Released under the ISC license.

Copyright (c) 2020-2022 Bryan Frimin <bryan@frimin.fr>.
Re-branding 2022 Vinzenz Vietzke <vinzv@siduction.org>.

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
