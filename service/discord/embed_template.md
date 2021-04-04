One of your alerts has been triggered for the following reason(s):
```c
{{range .TrimmedReasons}}{{.}}
{{end}}
{{if .Trimmed}}{{.TrimmedCount}} more...{{end}}
```
You can view the item page on Universalis by clicking [this link]({{.PageURL}}).