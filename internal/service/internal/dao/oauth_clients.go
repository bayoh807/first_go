// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_0620/internal/service/internal/dao/internal"
)

// oauthClientsDao is the data access object for table oauth_clients.
// You can define custom methods on it to extend its functionality as you wish.
type oauthClientsDao struct {
	*internal.OauthClientsDao
}

var (
	// OauthClients is globally public accessible object for table oauth_clients operations.
	OauthClients = oauthClientsDao{
		internal.NewOauthClientsDao(),
	}
)

// Fill with you ideas below.