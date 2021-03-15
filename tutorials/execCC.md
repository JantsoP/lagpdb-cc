# Executing a custom command with a custom command
So, you've came along advanced custom commands with YAGPDB and wondered how you can execute another custom command with a custom command and what all this fuss is about? Then this detailed introduction is exactly for you!

## execCC
For now, we'll concentrate on `execCC`; There are two more functions - `scheduleUniqueCC` and `cancelScheduledUniqueCC`- but these will be investigated in a later section.

### Syntax
Let's get started with the syntax first, as this is most likely the reason why you are here in the first place:

```go
{{execCC CC-ID Channel Delay Data}}
```
Alright, that seems pretty strange for now, but don't worry, it is actually quite easy:

#### CC-ID
This is typically one of either a set ID like 5 or the special `.CCID` template, representing the current custom command's ID.

##### How do I get that CC-ID?
You can do so by either using the `.CCID` template, or by taking a look at the following example; The number inside the marker box is the custom command ID, without the `#`:

![example](https://i.imgur.com/JA19tcC.png)

#### Channel
This obviously represents the channel of execution. You may use `nil` or `.Channel.ID` for the current channel, however a set channel ID or a channel name may also be used.

Despite being possible, it is not recommended to use a channel name, as this might break your custom command should the channel name change. Using an ID is definitely more robust.

#### Delay
This argument should be pretty straigt forward: You give a delay in seconds - This might range from `0` to a little over 292 years (9.2 x 10<sup>9</sup> seconds).

Although definitely possible, it should be obvious that using such long delays is quite unfeasible; On the one hand, these numbers will get ridicously long, on the other hand of what use would be a delay of, say, 100 years?

It is worth to note that, although floating point numbers such as `0.5` are supported, they will be truncuated to just their integer value, is this case `0`.

#### Data
Now lastly `Data` can be anything you want, to let your executed CC deal with it!
This section is a little more elaborate, so be sure to stick with me!

##### What exactly is that ominous data?
This is quite simple: Any data you want to hand over. Be that a slice, sdict, dict, single strings, numbers, etc.
You then can access this data with `.ExecData` and continue operating on that.
Note that sdicts and slices are a bit weird to deal with in `.ExecData`, as you have to convert them when retrieving, such as:

###### sdict in .ExecData example
```go
{{if not .ExecData}}
    {{/* Initializing an sdict and sending it over as data with the execCC template. */}}
    {{$sdictToSend := sdict "key1" "This is a test" "key2" 123456789}}
    {{execCC .CCID nil 30 $sdictToSend}}
{{else}}
    {{/* Accessing the data through .ExecData and converting it back to an sdict. */}}
    {{$sdictRetrieved := sdict .ExecData}}
{{end}}
```
You can read on that behaviour [here](https://docs.yagpdb.xyz/reference/templates#templates-sdict).

#### So, how would a full example look like?

> Assume we want to hand over a slice of strings `"This" "is" "a" "test"` to the CC with the ID `5` in 30 seconds in the current channel. Please take a moment to think about it before revealing the solution.

<details>
<summary>Click here to reveal a possible solution.</summary>

```go
{{$slice := cslice "This" "is" "a" "test"}}
{{execCC 5 nil 30 $slice}}
```
*Similarily, for dict and sdict you can do the same. Just keep their syntax in mind. For simple variables such as the string `"Hello"` or a number `123`, you can just put that in place of `$slice` on the second line.*
</details> 

### Why execCC?
If you read this far, you probably want to know what the benefits of `execCC`! If not, you can now go away : )

* Delayed operations - execCC can handle delays up to approximately 293 years, while `sleep` (the only real alternative) has significant limitations and is limited to 60 seconds at maximum.
* Chaining of custom commands - Useful to "increase" the limit of 10 000 characters per custom command.
* Creates a new instance (or execution) of the custom command - Useful to increase a few limits in custom commands.

| ⚠ To prevent abuse of this, custom commands executed with `execCC` are strictly limited to 10 per channel per minute. |
| --- |

## scheduleUniqueCC
Let us take a look at a similar but also quite different function to `execCC`: `scheduleUniqueCC`.

Contrary to `execCC`, `scheduleUniqueCC` overrides the old instance of the executed custom command - This is the reason we must provide a key:

### Syntax
The syntax is very similar to the one of `exeCC`, with one minor difference:
```go
{{scheduleUniqueCC CC-ID Channel Delay Key Data}}
```
As you can see, the only difference to `execCC` is that we give it a key, the rest is the same.

#### What is that "key"?
The key is used to keep track of the currently executed instance of the custom command. You can use anything as key, however it is a good practice to use something that tells you why and what, such as a single word or a combination of words.

If you run the function again, it will override the old execution.

## cancelScheduledUniqueCC
As the name suggests, this function is used to cancel a scheduled CC. This is only available through using `scheduleUniqueCC` earlier, as this function requires a key and the CC-ID:
```go
{{cancelScheduledUniqueCC CC-ID Key}}
```
Make sure to give it the correct key, as those are case sensitive. You cannot use it in combination with `execCC`, because that is lacking the required key argument: Without a key, YAGPDB does not know to what instance you are referring to.

### Why schedule a CC?
Now that you came this far, you surely want to know why you would want to ever schedule a custom command:

* Only one instance at all times (per key)
* Execution of the custom command can be stopped with `cancelScheduledUniqueCC`

## When do I use what?
That comes down to the individual use-case, actually. In general you can say when you want only one instance of the executed CC at all times, for example in bumping system (the bump can only be done once per interval), you use `scheduleUniqueCC`. 

When you want to use the executed custom command to make user-specific operations, resort to `execCC`.

## Conclusion
* you can execute a custom command with a custom command by using the functions `execCC` or `scheduleUniqueCC`.
* `execCC` creates a new instance of the executed custom command, while `scheduleUniqueCC` overrides the old instance with the given key.
* Think about when what would be of better use; sometimes you want to schedule a CC, sometimes you don't want to do that.

#### Footnotes
* [execCC related functions in the documentation](https://docs.yagpdb.xyz/reference/templates#execcc)
* [serialization behaviour in database interactions and execCC data](https://docs.yagpdb.xyz/reference/templates#templates-sdict)
