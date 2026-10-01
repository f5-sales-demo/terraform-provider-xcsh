---
page_title: "xcsh_cloud_credentials landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials landing."
---

# xcsh_cloud_credentials landing

<a id="canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c8fd02e843f5debf7ce12ccc8c4c07c084dbd26f045a71d03fa2c4ca6d733eb"></a>

## xcsh_cloud_credentials — xcsh_cloud_credentials / e994caf30cb8 / 2

Breadcrumbs:

- xcsh_cloud_credentials

Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud\_credentials
object. configuration.

<a id="canonical-6465d16e4706da82e1f267c79bbb7fb8e049a50de2b08a3424cd7dc75d15fd86"></a>

## Prerequisites — xcsh_cloud_credentials / e994caf30cb8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2568c05b70c36681187639a95eb56977cb6c67d054ffeb4b310c309b7a82021e"></a>

## Minimal configuration — xcsh_cloud_credentials / e994caf30cb8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudCredentials Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudCredentials by name
data "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}

output "cloud_credentials_id" {
  value = data.xcsh_cloud_credentials.example.id
}
```

<a id="canonical-cbdf325e26831f48c2e8aa9ff7984b06552b43d04b7c514a335403bf40cbb4a9"></a>

## Root configuration — xcsh_cloud_credentials / e994caf30cb8 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-99c690d995a0792d3d495b786bc242a846ad6fb10b0bc259125d7031c8011007"></a>

## Next pages — xcsh_cloud_credentials / e994caf30cb8 / 6

- [Property reference](../guides/data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [Examples](../guides/data-sources--cloud_credentials--examples--group-001.md#canonical-759b31a35860303b021187e20febff6d1b601860ab93e3dbfe1d30f44fd2e998)
