package main

import "fmt"

func FormatEvent(e Event) string {
	switch e.Type {

	case "PushEvent":
		return fmt.Sprintf(
			"Pushed %d commits to %s",
			len(e.Payload.Commits),
			e.Repo.Name,
		)

	case "IssuesEvent":
		return fmt.Sprintf("Opened issue in %s", e.Repo.Name)

	case "WatchEvent":
		return fmt.Sprintf("Starred %s", e.Repo.Name)

	case "ForkEvent":
		return fmt.Sprintf("Forked %s", e.Repo.Name)

	case "CreateEvent":
		return fmt.Sprintf("Created something in %s", e.Repo.Name)

	default:
		return fmt.Sprintf("%s in %s", e.Type, e.Repo.Name)
	}
}
