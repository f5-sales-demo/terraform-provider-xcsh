---
page_title: "xcsh_forwarding_class landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class landing."
---

# xcsh_forwarding_class landing

<a id="canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-524477b66dc5153c643fd817f40286244ddd03c6dc25260ee51864946a5a26a5"></a>

## xcsh_forwarding_class — xcsh_forwarding_class / 690d8dc10edc / 2

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

<a id="canonical-cc881c3c552e119cb33b5285272c728b19e3d0277322c3e11ff2c72fd9f7cf6c"></a>

## Prerequisites — xcsh_forwarding_class / 690d8dc10edc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-687ea6c11c0e90c0a05570e2a1e188dee4bd1f1e631a9179ca6cd341d59c84fe"></a>

## Minimal configuration — xcsh_forwarding_class / 690d8dc10edc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

<a id="canonical-607bdbac013278945a3f016cfabe1a6783f71ba62acc57f4a89cb7b4fcb78d21"></a>

## Root configuration — xcsh_forwarding_class / 690d8dc10edc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-747d0a02131ac87c5bdfac14aecd2640aa13cdaac289c1936974f505bcec5c2a"></a>

## Next pages — xcsh_forwarding_class / 690d8dc10edc / 6

- [Property reference](../guides/data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [Examples](../guides/data-sources--forwarding_class--examples--group-001.md#canonical-75575c16923f49aa304e0a71ec8c3ee586191de8d37a0f4e564b37791041c33e)
