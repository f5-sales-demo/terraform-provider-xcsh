---
page_title: "xcsh_app_api_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group landing."
---

# xcsh_app_api_group landing

<a id="canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9d6035b2261df54c9ea7dd455ec978ee5382e20a3700f422574b397a0bebb28"></a>

## xcsh_app_api_group — xcsh_app_api_group / 961556176746 / 2

Breadcrumbs:

- xcsh_app_api_group

Manages app\_api\_group creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-5a4245defa836e26a5f81f9d28310647515fbde5d7848872890fed6b3efc2c02"></a>

## Prerequisites — xcsh_app_api_group / 961556176746 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a3a13df6d1945ffefa0f9c1a0a4997453727bd977464e1af892346d368b42309"></a>

## Minimal configuration — xcsh_app_api_group / 961556176746 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

<a id="canonical-52cf32aef251e2da67dd42586ea04f19bfa5c99c10a25ed4671c0ce4e5366534"></a>

## Root configuration — xcsh_app_api_group / 961556176746 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b407d0375cfaea983feddee8780ebbd5fbab75708c47345f4a6676d9b8ccb03c"></a>

## Next pages — xcsh_app_api_group / 961556176746 / 6

- [Property reference](../guides/data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [Examples](../guides/data-sources--app_api_group--examples--group-001.md#canonical-3241f436996013f20af4a7b6f436914d3e43c74b62cf9d26ba502193fd04cc4e)
