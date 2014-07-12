#elementary Fixes Bot

This bot fetches fixed bugs from launchpad and then tweets about it.

tweets accessible at [@elementaryfixes](https://twitter.com/elementaryfixes)

The old bot was written in python and suffered from lots of crashes due to launchpadlib, this bot uses a custom function to fetch the bugs and is much more stable.


##To run

There is a sample config file provided, add your configuration to that file and rename it to bot.cfg to use it.

    $ make
    $ ./elementary-fixes
