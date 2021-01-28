# Report System v2

## Table of Contents
<details>
<summary>Table of Contents</summary>

* [Features](#Features)
* [Setting Up](#Setting-Up)
* [Usage](#Usage)
    * [Interface](#Interface)
        * [Reaction Menu](#Reaction-Menu)
        * [Default Reaction Inferface](../#Default-Reaction-Inferface)
        * [Pending Cancellation Request](..#Pending-Cancellation-Request)
        * [Notification Message](#Notification-Message)
* [Acknowledgements](../#Acknowledgements)
* [Planned Features](../#Planned-Features)
* [Author](../#Author)

</details>

## Features
* Logging Channel for report messages
* Cancellation requests
* Report history
* Reaction Menu
* Notifications
    * Custom Message in Notification

## Setting Up
| ⚠ You need `Manage_Server` permission in order to run the setup command! |
| --- |

Make for each custom command file a separate custom command, preferrably in the same category to keep them neat and organized. Please make sure to follow each step precisely and to use the correct trigger and trigger type, as well.

#### Here's what you have to do:
1. Disable the native report command, found here: `Control Panel > Tools & Utilities > Moderation`, like [this](#Disabled-Report-Command)
2. Configure the variables in [the main command](customReport.go.tmpl) as described there.
4. Run the **case sensitive** command `-ru dbSetup`
5. Done! YAGPDB.xyz will now take care of the rest and confirms setting up with an appropiate response.  

---
| 🛑 Pay attention to the trigger types and triggers! |
| --- |

For [the main command](customReport.go.tmpl) it is a RegEx trigger type with the trigger `\A-r(eport)?(?:u(ser)?)?(\s+|\z)`.  
[The second command](cancelReport.go.tmpl) requires a RegEx trigger as well, with the trigger being `\A-c(ancel)?r(eport)?(\s+|\z)`.  
The [reaction handler](reactionHandler.go.tmpl) needs a Reaction trigger with "Added reactions only" selected.  
Lastly, the [message listener](messageListener.go.tml) is a `\A` or `.*` RegEx trigger.

| ℹ Make sure to change `-` in both RegEx triggers to match YAGPDB's prefix in your server!<br/>It is also recommened to create a [command override](#Command-Override-Example) disabling the `report` command completely. |
| --- |

## Usage
### Commands
`-ru <User:Mention/ID> <Reason:Text>` - Sends the report.  
`-cr <MessageID:Text> <Key:Text> <Reason:Text>` - Requests cancellation of the latest report. The key is given to the reporting user in a custom command DM.

### Interface
| ℹ Only members with `Manage_Messages` permission will be able to use the reaction menu. |
| --- |

#### Reaction Menu
The bottom-most field in the embed will give you a short explanation on what the buttons do.
Please take some time to read the intention behind a few options:

* **Opening a report:** In order to be able to interact with a report, it must be opened; First come, first serve - Only the moderator who opened the report can interact with it.
* **Dismissing a report:** Some call it "ignoring". Both is fine. Basically it tells the reporting user that their report has no ground to stand upon on.
* **Requesting information:** You can see it as a step before ignoring, in case the reported user is a known case, but the report reason is not a very substantive one.  
* **Starting investigation:** This one should be obvious. Looking into it, reading the logs, discussing with other staff, talking with the reported user, finding a solution.  
* **Resolving a report:** Used when the reported user was punished accordingly or the report turns out to be for a bagatelle.

Of course, there are more options than just these four, however the missing ones are a fair bit clearer than these.

| ✅ Once a report is closed, YAGPDB.xyz will add a white flag (🏳️) as reaction to signalize a closed report. |
| --- |

#### Notification Message
![Notification Example](https://cdn.discordapp.com/attachments/767771719720632350/804028386255175680/unknown.png)

Your moderators are from time to time prompted to enter a message - This message will show up in the description of the notification message.

#### Screenshots

##### Command Override Example
![Command Override](https://cdn.discordapp.com/attachments/767771719720632350/795328377158369330/unknown.png)

##### Disabled Report Command
![Disable Report Cmd](https://cdn.discordapp.com/attachments/767771719720632350/795330583303028746/unknown.png)