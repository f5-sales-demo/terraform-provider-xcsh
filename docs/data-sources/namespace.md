---
page_title: "xcsh_namespace landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace landing."
---

# xcsh_namespace landing

<a id="canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b941d55d1b44420a39f985d214db8ddd67c56467acb9dfd8a93e1706f895931"></a>

## xcsh_namespace — xcsh_namespace / aaeb769cadc6 / 2

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.

<a id="canonical-dbb4b093853033107dd97707dedce4c9138bae27b5153474630b1e0371c0a2ba"></a>

## Prerequisites — xcsh_namespace / aaeb769cadc6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8d3fe21097d97ce053d6577118c3d663d9ad2775edf70d7f88e046c3a2f46695"></a>

## Minimal configuration — xcsh_namespace / aaeb769cadc6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

<a id="canonical-1b7ddda40dbd5fb92b41e5b4450de23800e7a3f13ae53c74e0f8a028c5955d94"></a>

## Root configuration — xcsh_namespace / aaeb769cadc6 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3b0c46e14502279701f820a5421bcc79c29449fd9859445a780e8413803baf72"></a>

## Next pages — xcsh_namespace / aaeb769cadc6 / 6

- [Property reference](../guides/data-sources--namespace--reference--group-001.md#canonical-6f4493fe6a03e423b14ed16b610b1d61c98ab9ef17f2f98fa04f2fb20e9e9fd7)
- [Examples](../guides/data-sources--namespace--examples--group-001.md#canonical-13126340d43a7fef70e3c42c1e8d7b7fb39680618a33f6e79757825b1db49b7e)
