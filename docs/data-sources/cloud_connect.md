---
page_title: "xcsh_cloud_connect landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect landing."
---

# xcsh_cloud_connect landing

<a id="canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6436ef88d3b4287b48a5279af233504354a44daf7ad3981bc8d18aeb9cc555d1"></a>

## xcsh_cloud_connect — xcsh_cloud_connect / 7944851ba996 / 2

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

<a id="canonical-561b8caea293392debde534bbb49b2e57640c4e0dd8e247db900777bb375fad7"></a>

## Prerequisites — xcsh_cloud_connect / 7944851ba996 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-7849e14d1cbbd3488cd0ed1be6ee027097eefe96ab9a8c5b87b062da729dc577"></a>

## Minimal configuration — xcsh_cloud_connect / 7944851ba996 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudConnect Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudConnect by name
data "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}

output "cloud_connect_id" {
  value = data.xcsh_cloud_connect.example.id
}
```

<a id="canonical-7e30ea594741aef3a306b03811f0e6901f141a54b85b77a83a3fa1d2f98b5259"></a>

## Root configuration — xcsh_cloud_connect / 7944851ba996 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-68d10b9c8e143ce6d1c1828b68f65d9d34da0eaa48ce9a2305a4ceb4b5dfe915"></a>

## Next pages — xcsh_cloud_connect / 7944851ba996 / 6

- [Property reference](../guides/data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [Examples](../guides/data-sources--cloud_connect--examples--group-001.md#canonical-0e447ba11f2ec3291f89ffa097e0f893c0f5cb51189a0746945199aa4d3c314f)
