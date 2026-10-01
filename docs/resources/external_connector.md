---
page_title: "xcsh_external_connector landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector landing."
---

# xcsh_external_connector landing

<a id="canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd9780b06f9370cfc23a3d14f86a62c3f1608a4b143bced8a1f84bf5737b52bd"></a>

## xcsh_external_connector — xcsh_external_connector / b8246874aec2 / 2

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

<a id="canonical-ca57f6533bac313ec2331f9aeb918e62e199a1a495d912818e685499e8d25e69"></a>

## Prerequisites — xcsh_external_connector / b8246874aec2 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2a567b70a181c7bd5ce58fcf1d17c2878a9a8924742f68ba1b80c9994096e018"></a>

## Minimal configuration — xcsh_external_connector / b8246874aec2 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```

<a id="canonical-cc6c517a4cd88204c0109eb350fe072e482cf61977ddeec32a309b74d21d7a9b"></a>

## Root configuration — xcsh_external_connector / b8246874aec2 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-43d0c73f1ad21d88afc376604f55d7975fdb7066baf7ee1da3389f50a11ba270"></a>

## Next pages — xcsh_external_connector / b8246874aec2 / 6

- [Property reference](../guides/resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [Examples](../guides/resources--external_connector--examples--group-001.md#canonical-4b2095e73af6f07a0640ac9c476f2bec2e3a1067e307b42dfc92f1a141556abf)
- [Import](../guides/resources--external_connector--lifecycle--group-001.md#canonical-d18f2387ebcbaeea866ff767155d6b9e69a9544e7e6ba800bc2b7b3379317938)
- [Timeouts](../guides/resources--external_connector--lifecycle--group-001.md#canonical-35d480a115e2886c6985bdbbc2cbd9a1189c0b4a0baa190a99de15229dd5ab34)
