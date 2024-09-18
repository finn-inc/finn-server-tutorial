package registry_mock

type MockDB struct {
	client MockDBClient
}

func (d *MockDB) GetClient() MockDBClient {
	return d.client
}

type MockDBClient struct{}

func NewMockDB() MockDBClient {
	return MockDBClient{}
}
