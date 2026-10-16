---
page_title: "xcsh_site_registrations_by_site"
subcategory: ""
description: "List registrations for a Customer Edge site."
xcsh_docs: {"aliases": ["site registrations by site"], "body_bytes": 1334, "body_sha256": "sha256:17a1823857bf4fc006e8ee97ddfd66109950800aa779f9c538ae810a72986eeb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:reference", "xcsh-docs:data-sources:site_registrations_by_site:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_registrations_by_site/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2111303110111030-3113020121012123-1313322201232310-1013111203201222-0203220021201023-0033331330113033-0303303330002221-2203131220023320", "registry_path": "docs/data-sources/site_registrations_by_site.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "List registrations for a Customer Edge site.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_registrations_by_site

Breadcrumbs:

- xcsh_site_registrations_by_site

List registrations for a Customer Edge site.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```

## Root configuration

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/examples/)
