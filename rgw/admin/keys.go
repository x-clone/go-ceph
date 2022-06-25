//go:build ceph_preview
// +build ceph_preview

package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// KeySpec is the user credential management configuration
// https://docs.ceph.com/en/latest/radosgw/adminops/#id37
type KeySpec struct {
	UID         string `url:"uid"`          // The user ID to receive the new key
	SubUser     string `url:"subuser"`      // The subuser ID to receive the new key
	KeyType     string `url:"key-type"`     // Key type to be generated, options are: swift, s3 (default)
	AccessKey   string `url:"access-key"`   // Specify the access key
	SecretKey   string `url:"secret-key"`   // Specify secret key
	GenerateKey *bool  `url:"generate-key"` // Generate a new key pair and add to the existing keyring
}

// CreateKey will generate new keys or add specified to keyring
// https://docs.ceph.com/en/latest/radosgw/adminops/#create-key
// PREVIEW
func (api *API) CreateKey(ctx context.Context, key KeySpec) (*[]UserKeySpec, error) {
	if key.UID == "" && key.SubUser == "" {
		return nil, errMissingUserID
	}

	body, err := api.call(ctx, http.MethodPut, "/user?key", valueToURLParams(key))
	if err != nil {
		return nil, err
	}

	ref := []UserKeySpec{}
	err = json.Unmarshal(body, &ref)
	if err != nil {
		return nil, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return &ref, nil
}

// RemoveKey will remove an existing key
// https://docs.ceph.com/en/latest/radosgw/adminops/#remove-key
// KeySpec.SecretKey parameter shouldn't be provided and will be ignored
// PREVIEW
func (api *API) RemoveKey(ctx context.Context, key KeySpec) error {
	if key.UID == "" && key.SubUser == "" {
		return errMissingUserID
	}
	if key.UID != "" && key.AccessKey == "" {
		return errMissingUserAccessKey
	}

	_, err := api.call(ctx, http.MethodDelete, "/user?key", valueToURLParams(key))
	if err != nil {
		return err
	}

	return nil
}
