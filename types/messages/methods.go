package messages

import (
	"github.com/statping-ng/statping-ng/utils"
)

func (m *Message) IsActive() bool {
	curr := utils.Now()
	return curr.Before(m.EndOn) && curr.After(m.StartOn)
}
