# Contributing Guidelines

If you're reading this, it is fair to assume that you're considering a contribution to this repository. We appreciate any form of contribution, whether it be a new command, bugfixes, typo, or better documentation. The following guidelines will help you with your contribution.

## Folder Structure

To understand how this repository is structured:

* `.github/` contains all information necessary for GitHub: Issue and pull request templates, as well as workflows.
* `assets/` is where image files used in the documenting READMEs are stored.
* `src/` is the folder where all custom command source files are stored.

Browse for a while in this repository to get used to the folder structure.

# Guide

Now that you're familiar with the repository's structure, please take your time to read this guide carefully. Its main purpose is to make things easier for *you*, not us.

## Code Formatting

### Leading Comment

Add a brief comment to the start of your custom command explaining what this script does. Please note that this is not the place for an essay - users are asked to keep this comment at all costs when adding these scripts to their server, and as you know, this space is precious.

If necessary, refer to an accompanying README which further documents your custom command.

###### Do This

```text
This custom command does foo bar buz.
Refer to the README for more information.
```

###### Don't Do This

```text
This command does foo bar buz.

Usage:
    -foo
    -foo bar

Trigger Type: Command
Trigger: foo
```

After that, add some information about you, preferably in the following format:

###### Author

```text
Copyright (c): Name <https://github.com/your-handle-on-github>
```

Finally, add the source (this repository!) and the license, as such:

```text
Source: https://github.com/l-zeuch/lagpdb-cc
Licensed under the terms of BSD-3-Clause
```

That way, people will always know where your script is from and can easily add it themselves. After all this, your leading comment may look as follows:

###### Leading Comment Example

```go
{{/*
    This command sends a message in the chat.

    Copyright (c): Luca Z. <https://github.com/l-zeuch>
    Source: https://github.com/l-zeuch/lagpdb-cc
    Licensed under the terms of BSD-3-Clause
*/}}
```

### Add the Code Itself

Now it's time to find a fitting spot for your new custom command; Is it more focused towards games or having a fun time? It's probably fitting for `src/fun/`. In any case, try not to think too hard about it: When you make your PR, we can talk about this during the review process.

If you're adding a set of custom commands, please arrange them in a new folder within the appropriate category folder.

## Naming Conventions

* Use `snake_case` for folders and files, as per GoLang naming convention: `foor_bar/baz_buz.gotmpl`, not `fooBar/bazBuz.gotmpl`.
* Use consistent variable casing - if you start with `camelCase`, stick to it and don't switch to `snake_case` in the middle of your code.
  - Info for the conventionists: The Go naming convention dictates to use `camelCase` for multi-word names.

    It is also highly recommended to keep it short and simple, yet descriptive, meaning names like `targetUsernameAndDiscriminator` aren't really helpful.
* Use the file extension `.gotmpl` - This extension will grant syntax highlighting here on GitHub, as well as language recognition.

### Configuration Variables

This might be the only big change you really have to make to your code: Extract variables that can/should be used for configuration into constants in `UPPER_SNAKE_CASE` to the top, right underneath the leading comment. Always keep in mind what the end-user would probably want to configure - with that said, it's often enough more than you think.

Consider the following example of a "Hello World" program: A user might want to configure the message, so we'll extract it as follows:

###### Configuration Variable Example

```go
{{/*
    Leading comment goes here
*/}}

{{/* CONFIGURATION START */}}
{{$MESSAGE_SENT := "Hello, World!"}}
{{/* CONFIGURATION END */}}

{{$MESSAGE_SENT}}
```

It is recommended to provide default values where appropiate as well as use meaningful variable names, so that it's easier to see what exactly the configuration is for.

## Documentation

Good code has good documentation. These custom commands are no exception - how would you expect users to know how to use your command if you don't tell them?

### Documenting a Single Command

This is perhaps the easiest of all things: simly add your documentation to the folder's README. For your convenience, there is a [template](./.github/doc_template_single.md) which you can work with.

### Documenting a System

If you're adding a system, it's a bit more complicated.

### Is It Hard to Install?

Think about how hard it is to install your system, relative to other custom commmands. You might need to write an installation guide, especially when your system has parts users might be confused about, or lots of complex configuration. There is no template for this, but there are several good examples on the yagpdb-cc GitHub page, namely:

* [Version 2 of the starboard system](https://yagpdb-cc.github.io/fun/starboard/overview)
* [Version 1 of the custom report replacement](https://yagpdb-cc.github.io/moderation/report-system/overview)

Note that these use special markdown features which aren't available here, but you should get the gist of it.

### Is It Easy to Install?

If there's little to no configuration necessary, feel free to skip the installation guide and instead write a quick overview of your system, i.e. its variables and trigger(s).

For some tips on writing documentation in markdown, please view the next section.

# Markdown

Feel free to use any markdown syntax known to you - please note however that not everything that's markdown will be rendered on GitHub. GitHub uses a [special flavour](https://guides.github.com/features/mastering-markdown/), which you should at least read through if you're new to this.

You can generate annotations with the following syntax, which is essentially just a table header which then renders similar to annotations as seen on Docusaurus or Gitbook.

###### Annotation Example

  ```markdown
  | ℹ️ This is a hint. |
  | ---- |
  ```

Although technically permitted, please keep the usage of inline-HTML to a minimum, as this will make the raw document harder to read.

Only use HTML when it can't be done with pure markdown, for example dropdowns:

###### HTML Dropdown Demo

  ```html
  <details>
  <summary>Click me!</summary>

  More text which is hidden initially.
  </details>
  ```

Label your codeblocks with a level-6 heading (i.e. `######` 6 hashes), like presented throughout this document. Always give your codeblock a language tag - if there is no language and it's only text, use `txt` to show this intention:

###### Codeblock Language tagging

```markdown
    ```txt
    your text would go here.
    ```
```

# Submitting Your Code

After you've added your code, wrote the documentation and did a final proof-read, you're ready to open a pull request.

If you're new to this, or need a refresher, the [GitHub documentation](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request "Creating a Pull Request") explain it well.

Make sure to open your pull request against the `dev` branch . There is also a pull request template which is recommened to be used.

## What then?

Now all you have to do is to wait for a review. Please make sure to answer in a timely manner, to speed up things - you surely would want to see your command being added as fast as possible.

----
Thanks for considering a contribution!
