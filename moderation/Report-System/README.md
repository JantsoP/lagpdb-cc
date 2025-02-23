# Report System

A custom command package for your server extending the functionality of the native `report` command!

## Features

* Larger Loggings (250 messages instead of 100)
* A report history listing the past reports
* Notifications to keep track and encourage communication
* And more!

## Planned features

| ℹ These features are planned and still under development. They'll be added in the future, most likely as optional user variables. |
| --- |

* [x] Custom messages in notification.
* [ ] Moderation menu using reactions for quicker access (actions being taken on the reported user).
* [ ] Quick report using a reaction.
* [ ] Warn users with too many reports automatically.

## Author, License, and Disclaimer

This custom command package was authored and is maintained by [l-zeuch](https://github.com/l-zeuch "l-zeuch's profile on GitHub").  
The author does not take any responsibility for bugs or other unintended behaviour caused by altered code beyond the intended and described necessary configuration, as stated in the [license agreement](../../LICENSE "License").

## Acknowledgements

Huge kudos to [NaruDevnote](https://www.github.com/NaruDevnote "NaruDevnote's profile on GitHub") as he participated in a manner I highly appreciate.
Thank you very much for helping out!

----

## Installing

Alright, now that the classic mumbo jumbo is done, let's get down to business!

Before you continue to read further, please take your time to read the leading comment in each custom command script.

The [main command](customReport.go.tmpl "main command") has a configuration area. Please note this short explanation on each variable:

* The report log is the channel where the reports are being sent to.
* Report-Discussion is the channel where users are being notified of updates regarding their report.
* Message-Listener is the custom command with the `\A` RegEx trigger. This custom command will listen for a message when prompted.

As you might have already seen, I noted down the trigger and trigger type in the leading comment in each file.

These are important - they'll allow you to use aliases such as `report`, `reportuser`, `ru`; `cancel`, `cancelreport`, `cr`; respectively. If you do not wish to use aliases, then make sure you choose a command trigger and one single command name - the command trigger type does not allow for multiple triggers, hence the usage of a RegEx trigger type.

This goes for all kind of information: The usage, trigger, trigger type, and a short description as well as license and author are all noted down in the heading comment, so reading it can save you quite a lot of troubleshooting.

Really think about what you do when you change anything! If you are unsure if what you currently do is a good idea, it probably is not!

| ⚠ The following section assumes that you have write access to the server's control panel! |
| ---- |

Navigate to your [control panel](https://yagpdb.xyz/manage "YAGPDB.xyz control panel") and select the server you want to add this package to. Make sure you are logged into the right account.

First off, navigate on your control panel to **Tools & Utilities > Moderation** and select the **General** tab. Find the switch labelled **Enable report command?** and confirm that it is off (it will turn red when that is the case). For a visual reference, take a look at the following screenshot:

###### Disabled report command

![disabled report command](https://i.imgur.com/9VW7BuS.png)

Now navigate to **Core > Command settings** and select the **Global** tab. We are going to create a new command override affecting the report command. Select this command in the dropdown and hit the switch **Enable specified commands?** so that it turns red. You'll find a visual reference as the following screenshot:

###### Command override

![command override](https://i.imgur.com/BnJSZE3.png)

----
| ℹ If you don't know how to add custom commands, or need a refresher, please follow [this tutorial](https://learn.yagpdb.xyz/the-custom-command-interface "How to add a custom command"). |
| --- |

Navigate to **Core > Custom commands**. If you want to keep it all neat and organised, create a new category called "Report System" or similar, just so that you know what kind of commands are in there.

Copy the code and create for each file a new custom command. Look at the leading comment and set up the trigger type and trigger accordingly.

As a final step on the control panel, exchange `CHANNEL-ID` for the according channel ID of your report log as well as your report-discussion, as well as `CC-ID` for the custom command ID of your `\A` RegEx trigger.

Navigate back to Discord and run `-ru dbSetup`. Please note that this sub-command is case sensitive. Upon success, YAGPDB will respond with
> Database primed, report numbers reset, system is ready to use!

| ✅ Your report-system is now set up! |
| ---- |

### Gallery / Various

The following section showcases various notes and screenshots related to this custom command package.

###### Default interface

The message each report starts with. No cancellation request, no responsible moderator.

![default message](https://i.imgur.com/tkHJmr7.png)

###### Pending request

This is how a report looks like when the reporting user requested a cancellation.

![report with cancellation](https://i.imgur.com/QMUaV6I.png)

###### Notification

The message that gets sent along certain actions.

![example notification](https://i.imgur.com/DY5VAc4.png)

#### Colour

There are several states a report can be in, and each one has an individual colour. This will make it easier for you and your staff to see which report needs the most attention (also everybody likes a little colour):

* ![#808080](https://cdn.discordapp.com/attachments/767771719720632350/793546124903317554/000000.png) Pending moderator review
* ![#FF00FF](https://cdn.discordapp.com/attachments/767771719720632350/793546157316898857/000000.png) Pending cancellation request
* ![#FFFF00](https://cdn.discordapp.com/attachments/767771719720632350/793546178070446140/000000.png) Under investigation
* ![#0000FF](https://cdn.discordapp.com/attachments/767771719720632350/793546199532699678/000000.png) Information requested
* ![#00FF00](https://cdn.discordapp.com/attachments/767771719720632350/793546218068115486/000000.png) Report resolved
* ![#FF0000](https://cdn.discordapp.com/attachments/767771719720632350/793546237483024394/000000.png) Cancellation request denied
