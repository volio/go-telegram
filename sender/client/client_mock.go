package client

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m *MockClient) DoPost(method string, v interface{}) error {
	args := m.Called(method, v)
	return args.Error(0)
}
