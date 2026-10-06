---
page_title: "xcsh_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer examples."
---

# xcsh_policer examples

<a id="canonical-1311202001001303-0121132302210321-3212302221232023-2331311030121133-2003102010101310-2221203033211131-3300210210103331-2100210312012321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)
- Examples

<a id="canonical-0111011323211311-0330212111132111-3110321023233012-0213000310302303-3223322002021210-3301302011123112-1023102112131221-0002211210222023"></a>

### Complete configurations for `xcsh_policer`

- [Resource](resources--policer--examples--group-001.md#canonical-1112112102112202-1312020331023111-1012320113033311-0102113031110021-0121131031221111-0133031133130310-2102031120132310-0323020111332102): valid configuration.

<a id="canonical-1112112102112202-1312020331023111-1012320113033311-0102113031110021-0121131031221111-0133031133130310-2102031120132310-0323020111332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)
- [Examples](resources--policer--examples--group-001.md#canonical-1311202001001303-0121132302210321-3212302221232023-2331311030121133-2003102010101310-2221203033211131-3300210210103331-2100210312012321)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policer/resource.tf`; digest `sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443`.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```
