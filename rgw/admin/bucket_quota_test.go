package admin

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *RadosGWTestSuite) TestBucketQuota() {
	suite.SetupConnection()
	co, err := New(suite.endpoint, suite.accessKey, suite.secretKey, newDebugHTTPClient(http.DefaultClient))
	assert.NoError(suite.T(), err)

	s3, err := newS3Agent(suite.accessKey, suite.secretKey, suite.endpoint, true)
	assert.NoError(suite.T(), err)

	err = s3.createBucket(suite.bucketTestName)
	assert.NoError(suite.T(), err)

	suite.T().Run("set bucket quota but no user is specified", func(t *testing.T) {
		err := co.SetIndividualBucketQuota(context.Background(), IndividualBucketQuotaSpec{})
		assert.Error(suite.T(), err)
		assert.EqualError(suite.T(), err, errMissingUserID.Error())
	})

	suite.T().Run("set bucket quota but no bucket is specified", func(t *testing.T) {
		err := co.SetIndividualBucketQuota(context.Background(), IndividualBucketQuotaSpec{UID: "admin"})
		assert.Error(suite.T(), err)
		assert.EqualError(suite.T(), err, errMissingBucket.Error())
	})

	suite.T().Run("set bucket quota", func(t *testing.T) {
		err := co.SetIndividualBucketQuota(context.Background(), IndividualBucketQuotaSpec{UID: "admin", Bucket: suite.bucketTestName})
		assert.NoError(suite.T(), err)
	})
}
