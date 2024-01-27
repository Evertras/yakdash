package captured

import "sync"

type capturedOutput struct {
	sync.RWMutex
	data []byte
}

func newCapturedOutput() *capturedOutput {
	return &capturedOutput{
		data: make([]byte, 0, 1024),
	}
}

func (c *capturedOutput) Write(data []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()

	c.data = append(c.data, data...)

	return len(data), nil
}

func (c *capturedOutput) reset() {
	c.Lock()
	defer c.Unlock()

	c.data = make([]byte, 0, 1024)
}

func (c *capturedOutput) String() string {
	c.RLock()
	defer c.RUnlock()

	return string(c.data)
}
