package validators

const (
	lostStatus   = "lost"
	wonStatus    = "won"
	activeStatus = "active"
)

type BetRequestValidator struct{}

func NewBetRequestValidator() *BetRequestValidator {
	return &BetRequestValidator{}
}

func (v *BetRequestValidator) StatusIsValid(status string) bool {
	if status == lostStatus || status == wonStatus || status == activeStatus {
		return true
	}
	return false
}
