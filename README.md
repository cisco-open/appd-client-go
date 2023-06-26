# AppDynamics Go Client
  
This project provides access to the [AppDynamics REST API](https://docs.appdynamics.com/appd/23.x/latest/en/extend-appdynamics/appdynamics-apis) in Go language.

## Usage ##
```go
import "github.com/cisco-open/appd-client-go"
```

Create a client, get every Business Transaction for every Application

```go
client, _ :=  appdrest.NewClient("http", "192.168.33.10", 8090, "admin", "password", "customer1")

apps, err := client.Application.GetApplications()
	if err != nil {
		panic(err.Error())
	}
	for _, app := range apps {

		bts, err := client.BusinessTransaction.GetBusinessTransactions(app.ID)
		if err != nil {
			panic(err.Error())
		}
		for _, bt := range bts {
			fmt.Printf("App: %s, Tier: %s, BT: %s, Type: %s\n", app.Name, bt.TierName, bt.Name, bt.EntryPointType)
		}

	}
```


## Projects using this library ##

* [AppDynamics Swagger Tool](https://github.com/cisco-open/swagger-appd-tool)

## Support ##

We welcome feedback, questions, issue reports and pull requests.

Maintainer's email address: mdivis@cisco.com

[GitHub Issues](https://github.com/cisco-open/appd-client-go/issues)

## Acknowledgements ##

Based on work of https://github.com/dlopes7/go-appdynamics-rest-api

## License ##

MIT License


Copyright (c) 2023 David Lopes<br>
Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
