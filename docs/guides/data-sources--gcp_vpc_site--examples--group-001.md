---
page_title: "xcsh_gcp_vpc_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site examples."
---

# xcsh_gcp_vpc_site examples

<a id="canonical-5b11c53bee7c17cb53c566c755e7e65a5bb1035e2aedfe5a5351a6f68d283a1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1461fa208160060c4e124f6f45013a421a8a0d17c823152b2a7a17d5cf259b5d"></a>

## Examples — Examples / 0970c827a5fb / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- Examples

<a id="canonical-8862a935b4530dbffa7255b5dc4b2621888818b5a0d05d16d22619ea7551c699"></a>

## Complete configurations — Examples / 0970c827a5fb / 3

- [Data source](data-sources--gcp_vpc_site--examples--group-001.md#canonical-104b36094734c1423320e5334bb942fbe4fa0e74ee13289436deb78174a21c6b): valid configuration.

<a id="canonical-224c128770efcba5dcf224b43077ea28fc6dad1d9272d5fc2c6d6d3932c6de5a"></a>

## Next pages — Examples / 0970c827a5fb / 4

- [Data source](data-sources--gcp_vpc_site--examples--group-001.md#canonical-104b36094734c1423320e5334bb942fbe4fa0e74ee13289436deb78174a21c6b)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-104b36094734c1423320e5334bb942fbe4fa0e74ee13289436deb78174a21c6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-179e390a4ca6109ad5bb45adb46a6151893a8dd6184d87b5b055dacd837813db"></a>

## Data source — Data source / 3ce15b7e77bc / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Examples](data-sources--gcp_vpc_site--examples--group-001.md#canonical-5b11c53bee7c17cb53c566c755e7e65a5bb1035e2aedfe5a5351a6f68d283a1d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_gcp_vpc_site/data-source.tf`; digest `sha256:6ff7b0833cac6e9cc603f3fd4746eb2dd86fa9ff580a9017d62c97b627589c33`.

```terraform
# GCPVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GCPVPCSite by name
data "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"
}

output "gcp_vpc_site_id" {
  value = data.xcsh_gcp_vpc_site.example.id
}
```

<a id="canonical-4d350e0675a550e1d5ead4b6c970bd2e9e239226690346d33545c1b938d7f175"></a>

## Next pages — Data source / 3ce15b7e77bc / 3

- [Examples](data-sources--gcp_vpc_site--examples--group-001.md#canonical-5b11c53bee7c17cb53c566c755e7e65a5bb1035e2aedfe5a5351a6f68d283a1d)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
