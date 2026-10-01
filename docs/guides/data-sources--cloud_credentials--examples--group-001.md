---
page_title: "xcsh_cloud_credentials examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials examples."
---

# xcsh_cloud_credentials examples

<a id="canonical-759b31a35860303b021187e20febff6d1b601860ab93e3dbfe1d30f44fd2e998"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8319d617a287e7cd4ca57d453a3d3b9fc466b3b7669ca4aa5caace4ed00d8e07"></a>

## Examples — Examples / 63b57f8dd3e5 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- Examples

<a id="canonical-1ad2fa43a5685a2d107d2feed47285c4b14a37993067fb34a4ea830cc89b5ccc"></a>

## Complete configurations — Examples / 63b57f8dd3e5 / 3

- [Data source](data-sources--cloud_credentials--examples--group-001.md#canonical-e5078dd6e433052b3a5780a9b80d4dcfa49c800ce4fffd57722b9299178c54d7): valid configuration.

<a id="canonical-bf11d65f9a2f1e3f274a3e82962c64b32475f084a812e6815f3f742266de6c38"></a>

## Next pages — Examples / 63b57f8dd3e5 / 4

- [Data source](data-sources--cloud_credentials--examples--group-001.md#canonical-e5078dd6e433052b3a5780a9b80d4dcfa49c800ce4fffd57722b9299178c54d7)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-e5078dd6e433052b3a5780a9b80d4dcfa49c800ce4fffd57722b9299178c54d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55cd0e9385e580868cbbbc70dbb3c5bf5b035f8a4ac17a61b3ca1b73c50615d4"></a>

## Data source — Data source / 7f1c4d2c8ea9 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Examples](data-sources--cloud_credentials--examples--group-001.md#canonical-759b31a35860303b021187e20febff6d1b601860ab93e3dbfe1d30f44fd2e998)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_credentials/data-source.tf`; digest `sha256:6a62066c747f97e4baff52829aa9bc36ad62c8b9304220fe45b6e8d9ef9839f1`.

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

<a id="canonical-d33686214abf1a2812b8bf504eea478621845d60d3a15da2079ab8d11344faa7"></a>

## Next pages — Data source / 7f1c4d2c8ea9 / 3

- [Examples](data-sources--cloud_credentials--examples--group-001.md#canonical-759b31a35860303b021187e20febff6d1b601860ab93e3dbfe1d30f44fd2e998)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
