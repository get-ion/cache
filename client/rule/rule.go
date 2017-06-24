package rule

import (
	"github.com/get-ion/ion/context"
)

// Rule a superset of validators
type Rule interface {
	Claim(ctx context.Context) bool
	Valid(ctx context.Context) bool
}
