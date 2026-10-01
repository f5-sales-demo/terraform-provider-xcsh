---
page_title: "xcsh_application_profiles examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles examples."
---

# xcsh_application_profiles examples

<a id="canonical-4f157bf871ffc5e090bca94e3729fc5b9aa57f302edc85a8cb0c006e9a14aad4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5ebc2f7440d145b7f783e9b51e3327e98a1f8292409e7bf98fe8abecfb9a35e"></a>

## Examples — Examples / ce671fcd6c50 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- Examples

<a id="canonical-bab2e5340115a0fbb44274c617450283a5c9afd8ac26561cea5583a1e7c9594d"></a>

## Complete configurations — Examples / ce671fcd6c50 / 3

- [Data source](data-sources--application_profiles--examples--group-001.md#canonical-3e46cccd95f3388b43a1213f7211d632f8e9d177c4c9f1f2ccc815f82713492d): valid configuration.

<a id="canonical-c214c290dfab4a1c7520d891ee7d0147373374e09e38901c55dcf854ba29f5d2"></a>

## Next pages — Examples / ce671fcd6c50 / 4

- [Data source](data-sources--application_profiles--examples--group-001.md#canonical-3e46cccd95f3388b43a1213f7211d632f8e9d177c4c9f1f2ccc815f82713492d)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3e46cccd95f3388b43a1213f7211d632f8e9d177c4c9f1f2ccc815f82713492d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5a81d8f5d8a64d39e9713ebbb04cc8ec8687eb756782c7f154e5a8012183e90"></a>

## Data source — Data source / 44df030df38c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Examples](data-sources--application_profiles--examples--group-001.md#canonical-4f157bf871ffc5e090bca94e3729fc5b9aa57f302edc85a8cb0c006e9a14aad4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_application_profiles/data-source.tf`; digest `sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc`.

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

<a id="canonical-bcc32cb4ff8abd46bdbe525a362522f7755469fe1eaf30b40cce5f3a36485403"></a>

## Next pages — Data source / 44df030df38c / 3

- [Examples](data-sources--application_profiles--examples--group-001.md#canonical-4f157bf871ffc5e090bca94e3729fc5b9aa57f302edc85a8cb0c006e9a14aad4)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
