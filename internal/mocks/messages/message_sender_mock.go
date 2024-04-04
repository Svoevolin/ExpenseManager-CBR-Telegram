// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package messages

//go:generate minimock -i github.com/Svoevolin/workshop_1_bot/internal/model/messages.MessageSender -o message_sender_mock_test.go -n MessageSenderMock -p messages

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MessageSenderMock implements messages.MessageSender
type MessageSenderMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSendMessage          func(text string, userID int64, keyboardRows ...map[string]string) (err error)
	inspectFuncSendMessage   func(text string, userID int64, keyboardRows ...map[string]string)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mMessageSenderMockSendMessage
}

// NewMessageSenderMock returns a mock for messages.MessageSender
func NewMessageSenderMock(t minimock.Tester) *MessageSenderMock {
	m := &MessageSenderMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMessageMock = mMessageSenderMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*MessageSenderMockSendMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageSenderMockSendMessage struct {
	mock               *MessageSenderMock
	defaultExpectation *MessageSenderMockSendMessageExpectation
	expectations       []*MessageSenderMockSendMessageExpectation

	callArgs []*MessageSenderMockSendMessageParams
	mutex    sync.RWMutex
}

// MessageSenderMockSendMessageExpectation specifies expectation struct of the MessageSender.SendMessage
type MessageSenderMockSendMessageExpectation struct {
	mock    *MessageSenderMock
	params  *MessageSenderMockSendMessageParams
	results *MessageSenderMockSendMessageResults
	Counter uint64
}

// MessageSenderMockSendMessageParams contains parameters of the MessageSender.SendMessage
type MessageSenderMockSendMessageParams struct {
	text         string
	userID       int64
	keyboardRows []map[string]string
}

// MessageSenderMockSendMessageResults contains results of the MessageSender.SendMessage
type MessageSenderMockSendMessageResults struct {
	err error
}

// Expect sets up expected params for MessageSender.SendMessage
func (mmSendMessage *mMessageSenderMockSendMessage) Expect(text string, userID int64, keyboardRows ...map[string]string) *mMessageSenderMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageSenderMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageSenderMockSendMessageExpectation{}
	}

	mmSendMessage.defaultExpectation.params = &MessageSenderMockSendMessageParams{text, userID, keyboardRows}
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the MessageSender.SendMessage
func (mmSendMessage *mMessageSenderMockSendMessage) Inspect(f func(text string, userID int64, keyboardRows ...map[string]string)) *mMessageSenderMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for MessageSenderMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by MessageSender.SendMessage
func (mmSendMessage *mMessageSenderMockSendMessage) Return(err error) *MessageSenderMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageSenderMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageSenderMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &MessageSenderMockSendMessageResults{err}
	return mmSendMessage.mock
}

// Set uses given function f to mock the MessageSender.SendMessage method
func (mmSendMessage *mMessageSenderMockSendMessage) Set(f func(text string, userID int64, keyboardRows ...map[string]string) (err error)) *MessageSenderMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the MessageSender.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the MessageSender.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	return mmSendMessage.mock
}

// When sets expectation for the MessageSender.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mMessageSenderMockSendMessage) When(text string, userID int64, keyboardRows ...map[string]string) *MessageSenderMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageSenderMock.SendMessage mock is already set by Set")
	}

	expectation := &MessageSenderMockSendMessageExpectation{
		mock:   mmSendMessage.mock,
		params: &MessageSenderMockSendMessageParams{text, userID, keyboardRows},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up MessageSender.SendMessage return parameters for the expectation previously defined by the When method
func (e *MessageSenderMockSendMessageExpectation) Then(err error) *MessageSenderMock {
	e.results = &MessageSenderMockSendMessageResults{err}
	return e.mock
}

// SendMessage implements messages.MessageSender
func (mmSendMessage *MessageSenderMock) SendMessage(text string, userID int64, keyboardRows ...map[string]string) (err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(text, userID, keyboardRows...)
	}

	mm_params := MessageSenderMockSendMessageParams{text, userID, keyboardRows}

	// Record call args
	mmSendMessage.SendMessageMock.mutex.Lock()
	mmSendMessage.SendMessageMock.callArgs = append(mmSendMessage.SendMessageMock.callArgs, &mm_params)
	mmSendMessage.SendMessageMock.mutex.Unlock()

	for _, e := range mmSendMessage.SendMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendMessage.SendMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessage.SendMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessage.SendMessageMock.defaultExpectation.params
		mm_got := MessageSenderMockSendMessageParams{text, userID, keyboardRows}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("MessageSenderMock.SendMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the MessageSenderMock.SendMessage")
		}
		return (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(text, userID, keyboardRows...)
	}
	mmSendMessage.t.Fatalf("Unexpected call to MessageSenderMock.SendMessage. %v %v %v", text, userID, keyboardRows)
	return
}

// SendMessageAfterCounter returns a count of finished MessageSenderMock.SendMessage invocations
func (mmSendMessage *MessageSenderMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of MessageSenderMock.SendMessage invocations
func (mmSendMessage *MessageSenderMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to MessageSenderMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mMessageSenderMockSendMessage) Calls() []*MessageSenderMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*MessageSenderMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *MessageSenderMock) MinimockSendMessageDone() bool {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendMessageInspect logs each unmet expectation
func (m *MessageSenderMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageSenderMock.SendMessage with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageSenderMock.SendMessage")
		} else {
			m.t.Errorf("Expected call to MessageSenderMock.SendMessage with params: %#v", *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		m.t.Error("Expected call to MessageSenderMock.SendMessage")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageSenderMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendMessageInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageSenderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageSenderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendMessageDone()
}
