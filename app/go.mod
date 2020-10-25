module github.com/Fukkatsuso/oauth-sample

go 1.14

replace github.com/Fukkatsuso/oauth-sample/app/lib/twitter => ./lib/twitter

require (
	github.com/Fukkatsuso/oauth-sample/app/lib/twitter v0.0.0-00010101000000-000000000000
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/contrib v0.0.0-20201005132743-ca038bbf2944
	github.com/gin-gonic/gin v1.6.3
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kidstuff/mongostore v0.0.0-20181113001930-e650cd85ee4b // indirect
	github.com/kimiazhu/ginweb-contrib v0.0.0-20160408072953-eb842df46c92
	github.com/kimiazhu/golib v0.0.0-20181123075453-970a1a889a1d // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mrjones/oauth v0.0.0-20190623134757-126b35219450
	github.com/ugorji/go v1.1.13 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/sys v0.0.0-20201020230747-6e5568b54d1a // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
