module gin-my

go 1.16

// FIXME:这都留的啥？没用就删掉.
require (
	gin v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	gopkg.in/ini.v1 v1.62.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.21.2
)

replace gin v0.0.0 => ./gin
