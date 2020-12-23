{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command is basically the native report-command, but adds some back-end functionalites in order for the rest to work :)

    Usage: `-ru <User:Mention/ID> <Reason:Text>`

    Recommended Trigger type and trigger: Regex trigger with trigger `\A-r(eport)?u(ser)?(\s+|\z)`

    Created by: Olde#7325
*/}}

{{/*CONFIG AREA START*/}}

{{$reportLog := 772251753173221386}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 766370841196888104}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL CODE*/}}
{{$p := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "prefix")) 0 1}}
{{$Escaped_Prefix := reReplace `[\.\[\]\-\?\!\\\*\{\}\(\)\|]` $p `\${0}`}}
{{if not (reFind (print `\A` $Escaped_Prefix `|<@!?204255221017214977>`) .Message.Content)}}
{{$response:= sendMessageRetID nil "Did not set regex to match Server Prefix!"}}{{deleteTrigger}}
{{deleteMessage nil $response}}{{else}}
{{if and .CmdArgs (lt (len .CmdArgs) 2)}}
    {{if eq (index .CmdArgs 0) "dbSetup"}}
        {{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") "ManageServer")}}
            {{dbSet 2000 "reportSettings" (sdict "reportLog" (str $reportLog) "reportDiscussion" (str $reportDiscussion))}}
            {{dbSet 2000 "ReportNo" 0}}
            {{sendMessage nil "**Database primed, report count reset, system is ready to use!**"}}
        {{else}}
            {{sendMessage nil "You do not have permission to use this command!"}}
        {{end}}
    {{else}}
    {{sendMessage nil (printf "```%s <User:Mention/ID> <Reason:Text>``` \n Not enough arguments passed." .Cmd)}}
    {{end}}
{{else if not .CmdArgs}}
    {{sendMessage nil (printf "```%s <User:Mention/ID> <Reason:Text>``` \n Not enough arguments passed." .Cmd)}}
{{else}}
    {{$user := userArg (index .CmdArgs 0)}}
    {{if eq $user.ID .User.ID}}
        {{$silly := sendMessageRetID nil "You can't report yourself, silly."}}
        {{deleteTrigger}}
        {{deleteMessage nil $silly}}
    {{else}}
        {{$secret := adjective}}
        {{$logs250 := execAdmin "log" "250"}}
        {{$reason := joinStr " " (slice .CmdArgs 1)}}
        {{$history := ""}}
        {{if (dbGet $user.ID "rhistory")}}
            {{range (dbGetPattern $user.ID "rhistory%" 7 0)}}
                {{$history = .Value}}
            {{end}}
            {{dbSet $user.ID "rhistory" (print (dbGet $user.ID "rhistory").Value "\n" (currentTime.Format "02-01-2006-15:04:05") ": " $reason)}}
        {{else}}
            {{dbSet $user.ID "rhistory" (print (currentTime.Format "02-01-2006-15:04:05") ": " $reason)}}
        {{end}}
        {{$reportNo := dbIncr 2000 "ReportNo" 1}}
        {{$reportEmbed := cembed "title" (print "Report No. " $reportNo)
            "author" (sdict "name" (printf "%s (ID %d)" .User.String .User.ID) "icon_url" (.User.AvatarURL "256"))
            "thumbnail" (sdict "url" ($user.AvatarURL "512"))
            "description" (printf "<@%d> reported <@%d> in <#%d>." .User.ID $user.ID .Channel.ID)
            "fields" (cslice
                (sdict "name" "Current State" "value" "__Not reviewed yet.__")
                (sdict "name" "Reason for Report" "value" $reason)
                (sdict "name" "Reported user" "value" (printf "<@%d> (ID %d)" $user.ID $user.ID))
                (sdict "name" "Message Logs" "value" (printf "[last 250 messages](%s) \nTime - `%s`" $logs250 (currentTime.Format "Mon 02 Jan 15:04:05")))
                (sdict "name" "History" "value" (print "```\n" (or $history "None recorded") "\n```"))
                (sdict "name" "Reaction Menu Options" "value" (printf "\nDismiss report with ❌, put under investigation with 🛡️, or request more background information with ⚠️."))
            )
            "footer" (sdict "text" "No moderator yet • Claim with any reaction")
        }}
        {{$x := sendMessageRetID $reportLog $reportEmbed}}
        {{addMessageReactions $reportLog $x "❌" "🛡️" "⚠️"}}
        {{$response := sendMessageRetID nil "User reported to the proper authorites!"}}
        {{dbSet .User.ID "key" $secret}}
        {{deleteTrigger}}
        {{deleteMessage nil $response}}
        {{sendDM (printf "User reported to the proper authorities! If you wish to cancel your report, simply type \n```-cancelr %d %s``` in any channel.\n **A reason is required.**" $x $secret)}}
    {{end}}
{{end}}{{end}}