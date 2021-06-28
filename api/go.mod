module ezyo/forum

go 1.13

require (
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/febytanzil/gomailer v0.8.4
	github.com/joho/godotenv v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	gopkg.in/mail.v2 v2.3.1
	gorm.io/gorm v1.21.9
	goyave.dev/goyave/v3 v3.9.1
)

replace goyave.dev/goyave/v3 => github.com/ezyostudio/goyave/v3 v3.10.0
