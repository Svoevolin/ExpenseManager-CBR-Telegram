// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package services

//go:generate minimock -i github.com/Svoevolin/workshop_1_bot/internal/services.RateStorage -o rate_storage_mock.go -n RateStorageMock -p services

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"github.com/gojuno/minimock/v3"
)

// RateStorageMock implements services.RateStorage
type RateStorageMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddRate          func(ctx context.Context, rate domain.Rate) (err error)
	inspectFuncAddRate   func(ctx context.Context, rate domain.Rate)
	afterAddRateCounter  uint64
	beforeAddRateCounter uint64
	AddRateMock          mRateStorageMockAddRate
}

// NewRateStorageMock returns a mock for services.RateStorage
func NewRateStorageMock(t minimock.Tester) *RateStorageMock {
	m := &RateStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddRateMock = mRateStorageMockAddRate{mock: m}
	m.AddRateMock.callArgs = []*RateStorageMockAddRateParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRateStorageMockAddRate struct {
	mock               *RateStorageMock
	defaultExpectation *RateStorageMockAddRateExpectation
	expectations       []*RateStorageMockAddRateExpectation

	callArgs []*RateStorageMockAddRateParams
	mutex    sync.RWMutex
}

// RateStorageMockAddRateExpectation specifies expectation struct of the RateStorage.AddRate
type RateStorageMockAddRateExpectation struct {
	mock    *RateStorageMock
	params  *RateStorageMockAddRateParams
	results *RateStorageMockAddRateResults
	Counter uint64
}

// RateStorageMockAddRateParams contains parameters of the RateStorage.AddRate
type RateStorageMockAddRateParams struct {
	ctx  context.Context
	rate domain.Rate
}

// RateStorageMockAddRateResults contains results of the RateStorage.AddRate
type RateStorageMockAddRateResults struct {
	err error
}

// Expect sets up expected params for RateStorage.AddRate
func (mmAddRate *mRateStorageMockAddRate) Expect(ctx context.Context, rate domain.Rate) *mRateStorageMockAddRate {
	if mmAddRate.mock.funcAddRate != nil {
		mmAddRate.mock.t.Fatalf("RateStorageMock.AddRate mock is already set by Set")
	}

	if mmAddRate.defaultExpectation == nil {
		mmAddRate.defaultExpectation = &RateStorageMockAddRateExpectation{}
	}

	mmAddRate.defaultExpectation.params = &RateStorageMockAddRateParams{ctx, rate}
	for _, e := range mmAddRate.expectations {
		if minimock.Equal(e.params, mmAddRate.defaultExpectation.params) {
			mmAddRate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddRate.defaultExpectation.params)
		}
	}

	return mmAddRate
}

// Inspect accepts an inspector function that has same arguments as the RateStorage.AddRate
func (mmAddRate *mRateStorageMockAddRate) Inspect(f func(ctx context.Context, rate domain.Rate)) *mRateStorageMockAddRate {
	if mmAddRate.mock.inspectFuncAddRate != nil {
		mmAddRate.mock.t.Fatalf("Inspect function is already set for RateStorageMock.AddRate")
	}

	mmAddRate.mock.inspectFuncAddRate = f

	return mmAddRate
}

// Return sets up results that will be returned by RateStorage.AddRate
func (mmAddRate *mRateStorageMockAddRate) Return(err error) *RateStorageMock {
	if mmAddRate.mock.funcAddRate != nil {
		mmAddRate.mock.t.Fatalf("RateStorageMock.AddRate mock is already set by Set")
	}

	if mmAddRate.defaultExpectation == nil {
		mmAddRate.defaultExpectation = &RateStorageMockAddRateExpectation{mock: mmAddRate.mock}
	}
	mmAddRate.defaultExpectation.results = &RateStorageMockAddRateResults{err}
	return mmAddRate.mock
}

// Set uses given function f to mock the RateStorage.AddRate method
func (mmAddRate *mRateStorageMockAddRate) Set(f func(ctx context.Context, rate domain.Rate) (err error)) *RateStorageMock {
	if mmAddRate.defaultExpectation != nil {
		mmAddRate.mock.t.Fatalf("Default expectation is already set for the RateStorage.AddRate method")
	}

	if len(mmAddRate.expectations) > 0 {
		mmAddRate.mock.t.Fatalf("Some expectations are already set for the RateStorage.AddRate method")
	}

	mmAddRate.mock.funcAddRate = f
	return mmAddRate.mock
}

