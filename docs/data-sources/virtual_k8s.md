---
page_title: "xcsh_virtual_k8s landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s landing."
---

# xcsh_virtual_k8s landing

<a id="canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66083e5d61fc69364dea5dd2574c5932c2321c060ec9337b64293ba04c2ebe1f"></a>

## xcsh_virtual_k8s — xcsh_virtual_k8s / 47c10fb865b0 / 2

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0a810d6c5f871ed9e087166f0203d00d59f552bff8c486c7d0a1d30f161e3555"></a>

## Prerequisites — xcsh_virtual_k8s / 47c10fb865b0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

<a id="canonical-f1151e259eb409ae83c2d41c39e86789893708cb50556128378b846d9b42e3c1"></a>

## Minimal configuration — xcsh_virtual_k8s / 47c10fb865b0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

<a id="canonical-ce4b9a6d4c3a87370a11f35a3f55f811ba6ce44a08dca28f1ecfd6dd2a2389b7"></a>

## Root configuration — xcsh_virtual_k8s / 47c10fb865b0 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9ee4924caafa5d1a3a5bd3b7036e64e738d21d4a854682f8093d134daa914e51"></a>

## Next pages — xcsh_virtual_k8s / 47c10fb865b0 / 6

- [Property reference](../guides/data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- [Examples](../guides/data-sources--virtual_k8s--examples--group-001.md#canonical-c39a52aefa52ab272de791fa63cb96987fe1981f36e2b0473eae32b77dc43ad8)
