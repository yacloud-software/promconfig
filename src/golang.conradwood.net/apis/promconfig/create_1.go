// client create: PromConfigServiceClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/promconfig/promconfig.proto
   gopackage : golang.conradwood.net/apis/promconfig
   importname: ai_0
   clientfunc: GetPromConfigService
   serverfunc: NewPromConfigService
   lookupfunc: PromConfigServiceLookupID
   varname   : client_PromConfigServiceClient_0
   clientname: PromConfigServiceClient
   servername: PromConfigServiceServer
   gsvcname  : promconfig.PromConfigService
   lockname  : lock_PromConfigServiceClient_0
   activename: active_PromConfigServiceClient_0
*/

package promconfig

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_PromConfigServiceClient_0 sync.Mutex
  client_PromConfigServiceClient_0 PromConfigServiceClient
)

func GetPromConfigClient() PromConfigServiceClient { 
    if client_PromConfigServiceClient_0 != nil {
        return client_PromConfigServiceClient_0
    }

    lock_PromConfigServiceClient_0.Lock() 
    if client_PromConfigServiceClient_0 != nil {
       lock_PromConfigServiceClient_0.Unlock()
       return client_PromConfigServiceClient_0
    }

    client_PromConfigServiceClient_0 = NewPromConfigServiceClient(client.Connect(PromConfigServiceLookupID()))
    lock_PromConfigServiceClient_0.Unlock()
    return client_PromConfigServiceClient_0
}

func GetPromConfigServiceClient() PromConfigServiceClient { 
    if client_PromConfigServiceClient_0 != nil {
        return client_PromConfigServiceClient_0
    }

    lock_PromConfigServiceClient_0.Lock() 
    if client_PromConfigServiceClient_0 != nil {
       lock_PromConfigServiceClient_0.Unlock()
       return client_PromConfigServiceClient_0
    }

    client_PromConfigServiceClient_0 = NewPromConfigServiceClient(client.Connect(PromConfigServiceLookupID()))
    lock_PromConfigServiceClient_0.Unlock()
    return client_PromConfigServiceClient_0
}

func PromConfigServiceLookupID() string { return "promconfig.PromConfigService" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("promconfig.PromConfigService")
   AddService("promconfig.PromConfigService")
}





