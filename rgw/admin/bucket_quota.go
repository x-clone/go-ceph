//go:build ceph_preview
// +build ceph_preview

package admin

import (
	"context"
	"net/http"
)

// IndividualBucketQuotaSpec describes an object store quota for an individual bucket
type IndividualBucketQuotaSpec struct {
	UID        string `json:"user_id" url:"uid"`
	Bucket     string `json:"bucket" url:"bucket"`
	QuotaType  string `url:"quota-type"`
	Enabled    *bool  `json:"enabled" url:"enabled"`
	CheckOnRaw bool   `json:"check_on_raw"`
	MaxSize    *int64 `json:"max_size" url:"max-size"`
	MaxSizeKb  *int   `json:"max_size_kb" url:"max-size-kb"`
	MaxObjects *int64 `json:"max_objects" url:"max-objects"`
}

// SetIndividualBucketQuota sets quota to a specific bucket
// https://docs.ceph.com/en/latest/radosgw/adminops/#set-quota-for-an-individual-bucket
//  PREVIEW
func (api *API) SetIndividualBucketQuota(ctx context.Context, quota IndividualBucketQuotaSpec) error {
	// Always for quota type to bucket
	quota.QuotaType = "bucket"

	if quota.UID == "" {
		return errMissingUserID
	}

	if quota.Bucket == "" {
		return errMissingBucket
	}

	_, err := api.call(ctx, http.MethodPut, "/bucket?quota", valueToURLParams(quota))
	if err != nil {
		return err
	}

	return nil
}
