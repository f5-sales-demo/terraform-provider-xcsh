---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_allowed_domain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1071, "body_sha256": "sha256:6645bfa540712de0387e30e62c0c6f1da6820a7cd5951a01c38678e159cef4e9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:allowed_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006", "source_path": "examples/data-sources/xcsh_allowed_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:allowed_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:allowed_domain:examples", "path": "documentation/data-sources/allowed_domain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-0100000003211222-1011030202302300-1223321113203003-2011211123233321-2112020030000123-0001300223021013-2131200103302111-3020112332210010", "registry_path": "docs/guides/data-sources--allowed_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/allowed_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Data source for xcsh_allowed_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_allowed_domain/data-source.tf`; digest `sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006`.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```
