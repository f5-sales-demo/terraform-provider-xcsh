---
page_title: "xcsh_site_image landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_image landing."
---

# xcsh_site_image landing

<a id="canonical-ddab21760c47eb7a0955908fcfd208fc71d1b2d519a90c01d28e9ccb3de38892"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57aebe17ff2655bef8ec308009a56bfadc8080b3af5d1833ac85b152707ce353"></a>

## xcsh_site_image — xcsh_site_image / 3c9611fa05d7 / 2

Breadcrumbs:

- xcsh_site_image

Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No
caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution
does not imply successful boot.

<a id="canonical-e2f0b65a3ce30999a49e9ce03f51e8cacba16ff5151e39fa386a026504295675"></a>

## Prerequisites — xcsh_site_image / 3c9611fa05d7 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-492d7ed15b65b2a10aabaf5a0de9175ccdcd009bf79436266d8f88b0e91edec0"></a>

## Minimal configuration — xcsh_site_image / 3c9611fa05d7 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

<a id="canonical-988abe69406831be4f766a113f36454da7d676b32b8ef9c91182166357734917"></a>

## Root configuration — xcsh_site_image / 3c9611fa05d7 / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-3c26b4dfd2a79494cf71496fe61bf1bc804d23e6f9b45a1ebc362beecde8a2c5"></a>

## Next pages — xcsh_site_image / 3c9611fa05d7 / 6

- [Property reference](../guides/data-sources--site_image--reference--group-001.md#canonical-ec8e309db65094e6e075991b728eb4bab2f7412acc5f567bcc844c1bea6db734)
- [Examples](../guides/data-sources--site_image--examples--group-001.md#canonical-466f5d8a3bd4437c6e0cb2fd0469ed916df5aee1038a8f0caa3bece7f63a73c3)
