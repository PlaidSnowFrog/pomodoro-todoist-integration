package todoist

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type TodoistResponse struct {
	Tasks      []Task  `json:"results"`
	NextCursor *string `json:"next_cursor"` // nil when there are no more pages
}

type Task struct {
	ID          string   `json:"id"`
	Content     string   `json:"content"` // the task title, e.g. "test"
	Description string   `json:"description"`
	ProjectID   string   `json:"project_id"`
	SectionID   *string  `json:"section_id"`
	Priority    int      `json:"priority"` // 1 = normal ... 4 = urgent
	Labels      []string `json:"labels"`
	Checked     bool     `json:"checked"`
	AddedAt     string   `json:"added_at"`
	Due         *Due     `json:"due"` // nil if no due date
}

type Due struct {
	Date        string  `json:"date"`   // "2026-11-20" or with time "2026-11-20T15:00:00"
	String      string  `json:"string"` // human text, e.g. "today"
	IsRecurring bool    `json:"is_recurring"`
	Timezone    *string `json:"timezone"`
}

func GetTasks() []Task {
	client := &http.Client{}
	var all []Task
	cursor := ""

	for {
		q := url.Values{}
		q.Set("query", "today")
		if cursor != "" {
			q.Set("cursor", cursor)
		}

		req, err := http.NewRequest("GET", "https://api.todoist.com/api/v1/tasks/filter?"+q.Encode(), nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+os.Getenv("TODOIST_API_KEY"))

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error status code: %d, Body: %s", resp.StatusCode, string(body))
			return nil
		}

		var response TodoistResponse
		if err := json.Unmarshal(body, &response); err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		all = append(all, response.Tasks...)

		if response.NextCursor == nil || *response.NextCursor == "" {
			break
		}
		cursor = *response.NextCursor
	}

	return all
}
