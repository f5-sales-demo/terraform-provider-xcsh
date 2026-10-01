---
page_title: "xcsh_tenant_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration examples."
---

# xcsh_tenant_configuration examples

<a id="canonical-ef992db0544684039f98ad0b10918386ea3ddd00f9017e4dffa394b0ee6a5569"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3adfb88a76a401f2103ae9a47547ed5384722d03528b2dc206874c42bae8f5a9"></a>

## Examples — Examples / 9f4a9c53059b / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-6c65d7df25dc79a7517bd40ba3201b123141e0bec2b95c1ab2d50ec783eb0521)
- Examples

<a id="canonical-c556720cfb92555a7fb64c66cd32d08e6d0e58cd4a752e57af0340577dd7044e"></a>

## Complete configurations — Examples / 9f4a9c53059b / 3

- [Data source](data-sources--tenant_configuration--examples--group-001.md#canonical-adcd8fb81ec2bf787b0619bddb2eba90fb507810a2ae08fb845d2fee4b723ede): valid configuration.

<a id="canonical-6f7702333c7a14d41495b7b20d8f3510e60aa143fc2b3e55d4944e6421af938a"></a>

## Next pages — Examples / 9f4a9c53059b / 4

- [Data source](data-sources--tenant_configuration--examples--group-001.md#canonical-adcd8fb81ec2bf787b0619bddb2eba90fb507810a2ae08fb845d2fee4b723ede)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-6c65d7df25dc79a7517bd40ba3201b123141e0bec2b95c1ab2d50ec783eb0521)

<a id="canonical-adcd8fb81ec2bf787b0619bddb2eba90fb507810a2ae08fb845d2fee4b723ede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99c8093ccfc999fc15766b94fa1bc5de42bb8e812b3e023f542e5c28f2e66774"></a>

## Data source — Data source / 3da172cb0994 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-6c65d7df25dc79a7517bd40ba3201b123141e0bec2b95c1ab2d50ec783eb0521)
- [Examples](data-sources--tenant_configuration--examples--group-001.md#canonical-ef992db0544684039f98ad0b10918386ea3ddd00f9017e4dffa394b0ee6a5569)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tenant_configuration/data-source.tf`; digest `sha256:82e6c50873cbaa19717f3b339ef9fae150aab727845b88eb85553bf86a8cc317`.

```terraform
# TenantConfiguration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TenantConfiguration by name
data "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}

output "tenant_configuration_id" {
  value = data.xcsh_tenant_configuration.example.id
}
```

<a id="canonical-bc764eb652db88c358fe86afba0f6ae6f952f7820710b872dccffd6546bda7bc"></a>

## Next pages — Data source / 3da172cb0994 / 3

- [Examples](data-sources--tenant_configuration--examples--group-001.md#canonical-ef992db0544684039f98ad0b10918386ea3ddd00f9017e4dffa394b0ee6a5569)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-6c65d7df25dc79a7517bd40ba3201b123141e0bec2b95c1ab2d50ec783eb0521)
