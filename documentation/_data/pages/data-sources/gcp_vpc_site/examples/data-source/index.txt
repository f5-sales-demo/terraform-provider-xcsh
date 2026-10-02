---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1277, "body_sha256": "sha256:e7f7c5cbee4bc89bc6adc2f0d104d88895560b4276f1065fd10bb00e373d63bc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6ff7b0833cac6e9cc603f3fd4746eb2dd86fa9ff580a9017d62c97b627589c33", "source_path": "examples/data-sources/xcsh_gcp_vpc_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:gcp_vpc_site:example:data-source", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:examples", "path": "documentation/data-sources/gcp_vpc_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0100102303120021-1013031030011002-0303020032110303-1023232110023323-3210332200321310-3232010302202110-0312313223132001-1310220201301223", "registry_path": "docs/guides/data-sources--gcp_vpc_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_gcp_vpc_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/examples/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
