package retry

import (
	"testing"

	"github.com/megaease/x/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestShouldMaxRetryNumReach(t *testing.T) {
	r := Default()
	err := r.SetOption(func(r *Retry) {
		r.Tries = 3
	}).SetOption(func(r *Retry) {
		r.RetryConditions = append(r.RetryConditions, func(r *Retry, err error) bool {
			return errors.IsConflict(err)
		})
	}).Do(func() error {
		return errors.Conflict("produce a confilic error")
	})

	assert.True(t, err == MaxRetryNumReachError || r.Attempts() != r.Tries+1,
		"retry should return max retry error")
}

func TestRetryShouldSucceed(t *testing.T) {
	r := Default()
	err := r.SetOption(func(r *Retry) {
		r.Tries = 3
	}).SetOption(func(r *Retry) {
		r.SucceedConditions = append(r.SucceedConditions, func(r *Retry, err error) bool {
			return errors.IsConflict(err)
		})
	}).Do(func() error {
		return errors.Conflict("produce a confilic error")
	})

	assert.True(t, err == nil && r.Attempts() == 1,
		"retry should succeed, but error")
}

func TestRetryShouldError(t *testing.T) {
	r := Default()
	err := r.SetOption(func(r *Retry) {
		r.Tries = 3
	}).SetOption(func(r *Retry) {
		r.SucceedConditions = append(r.SucceedConditions, func(r *Retry, err error) bool {
			return errors.IsConflict(err)
		})
	}).Do(func() error {
		return errors.BadSpec("produce a badSpec error")
	})

	assert.True(t, err != nil && r.Attempts() == 1,
		"retry should error, and try once")
}
