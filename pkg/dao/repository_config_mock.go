package dao

import (
	"testing"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryConfigDao struct {
	mock.Mock
}

func (r *MockRepositoryConfigDao) Create(newRepo api.RepositoryRequest) (api.RepositoryResponse, error) {
	args := r.Called(newRepo)
	rr, ok := args.Get(0).(api.RepositoryResponse)
	if ok {
		return rr, args.Error(1)
	} else {
		return api.RepositoryResponse{}, args.Error(1)
	}
}

func (r *MockRepositoryConfigDao) BulkCreate(newRepo []api.RepositoryRequest) ([]api.RepositoryResponse, []error) {
	args := r.Called(newRepo)
	rr, rrOK := args.Get(0).([]api.RepositoryResponse)
	er, erOK := args.Get(1).([]error)

	if rrOK && erOK {
		return rr, er
	} else if rrOK {
		return rr, nil
	} else if erOK {
		return nil, er
	} else {
		return nil, nil
	}
}

func (r *MockRepositoryConfigDao) Update(orgID string, uuid string, repoParams api.RepositoryRequest) error {
	args := r.Called(orgID, uuid, repoParams)
	return args.Error(0)
}

func (r *MockRepositoryConfigDao) Fetch(orgID string, uuid string) (api.RepositoryResponse, error) {
	args := r.Called(orgID, uuid)
	if args.Get(0) == nil {
		return api.RepositoryResponse{}, args.Error(0)
	}
	rr, ok := args.Get(0).(api.RepositoryResponse)
	if ok {
		return rr, args.Error(1)
	} else {
		return api.RepositoryResponse{}, args.Error(1)
	}
}

func (r *MockRepositoryConfigDao) List(
	orgID string,
	pageData api.PaginationData,
	filterData api.FilterData,
) (api.RepositoryCollectionResponse, int64, error) {
	args := r.Called(orgID, pageData, filterData)
	if args.Get(0) == nil {
		return api.RepositoryCollectionResponse{}, int64(0), args.Error(0)
	}
	rr, ok := args.Get(0).(api.RepositoryCollectionResponse)
	total, okTotal := args.Get(1).(int64)
	if ok && okTotal {
		return rr, total, args.Error(2)
	} else {
		return api.RepositoryCollectionResponse{}, int64(0), args.Error(2)
	}
}

func (r *MockRepositoryConfigDao) SavePublicRepos(urls []string) error {
	return nil
}

func (r *MockRepositoryConfigDao) Delete(orgID string, uuid string) error {
	args := r.Called(orgID, uuid)
	return args.Error(0)
}

func (r *MockRepositoryConfigDao) ValidateParameters(orgId string, req api.RepositoryValidationRequest, excludedUUIDs []string) (api.RepositoryValidationResponse, error) {
	r.Called(orgId, req)
	return api.RepositoryValidationResponse{}, nil
}

func NewMockRepositoryConfigDao(t *testing.T) *MockRepositoryConfigDao {
	m := &MockRepositoryConfigDao{}
	m.Mock.Test(t)

	t.Cleanup(func() { m.AssertExpectations(t) })

	return m
}
