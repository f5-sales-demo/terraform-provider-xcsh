---
page_title: "xcsh_gcp_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site landing."
---

# xcsh_gcp_vpc_site landing

<a id="canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72e07cae866efc2692323ff3fd94132684af880d4dbb4d6dd76ee746d5067be4"></a>

## xcsh_gcp_vpc_site — xcsh_gcp_vpc_site / 739c20cdeb47 / 2

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

<a id="canonical-634ef363811c99e928864c335d6fa5918f62a1e9ccdf2418fbbbe4909b6d539d"></a>

## Prerequisites — xcsh_gcp_vpc_site / 739c20cdeb47 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

<a id="canonical-c420823a4bc4235c1602bbaa0fe0242d6b07dfd469e1bb2e74e0e57dc380b56a"></a>

## Minimal configuration — xcsh_gcp_vpc_site / 739c20cdeb47 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-14ee8e66fb200f2f6ecf9219ea70e79c42b3126a67a94a97ed190d81c56d8888"></a>

## Root configuration — xcsh_gcp_vpc_site / 739c20cdeb47 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1fc4711ff1f039ea8add2d136bdefebfbd1804c9855b158f8225277ed6bc07db"></a>

## Next pages — xcsh_gcp_vpc_site / 739c20cdeb47 / 6

- [Property reference](../guides/data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [Examples](../guides/data-sources--gcp_vpc_site--examples--group-001.md#canonical-5b11c53bee7c17cb53c566c755e7e65a5bb1035e2aedfe5a5351a6f68d283a1d)
