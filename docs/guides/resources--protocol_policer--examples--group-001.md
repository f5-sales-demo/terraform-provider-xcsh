---
page_title: "xcsh_protocol_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer examples."
---

# xcsh_protocol_policer examples

<a id="canonical-3120323230231032-0233303201201322-3001301212102012-1331313222310222-2110131022000020-0021030121331322-0212120110332220-1100023333311320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- Examples

<a id="canonical-1301222233323000-1021222100020131-0303300320100121-0211312021223010-0011022001102320-2212211303010120-3302010303322101-1133132220311132"></a>

### Complete configurations for `xcsh_protocol_policer`

- [Resource](resources--protocol_policer--examples--group-001.md#canonical-2122230011130102-0023302203312021-0032100333210210-2203213210011030-0222321032023131-1321022121233233-0113001333120303-1212311031212332): valid configuration.

<a id="canonical-2122230011130102-0023302203312021-0032100333210210-2203213210011030-0222321032023131-1321022121233233-0113001333120303-1212311031212332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Examples](resources--protocol_policer--examples--group-001.md#canonical-3120323230231032-0233303201201322-3001301212102012-1331313222310222-2110131022000020-0021030121331322-0212120110332220-1100023333311320)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_policer/resource.tf`; digest `sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c`.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```
