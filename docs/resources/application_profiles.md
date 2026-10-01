---
page_title: "xcsh_application_profiles landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles landing."
---

# xcsh_application_profiles landing

<a id="canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4695f0cfcf157e0fb08f37fb89c67266f8939146b3f8b0866f82c47a88197fcc"></a>

## xcsh_application_profiles — xcsh_application_profiles / d6026c0b9b0f / 2

Breadcrumbs:

- xcsh_application_profiles

Manages Application Profiles in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-63501d372c057a89b7e04bc8b704ccac83ff98b7f9931397645fa36b032214de"></a>

## Prerequisites — xcsh_application_profiles / d6026c0b9b0f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0aabb761ff090b562aa682f50fd0fe002e26471b91d4f595e61dadb46bd6a681"></a>

## Minimal configuration — xcsh_application_profiles / d6026c0b9b0f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

<a id="canonical-12f40742f156243883267acc9ad6c81b784386a60bf89ce9f009b4847859d859"></a>

## Root configuration — xcsh_application_profiles / d6026c0b9b0f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-019bb8e1d0c2a57cd21e61bf1b1e991606975bd0b2ee24c55a32b2fb4516f9fd"></a>

## Next pages — xcsh_application_profiles / d6026c0b9b0f / 6

- [Property reference](../guides/resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [Examples](../guides/resources--application_profiles--examples--group-001.md#canonical-1997211f5f6bcf28b790e55429796b98903d009d88f6006d82ba263143c7649a)
- [Import](../guides/resources--application_profiles--lifecycle--group-001.md#canonical-338bb23e9d5cc47da4db36e3163347f188d4547c00c02410929f782b8d32418c)
- [Timeouts](../guides/resources--application_profiles--lifecycle--group-001.md#canonical-130f2b55446c962cc47192dbc460671ebc4f68ee4c3011c4847e977980096f0d)
