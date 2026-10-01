---
page_title: "xcsh_third_party_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_third_party_application examples."
---

# xcsh_third_party_application examples

<a id="canonical-1846a0b2efe35f7b3461c3775766d5787e6856ac42d7df49b74a9afbde5e3993"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97f3d8a294980c88e8f69707a28d8a65fe1b6b1c529f92b2675344fa5bb9dd57"></a>

## Examples — Examples / 38e67d9dfa9f / 2

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-b8aad6ca3bfaa99ef05f861b2a172f6bea0697789bde7db18a81450cf6385d78)
- Examples

<a id="canonical-80f0191770f9f404be4d924ada5d28cb864f6cd32529bb78daedf0f1f605fbc2"></a>

## Complete configurations — Examples / 38e67d9dfa9f / 3

- [Data source](data-sources--third_party_application--examples--group-001.md#canonical-2047086f84efb82e729cdd420302d5c79ed06a9982b981530a3c6ded5a63ca92): valid configuration.

<a id="canonical-ce8683a4ef9fef4e5239a4fae8e27f5c9b35317dea2f2d20e18200761a33dd41"></a>

## Next pages — Examples / 38e67d9dfa9f / 4

- [Data source](data-sources--third_party_application--examples--group-001.md#canonical-2047086f84efb82e729cdd420302d5c79ed06a9982b981530a3c6ded5a63ca92)
- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-b8aad6ca3bfaa99ef05f861b2a172f6bea0697789bde7db18a81450cf6385d78)

<a id="canonical-2047086f84efb82e729cdd420302d5c79ed06a9982b981530a3c6ded5a63ca92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a387f666645fc4dcd114aa03b451aa0a668df438dfc51dd4cfb96781ec58f3"></a>

## Data source — Data source / fbdf9ed6787f / 2

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-b8aad6ca3bfaa99ef05f861b2a172f6bea0697789bde7db18a81450cf6385d78)
- [Examples](data-sources--third_party_application--examples--group-001.md#canonical-1846a0b2efe35f7b3461c3775766d5787e6856ac42d7df49b74a9afbde5e3993)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_third_party_application/data-source.tf`; digest `sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53`.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```

<a id="canonical-041fb42461659709e55a5779b539021578bd366646c9d7531eb3ae928b79dcb7"></a>

## Next pages — Data source / fbdf9ed6787f / 3

- [Examples](data-sources--third_party_application--examples--group-001.md#canonical-1846a0b2efe35f7b3461c3775766d5787e6856ac42d7df49b74a9afbde5e3993)
- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-b8aad6ca3bfaa99ef05f861b2a172f6bea0697789bde7db18a81450cf6385d78)
