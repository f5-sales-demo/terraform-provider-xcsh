---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1071, "body_sha256": "sha256:4ae0add1f21f3a966591db74a3042867628ae4b0166454486662339b560cadf3", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6ff7b0833cac6e9cc603f3fd4746eb2dd86fa9ff580a9017d62c97b627589c33", "source_path": "examples/data-sources/xcsh_gcp_vpc_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:gcp_vpc_site:example:data-source", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:examples", "path": "docs/guides/data-sources--gcp_vpc_site--example--data-source.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Examples](data-sources--gcp_vpc_site--examples.md)
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

## Next pages

- [Examples](data-sources--gcp_vpc_site--examples.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
