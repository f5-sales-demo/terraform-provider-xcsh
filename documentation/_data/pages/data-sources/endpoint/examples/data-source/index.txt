---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_endpoint."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1229, "body_sha256": "sha256:8cada1b908a7941f9b97dcdd7c401e3fd69f71329cdcdb5facc2403667cea85a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b", "source_path": "examples/data-sources/xcsh_endpoint/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:endpoint:example:data-source", "parent_id": "xcsh-docs:data-sources:endpoint:examples", "path": "documentation/data-sources/endpoint/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0201022332111231-2211112003210301-2000031231200003-0300022131133233-0120131121000201-3212001123100330-2113220012031232-3021322031332022", "registry_path": "docs/guides/data-sources--endpoint--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_endpoint/data-source.tf`; digest `sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b`.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/examples/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
