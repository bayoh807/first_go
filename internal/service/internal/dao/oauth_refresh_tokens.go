// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_0620/internal/service/internal/dao/internal"
)

// oauthRefreshTokensDao is the data access object for table oauth_refresh_tokens.
// You can define custom methods on it to extend its functionality as you wish.
type oauthRefreshTokensDao struct {
	*internal.OauthRefreshTokensDao
}

var (
	// OauthRefreshTokens is globally public accessible object for table oauth_refresh_tokens operations.
	OauthRefreshTokens = oauthRefreshTokensDao{
		internal.NewOauthRefreshTokensDao(),
	}
)

// Fill with you ideas below.