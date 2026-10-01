---
page_title: "xcsh_public_ip landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip landing."
---

# xcsh_public_ip landing

<a id="canonical-4d58cc44b5d0834da51e718a5f01b7b87916d0e39e4050de470be23a21d6e26a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bbb91ecd0d6be90bd26423371d214af545e5e7f591f867b6df23dd68c3f3181"></a>

## xcsh_public_ip — xcsh_public_ip / ac2c512d6506 / 2

Breadcrumbs:

- xcsh_public_ip

Manages a Public IP resource in F5 Distributed Cloud for get public\_ip will get the object from the
storage backend for namespace metadata.namespace. configuration. (read-only data source)

<a id="canonical-7a05212473c34741f31fc836bc31ad38a05c29c551334dfaf7a51f7e36cb8b6d"></a>

## Prerequisites — xcsh_public_ip / ac2c512d6506 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-aa7d87cf27fb2717d26caea69d4a7f717f7b8331510d6270e4ddbda5067764c3"></a>

## Minimal configuration — xcsh_public_ip / ac2c512d6506 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```

<a id="canonical-76c5fbf3df92c5bb8b5c543b23242bbdf9c98e0fcc21794247e3bb213cac2a07"></a>

## Root configuration — xcsh_public_ip / ac2c512d6506 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-74384cc713deaf8c0ad829c2716a25317bd25ba2fe4be1a5dc550f456829962f"></a>

## Next pages — xcsh_public_ip / ac2c512d6506 / 6

- [Property reference](../guides/data-sources--public_ip--reference--group-001.md#canonical-78f60b2c54e55db478e018b4fa9587caff6345ca30632666a4315d5669f9c5ec)
- [Examples](../guides/data-sources--public_ip--examples--group-001.md#canonical-66ff80522379979087e1b46893f2f83bfdb92352b04293cec0b8fa51c4d25c3f)
