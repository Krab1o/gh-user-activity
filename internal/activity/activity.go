package activity

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Repo struct {
    Name    string  `json:"name"`
}

type Event struct {
    Type        string          `json:"type"`
    Repo        struct {
        Name    string
    } `json:"repo"`
    Payload struct {
        RefType string          `json:"ref_type"`
        Action  string          `json:"action"`
        Commits []struct {
            Message string      `json:"comment"`
        }
    }
    HappenedAt   time.Time      `json:"created_at"`
}

func Activity(username string, count uint32) {
    url := "https://api.github.com/users/" + username + "/events?per_page=" + strconv.FormatUint(uint64(count), 10)
    req, err := http.NewRequest(http.MethodGet, url, nil)
    code, codeErr := os.ReadFile("code.txt")
    if codeErr != nil {
        log.Fatal("No file with authorization token found")
    }
    req.Header.Set("Authorization", "Bearer " + string(code))
    req.Header.Set("Accept", "application/vnd.github+json")
    if err != nil {
        log.Fatal(err)
    }
    response, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    
    events := []Event{}
    json.NewDecoder(response.Body).Decode(&events)
    for _, event := range events {
        switch event.Type {
        case "CreateEvent":
            fmt.Printf("Created repository titled \"%s\" at %s\n", event.Repo.Name, event.HappenedAt.Format(time.DateTime))
        case "DeleteEvent":
            fmt.Printf("Deleted %s in \"%s\" repository at %s\n", event.Payload.RefType, event.Repo.Name, event.HappenedAt.Format(time.DateTime))
        case "ForkEvent":
            fmt.Printf("Forked \"%s\" repository at %s\n", event.Repo.Name, event.HappenedAt.Format(time.DateTime))
        case "PushEvent":
            fmt.Printf("Pushed %d commits to \"%s\" repository at %s\n", len(event.Payload.Commits), event.Repo.Name, event.HappenedAt.Format(time.DateTime))
        case "WatchEvent":
            fmt.Printf("Starred \"%s\" repository at %s\n", event.Repo.Name, event.HappenedAt.Format(time.DateTime))
        default:
            fmt.Printf("%s happened at %s\n", event.Type, event.HappenedAt.Format(time.DateTime))
        }
    }
}