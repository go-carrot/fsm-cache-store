package cachestore

import "errors"

type cachedTraverser struct {
	uuid         string
	currentState string
	Data         map[string]interface{}
}

func (c *cachedTraverser) UUID() string {
	return c.uuid
}

func (c *cachedTraverser) SetUUID(newUUID string) {
	c.uuid = newUUID
}

func (c *cachedTraverser) CurrentState() string {
	return c.currentState
}

func (c *cachedTraverser) SetCurrentState(newState string) {
	c.currentState = newState
}

func (c *cachedTraverser) Upsert(key string, value interface{}) error {
	c.Data[key] = value
	return nil
}

func (c *cachedTraverser) Fetch(key string) (interface{}, error) {
	if val, ok := c.Data[key]; ok {
		return val, nil
	}
	return nil, errors.New("Key `" + key + "` is not set")
}

func (c *cachedTraverser) Delete(key string) error {
	delete(c.Data, key)
	return nil
}
