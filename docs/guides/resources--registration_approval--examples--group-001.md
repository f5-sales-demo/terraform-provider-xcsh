---
page_title: "xcsh_registration_approval examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration_approval examples."
---

# xcsh_registration_approval examples

<a id="canonical-0120212003331230-0231120030031133-2330122110012231-1321111321200023-2022120032210303-3130002001000100-2300132013030121-2023021330011321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-1331103323121310-0002130210000313-2020131122020200-0233121113323131-3323011213331230-0033000123022130-3131333113210323-3013000301112322)
- Examples

<a id="canonical-1023130201203202-0022313223101000-2320322212002102-1312312310120222-1231113203000123-3310003213223220-3301133013132303-0300330223211232"></a>

### Complete configurations for `xcsh_registration_approval`

- [Resource](resources--registration_approval--examples--group-001.md#canonical-1031110231221333-0203300210212111-0223331033100203-1100320031333333-1222031031033003-3033130012233322-0223223213003023-0132133330322103): valid configuration.

<a id="canonical-1031110231221333-0203300210212111-0223331033100203-1100320031333333-1222031031033003-3033130012233322-0223223213003023-0132133330322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-1331103323121310-0002130210000313-2020131122020200-0233121113323131-3323011213331230-0033000123022130-3131333113210323-3013000301112322)
- [Examples](resources--registration_approval--examples--group-001.md#canonical-0120212003331230-0231120030031133-2330122110012231-1321111321200023-2022120032210303-3130002001000100-2300132013030121-2023021330011321)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration_approval/resource.tf`; digest `sha256:5a69b599688eff851094350bc6a6e7fffd81e09b55087a523abb0c8c970908cd`.

```terraform
# RegistrationApproval Resource Example
# Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RegistrationApproval configuration
resource "xcsh_registration_approval" "example" {
  name      = "example-registration-approval"
  namespace = "staging"

  cluster_size = 1
}
```
