---
page_title: "xcsh_securemesh_site_v2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 landing."
---

# xcsh_securemesh_site_v2 landing

<a id="canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e78fe7aff62871cb60674227d4fedd6befdec11588c2364a5ef2a2118a965e13"></a>

## xcsh_securemesh_site_v2 — xcsh_securemesh_site_v2 / baa760b82814 / 2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

<a id="canonical-0ad282e961286728e8e2d0457a13ec6cdd5e3d4c0807995802a67b281c0e0555"></a>

## Prerequisites — xcsh_securemesh_site_v2 / baa760b82814 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ca462458d2277307a50220f7516ffddf4bcc0bc59967cff60b9f7cc445a12aa6"></a>

## Minimal configuration — xcsh_securemesh_site_v2 / baa760b82814 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```

<a id="canonical-d322362c64f027d96676cf5e411bf8ec7e945a6fa392aff9998ba9e6667ecd07"></a>

## Root configuration — xcsh_securemesh_site_v2 / baa760b82814 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-ab15355bd88bce750c569ca13b89cae111fbf8c619d3214ef929758f3921fbde"></a>

## Next pages — xcsh_securemesh_site_v2 / baa760b82814 / 6

- [Property reference](../guides/data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [Examples](../guides/data-sources--securemesh_site_v2--examples--group-001.md#canonical-dd1bd31a95e67d7324b79547dcf212c3ece23b79f1d8be24343b53effe8e03f7)
