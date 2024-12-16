package projectError

import "errors"

var (
	ErrorNoData              = errors.New("can't find data")
	ErrorStartingTransaction = errors.New("can't start transaction")
	ErrorCommitTransaction   = errors.New("can't commit transaction")
	ErrorUserNotFound        = errors.New("user not found")
	ErrorCantUpdateUser      = errors.New("can't change user")
	ErrorUserExist           = errors.New("user exist")
	ErrorCantGetUser         = errors.New("can't get user")
	ErrorCantCreateUser      = errors.New("can't create user")
	ErrorCantDeleteUser      = errors.New("can't delete crypto")
	ErrorCryptoExist         = errors.New("crypto exist")
	ErrorCryptoNotFound      = errors.New("crypto not found")
	ErrorCantCreateCrypto    = errors.New("can't create crypto")
	ErrorCantUpdateCrypto    = errors.New("can't update crypto")
	ErrorCantDeleteCrypto    = errors.New("can't delete crypto")
	ErrorCantCreateLogger    = errors.New("can't  create logger")
)
