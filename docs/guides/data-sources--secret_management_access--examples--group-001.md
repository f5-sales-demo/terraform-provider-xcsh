---
page_title: "xcsh_secret_management_access examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access examples."
---

# xcsh_secret_management_access examples

<a id="canonical-af95116420ee9e87d5fe5c3f46181f8ef7aa75cc239ecdbe6971e0be1537c201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6557415ab408667557c834e1bdae312bfcb603e1bc397882899c3e49b58aba3"></a>

## Examples — Examples / 7b2b202c9b81 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- Examples

<a id="canonical-225cb2c6a6862072aee30c604a08e62166cf9c753714b5f51070b788902445c8"></a>

## Complete configurations — Examples / 7b2b202c9b81 / 3

- [Data source](data-sources--secret_management_access--examples--group-001.md#canonical-f291fec2403568de8661d649e45b8464c3449ef4e6b00e00b86de1ad11421fec): valid configuration.

<a id="canonical-de639d46c99f6e20671bbb1c16f7d417cd66f632520e55497a603ecf2d290e24"></a>

## Next pages — Examples / 7b2b202c9b81 / 4

- [Data source](data-sources--secret_management_access--examples--group-001.md#canonical-f291fec2403568de8661d649e45b8464c3449ef4e6b00e00b86de1ad11421fec)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-f291fec2403568de8661d649e45b8464c3449ef4e6b00e00b86de1ad11421fec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16b6d4135600f335e69c41336d9f5a83619e45c7628b59dc0e14245b04c6df2f"></a>

## Data source — Data source / 7af25a25c185 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Examples](data-sources--secret_management_access--examples--group-001.md#canonical-af95116420ee9e87d5fe5c3f46181f8ef7aa75cc239ecdbe6971e0be1537c201)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_secret_management_access/data-source.tf`; digest `sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2`.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```

<a id="canonical-27a16372524d02f51f202563d876b2253bd5a711e72800bcbd8c25aca8b9e51f"></a>

## Next pages — Data source / 7af25a25c185 / 3

- [Examples](data-sources--secret_management_access--examples--group-001.md#canonical-af95116420ee9e87d5fe5c3f46181f8ef7aa75cc239ecdbe6971e0be1537c201)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
