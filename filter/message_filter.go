package filters

import (
	"regexp"
	"strings"

	"github.com/purkhanov/gogram/commands"
	"github.com/purkhanov/gogram/types"
)

type MessageFilter func(*types.Message) bool

func TextEquals(text string) MessageFilter {
	return func(m *types.Message) bool {
		return m.Text == text
	}
}

func TextContains(substring string) MessageFilter {
	return func(m *types.Message) bool {
		return strings.Contains(m.Text, substring)
	}
}

func TextMatches(pattern string) MessageFilter {
	return func(m *types.Message) bool {
		matched, _ := regexp.MatchString(pattern, m.Text)
		return matched
	}
}

func IsCommand(command commands.Command) MessageFilter {
	return func(m *types.Message) bool {
		return m.Text == string(command)
	}
}
