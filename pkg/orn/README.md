# Object Resource Names (ORNs)

Object Resource Names (ORNs) uniquely identify Object resources.

## **ORN Format**

The following are the general formats for ORNs; the specific components and values used depend on the Application usage.

```
orn:partition:service:identity-id:resource
orn:partition:service:identity-id:resourcetype/resource
```

### Partition
The partition that the resource is in. For example, the partition for resources in the Compnay can be "acme".

### Service
The service that identifies the service (e.g. ContactDirectory, Judge, etc.).

### Identity

The ID of the identity that owns the resource. For example, 7bbcc039-6248-4a19-a664-51ebb61c8950.

> Note that the ORNs for some resources don't require an identity ID, so this component might be omitted.

### Resource, ResourceType/Resource

The content of this part of the ORN varies by service. It often includes an indicator of the type of resource (for example, an IAM identity) followed by a slash (/), followed by the resource name itself. Some service allows paths for resource names, as described in Paths in ORNs.

## **Paths in ORNs**

TODO: Define the ORN Path Spec properly.


## Package

The orn package provides provides some helpful methods to deal with the ORN format.

- Struct Go to represent the string
- Unmarshal/Marshal functions

See it in action:

```go
package yours

import (
  "github.com/gearnode/judge/orn"
  "fmt"
)

func YourFunction() {
  data := "orn:"foo-company:eatService::food:tomato"
  orn := orn.ORN{}
  if err := orn.Unmarshal(data, &orn); err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println(orn.Marshal(orn))
}
```
