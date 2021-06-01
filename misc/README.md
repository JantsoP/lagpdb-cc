# Miscelleanous
This folder contains various custom commands which don't fit well into other categories.

## List of Custom Commands
### Bookmark
This custom command bookmarks messages, sending context such as channel, a small custom message as well as a jump-link into your direct messages.

#### Installing
Add this custom command in a new slot, as usual. It is recommended to use the following configuration:

```yaml
Trigger Type: RegEx, case insensitive
Trigger: \A(?:-\s?|<@!?204255221017214977>\s*)b(?:ook)?m(?:ark)?(?:\s+|\z)
```

Don't forget to save and have your DMs open to YAGPDB.

#### Gallery
###### Bookmark Message
![demo bookmark](../assets/bookmark-demo.png)