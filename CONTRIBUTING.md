If you're reading this, it is fair to assume that you know how to write a custom command and want to see it being added here. Before that can happen, a few things need to be done:

* Make sure your command's concept or idea isn't already represented here
* Extract all possible configuration to the top of your script, for ease of access. It's better to have a little more configuration than too few
* Please open all of your pull requests against the `dev` branch, unless you want to be yelled at by a nice bot 😉

If you're PRing a new custom command, please add a header comment to each file describing the command's function, its usage, the author and the repository, the trigger as well as the trigger type, like this:

######  Example leading comment
```go
{{/*
    This command puts the famous "Hello, World!" message in chat.

    Recommended trigger and trigger type: Command trigger, `test`
    Usage: test

    Copyright (c): Jane Doe & John Smith; 2021
    Source: <https://github.com/l-zeuch/lagpdb-cc>
    Licensed under the terms of BSD-3-Clause
*/}}
```
> If you want to waive copyright, simply leave the copyright notice out and state it in your pull request.

Thanks for contributing and happy hacking!