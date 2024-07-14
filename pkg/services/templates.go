package services

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type TemplateData struct {
	CurrentYear int
	TimeElapsed string
}

func FormatDuration(duration string) string {
	var hours, minutes time.Duration
	if len(duration) < 8 && strings.Contains(duration, ":") {
		var formattedDuration strings.Builder
		formattedDuration.WriteString("00:")
		formattedDuration.WriteString(duration)
		return formattedDuration.String()
	} else if len(duration) < 8 {
		timestamp, _ := strconv.Atoi(duration)
		totalMinutes := time.Duration(timestamp)
		hours = (totalMinutes * time.Minute) / time.Hour
		minutes = totalMinutes - (hours*time.Hour)/time.Minute
		return fmt.Sprintf("%02d:%02d:00", hours, minutes)
	} else {
		return duration
	}
}

var functions = template.FuncMap{
	"formatDuration": FormatDuration,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}
	fmt.Println(pages)
	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))

		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))

		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