// When sets expectation for the RateStorage.AddRate which will trigger the result defined by the following
// Then helper
func (mmAddRate *mRateStorageMockAddRate) When(ctx context.Context, rate domain.Rate) *RateStorageMockAddRateExpectation {
	if mmAddRate.mock.funcAddRate != nil {
		mmAddRate.mock.t.Fatalf("RateStorageMock.AddRate mock is already set by Set")
	}

	expectation := &RateStorageMockAddRateExpectation{
		mock:   mmAddRate.mock,
		params: &RateStorageMockAddRateParams{ctx, rate},
	}
	mmAddRate.expectations = append(mmAddRate.expectations, expectation)
	return expectation
}

// Then sets up RateStorage.AddRate return parameters for the expectation previously defined by the When method
func (e *RateStorageMockAddRateExpectation) Then(err error) *RateStorageMock {
	e.results = &RateStorageMockAddRateResults{err}
	return e.mock
}

// AddRate implements services.RateStorage
func (mmAddRate *RateStorageMock) AddRate(ctx context.Context, rate domain.Rate) (err error) {
	mm_atomic.AddUint64(&mmAddRate.beforeAddRateCounter, 1)
	defer mm_atomic.AddUint64(&mmAddRate.afterAddRateCounter, 1)

	if mmAddRate.inspectFuncAddRate != nil {
		mmAddRate.inspectFuncAddRate(ctx, rate)
	}

	mm_params := RateStorageMockAddRateParams{ctx, rate}

	// Record call args
	mmAddRate.AddRateMock.mutex.Lock()
	mmAddRate.AddRateMock.callArgs = append(mmAddRate.AddRateMock.callArgs, &mm_params)
	mmAddRate.AddRateMock.mutex.Unlock()

	for _, e := range mmAddRate.AddRateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddRate.AddRateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddRate.AddRateMock.defaultExpectation.Counter, 1)
		mm_want := mmAddRate.AddRateMock.defaultExpectation.params
		mm_got := RateStorageMockAddRateParams{ctx, rate}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddRate.t.Errorf("RateStorageMock.AddRate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddRate.AddRateMock.defaultExpectation.results
		if mm_results == nil {
			mmAddRate.t.Fatal("No results are set for the RateStorageMock.AddRate")
		}
		return (*mm_results).err
	}
	if mmAddRate.funcAddRate != nil {
		return mmAddRate.funcAddRate(ctx, rate)
	}
	mmAddRate.t.Fatalf("Unexpected call to RateStorageMock.AddRate. %v %v", ctx, rate)
	return
}

// AddRateAfterCounter returns a count of finished RateStorageMock.AddRate invocations
func (mmAddRate *RateStorageMock) AddRateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddRate.afterAddRateCounter)
}

// AddRateBeforeCounter returns a count of RateStorageMock.AddRate invocations
func (mmAddRate *RateStorageMock) AddRateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddRate.beforeAddRateCounter)
}

// Calls returns a list of arguments used in each call to RateStorageMock.AddRate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddRate *mRateStorageMockAddRate) Calls() []*RateStorageMockAddRateParams {
	mmAddRate.mutex.RLock()

	argCopy := make([]*RateStorageMockAddRateParams, len(mmAddRate.callArgs))
	copy(argCopy, mmAddRate.callArgs)

	mmAddRate.mutex.RUnlock()

	return argCopy
}

// MinimockAddRateDone returns true if the count of the AddRate invocations corresponds
// the number of defined expectations
func (m *RateStorageMock) MinimockAddRateDone() bool {
	for _, e := range m.AddRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddRateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddRate != nil && mm_atomic.LoadUint64(&m.afterAddRateCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddRateInspect logs each unmet expectation
func (m *RateStorageMock) MinimockAddRateInspect() {
	for _, e := range m.AddRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RateStorageMock.AddRate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddRateCounter) < 1 {
		if m.AddRateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RateStorageMock.AddRate")
		} else {
			m.t.Errorf("Expected call to RateStorageMock.AddRate with params: %#v", *m.AddRateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddRate != nil && mm_atomic.LoadUint64(&m.afterAddRateCounter) < 1 {
		m.t.Error("Expected call to RateStorageMock.AddRate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RateStorageMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddRateInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RateStorageMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RateStorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddRateDone()
}
