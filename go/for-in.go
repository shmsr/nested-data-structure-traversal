package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type section struct {
	Title               string `json:"title"`
	ResetLessonPosition bool   `json:"reset_lesson_position"`
	Position            uint8  `json:"position,omitempty"`
	Lessons             []struct {
		Name     string `json:"name"`
		Position uint8  `json:"position,omitempty"`
	} `json:"lessons"`
}

var sample = []byte(`
[
  {
    "title": "Getting started",
    "reset_lesson_position": false,
    "lessons": [
      {"name": "Welcome"},
      {"name": "Installation"}
    ]
  },

  {
    "title": "Basic operator",
    "reset_lesson_position": false,
    "lessons": [
      {"name": "Addition / Subtraction"},
      {"name": "Multiplication / Division"}
    ]
  },

  {
    "title": "Advanced topics",
    "reset_lesson_position": true,
    "lessons": [
      {"name": "Mutability"},
      {"name": "Immutability"}
    ]
  }
]
`)

func do(data []byte) ([]byte, error) {
	var src []section

	if err := json.Unmarshal(data, &src); err != nil {
		return nil, err
	}

	sectionCounter, lessonCounter := uint8(1), uint8(1)

	for i := range src {
		if src[i].ResetLessonPosition {
			lessonCounter = 1
		}

		src[i].Position = sectionCounter
		sectionCounter++

		for j := range src[i].Lessons {
			src[i].Lessons[j].Position = lessonCounter
			lessonCounter++
		}
	}

	dst, err := json.MarshalIndent(src, "", "\t")
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func main() {
	res, err := do(sample)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, string(res))
}
