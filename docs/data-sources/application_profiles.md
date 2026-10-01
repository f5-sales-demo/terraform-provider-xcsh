---
page_title: "xcsh_application_profiles landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles landing."
---

# xcsh_application_profiles landing

<a id="canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d795432dc8c4967968115ff6fa7a1add32638cde0836a724c4d299be0e1328f"></a>

## xcsh_application_profiles — xcsh_application_profiles / bb7f2aeb04fc / 2

Breadcrumbs:

- xcsh_application_profiles

Manages Application Profiles in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-ee3babe6be88710ace7c7b01cbf4a0cc4366b452617fdf4b370acee5300e0cb7"></a>

## Prerequisites — xcsh_application_profiles / bb7f2aeb04fc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-792881b3863b7685ad7b7809d0aa6376cf136aebc064bf1e828fb958994feefb"></a>

## Minimal configuration — xcsh_application_profiles / bb7f2aeb04fc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ApplicationProfiles by name
data "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}

output "application_profiles_id" {
  value = data.xcsh_application_profiles.example.id
}
```

<a id="canonical-cad611ce011d09eb7c04f92a9394d9f864a5b38772699d566d02486f036539c0"></a>

## Root configuration — xcsh_application_profiles / bb7f2aeb04fc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-deb9472860031ec7e7a6e19f935ed2cafb27023546234b9bafe30bac7f274581"></a>

## Next pages — xcsh_application_profiles / bb7f2aeb04fc / 6

- [Property reference](../guides/data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [Examples](../guides/data-sources--application_profiles--examples--group-001.md#canonical-4f157bf871ffc5e090bca94e3729fc5b9aa57f302edc85a8cb0c006e9a14aad4)
