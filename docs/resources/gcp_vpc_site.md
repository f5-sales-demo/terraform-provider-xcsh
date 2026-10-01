---
page_title: "xcsh_gcp_vpc_site"
subcategory: "Infrastructure"
description: "xcsh_gcp_vpc_site for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1706, "body_sha256": "sha256:e45602039f2668912dbd6a591efd894fad677dc2adfd8ee119db457ce5a03881", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:fundamentals", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:reference", "xcsh-docs:resources:gcp_vpc_site:examples", "xcsh-docs:resources:gcp_vpc_site:import", "xcsh-docs:resources:gcp_vpc_site:timeouts"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:fundamentals", "parent_id": null, "path": "docs/resources/gcp_vpc_site.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_gcp_vpc_site for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_gcp_vpc_site

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GCPVPCSite Resource Example
# Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GCPVPCSite configuration
resource "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"

  gcp_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

## Root configuration

Required root properties: `gcp_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--gcp_vpc_site--reference.md)
- [Examples](../guides/resources--gcp_vpc_site--examples.md)
- [Import](../guides/resources--gcp_vpc_site--import.md)
- [Timeouts](../guides/resources--gcp_vpc_site--timeouts.md)
