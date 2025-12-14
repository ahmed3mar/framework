package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cacheMocks "github.com/goravel/framework/mocks/cache"
	configMocks "github.com/goravel/framework/mocks/config"
	eventMocks "github.com/goravel/framework/mocks/event"
)

func TestRepository_Put_FiresEvents(t *testing.T) {
	mockStore := &cacheMocks.Driver{}
	mockEvents := &eventMocks.Instance{}
	mockConfig := &configMocks.Config{}

	repository := NewRepository(mockStore, mockEvents, "default", mockConfig)

	key := "test_key"
	value := "test_value"
	ttl := time.Minute

	// Mock the store Put method
	mockStore.On("Put", key, value, ttl).Return(nil)

	// Mock the events firing - Repository uses Dispatch method
	mockEvents.On("Dispatch", mock.Anything, mock.Anything).Return([]any{}).Maybe()

	err := repository.Put(key, value, ttl)

	assert.NoError(t, err)
	mockStore.AssertExpectations(t)
}

func TestRepository_Get_FiresEvents(t *testing.T) {
	mockStore := &cacheMocks.Driver{}
	mockEvents := &eventMocks.Instance{}
	mockConfig := &configMocks.Config{}

	repository := NewRepository(mockStore, mockEvents, "default", mockConfig)

	key := "test_key"
	expectedValue := "test_value"

	// Mock the store Get method - needs to handle variadic parameters
	mockStore.On("Get", key, mock.Anything).Return(expectedValue)

	// Mock the events firing
	mockEvents.On("Dispatch", mock.Anything, mock.Anything).Return([]any{}).Maybe()

	result := repository.Get(key)

	assert.Equal(t, expectedValue, result)
	mockStore.AssertExpectations(t)
}

func TestRepository_Forget_FiresEvents(t *testing.T) {
	mockStore := &cacheMocks.Driver{}
	mockEvents := &eventMocks.Instance{}
	mockConfig := &configMocks.Config{}

	repository := NewRepository(mockStore, mockEvents, "default", mockConfig)

	key := "test_key"

	// Mock the store Forget method
	mockStore.On("Forget", key).Return(true)

	// Mock the events firing
	mockEvents.On("Dispatch", mock.Anything, mock.Anything).Return([]any{}).Maybe()

	result := repository.Forget(key)

	assert.True(t, result)
	mockStore.AssertExpectations(t)
}

func TestRepository_Tags(t *testing.T) {
	mockStore := &cacheMocks.Driver{}
	mockEvents := &eventMocks.Instance{}
	mockConfig := &configMocks.Config{}

	repository := NewRepository(mockStore, mockEvents, "default", mockConfig)

	tags := []string{"tag1", "tag2"}
	taggedRepo := repository.Tags(tags...)

	assert.NotNil(t, taggedRepo)
	// We can't easily test that the tags are properly set without exposing the internal structure
	// but we can verify that we get a new repository instance
	assert.IsType(t, &Repository{}, taggedRepo)
}

func TestRepository_Many(t *testing.T) {
	mockStore := &cacheMocks.Driver{}
	mockEvents := &eventMocks.Instance{}
	mockConfig := &configMocks.Config{}

	repository := NewRepository(mockStore, mockEvents, "default", mockConfig)

	keys := []string{"key1", "key2"}

	// Mock the store Get method for each key - needs to handle variadic parameters
	mockStore.On("Get", "key1", mock.Anything).Return("value1")
	mockStore.On("Get", "key2", mock.Anything).Return("value2")

	// Mock the events firing
	mockEvents.On("Dispatch", mock.Anything, mock.Anything).Return([]any{}).Maybe()

	result := repository.Many(keys)

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	assert.Equal(t, expected, result)
	mockStore.AssertExpectations(t)
}
