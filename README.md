# AppDynamics REST API
  
Library to access the [AppDynamics REST API](https://docs.appdynamics.com/appd/23.x/latest/en/extend-appdynamics/appdynamics-apis)  

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


Based on work of https://github.com/dlopes7/go-appdynamics-rest-api

Provides Go language REST API for AppDynamics


## License ##

MIT License

Copyright (c) 2023 David Lopes
Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
