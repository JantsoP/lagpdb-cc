<h1 align="center"><img src="https://yagpdb.xyz/static/img/logo_y.png" height=32px width=32px></img>&nbspYAGPDB Custom Commands</h1>

<div align="center">
<a href="https://github.com/l-zeuch/lagpdb-cc/stargazers/"><img src="https://img.shields.io/github/stars/l-zeuch/lagpdb-cc?logo=github&style=for-the-badge"></a>
<img src="https://img.shields.io/github/repo-size/l-zeuch/lagpdb-cc?logo=github&style=for-the-badge">
<a href="https://github.com/l-zeuch/lagpdb-cc/blob/master/LICENSE"><img src="https://img.shields.io/github/license/l-zeuch/lagpdb-cc?style=for-the-badge"></a>
<a href="https://github.com/l-zeuch"><img src="https://img.shields.io/static/v1?label=Maintainer&message=l-zeuch&color=1f8b4c&style=for-the-badge"></a>
</div>

A collection of well-documented, quality custom commands for the [YAGPDB](https://yagpdb.xyz) Discord Bot.

## Adding These Custom Commands

Using these assumes at least basic knowledge of custom commands in YAGPDB. The classic phrase "don't use things you don't understand" applies very well here.
Make sure you at least understand the basic logic, this is why the code for each script is public - so you can review it and learn from it. [Mandatory read](https://learn.yagpdb.xyz/general-tips)

If you for some reason forgot how to add custom commands stop now and [RTFM](https://learn.yagpdb.xyz/the-custom-command-interface).
This and other repositories won't be of much use if you don't know how to add things to your server.

This should be obvious if you've read the license but don't be upset if things break - they're provided "as is", without a guarentee to work as expected, despite best efforts to test and review them, humans still make mistakes. Open an issue and explain what happened. Make sure to double-check the documentation, sending in bug reports which aren't bugs but user error is an easy way to get your issue closed without reply.

All source files are located in `src`, which is further divided into **categories**. Some custom commands are standalone, there are, however, also "systems" which require you to add all scripts - They will be collected in an individual folder to signal that as well as accompanied by a README.md file documenting the installation process and usage of said system.

Browse the repository for a while and once you found one you like, read its documentation for installation instructions.

### Leading Comment

As you might've already noticed, there is a leading comment in each custom command file - this is your small documentation, keep it. It *will* help you troubleshooting things / looking for a certain custom command. Further, it discloses the author, source, and license, which **must** be kept according to the license under which this repository is. So yeah, although unlikely, people could actually DMCA you for removing that bit of information. The leading comment looks similar to the following:

###### Example Header

```txt
Copyright (c): Jane Doe & John Smith, 2021.
Source: https://github.com/l-zeuch/lagpdb-cc
Licensed under the terms of BSD-3-Clause.
```

## Contributing

Before you go ahead and create a fork, read the [contributing guidelines](CONTRIBUTING.md) closely and carefully. It contains crucial information which may get your PR rejected if it becomes evident you did not read what you're supposed to read.

If you're unsure about some weird phrasing, open an issue and describe your questions.

## Legal Mumbo Jumbo

### Disclaimer

1. The YAGPDB developer, staff, and/or support are not responsible for any difficulties caused by these custom commands. These commands are not guaranteed to be working, use them at your own risk.

2. By modifying a script beyond reasonable modification and intended configuration it is to be understood as the user's own design and may not receive full and/or any support.

3. All scripts in this repository are not meant to cause any offence of any kind and are provided "as is".

### License

Unless otherwise specified, all work in this repository is licensed under the terms of the [3-Clause BSD License](LICENSE.md).

### Maintainer

This repository is maintained by [@l-zeuch](https://github.com/l-zeuch)

## Useful Links

* YAGPDB [Documentation](https://docs.yagpdb.xyz)
* YAGPDB [Learning resources](https://learn.yagpdb.xyz)
* YAGPDB [control panel](https://yagpdb.xyz/manage)
* YAGPDB [Support Server](https://discord.com/invite/4udtcA5)

----
Copyright (c): Luca Zeuch, 2021

3-Clause BSD License
