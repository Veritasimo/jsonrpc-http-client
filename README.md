jsonrpc-http-client
===================

A simple JSON-RPC client for Go.

Usage Example
-----------

    import (
      ...
      "log"
      "github.com/Veritasimo/jsonrpc-http-client"
      ...
    )
    
    ...

    proxy := jsonrpc.NewProxy("http://localhost:8080/service/", "myapp") // Prefixes all methods with myapp.
    var payload = map[string]interface{}{
      "Cake": "Delicious",
      "Orly": true,
    }
    
    resp, err := proxy.Call("cakereport", payload)
    
    if err != nil {
      log.Fatal(err)
    }
    
    respPayload, ok := resp.(map[string]interface{})
    /* Per the JSON-RPC specification the response should be an object that looks something like this:
    {
      "id": 1,
      "error:  null,
      "result": true
    }
    
    The error field will not be null if a server error occurs. These are not caught by the Call error
    handling as you may need to handle it yourself. /*
    if !ok {
      log.Fatal("Terrible server.")
    }
    
    result, ok := respPaylad["result"].(bool)
    if !ok {
      log.Fatal("We were expecting a bool result!")
    }
    
    log.Printf("%+v", respPayload)
