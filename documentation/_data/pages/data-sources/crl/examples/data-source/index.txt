---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_crl."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1164, "body_sha256": "sha256:36b3920214bb18a96df4f7e7518c8ca0c44155b7e4e6e734e89ed726ce45d7f2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a", "source_path": "examples/data-sources/xcsh_crl/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:crl:example:data-source", "parent_id": "xcsh-docs:data-sources:crl:examples", "path": "documentation/data-sources/crl/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3020121233030311-3122033322301333-2312211200310011-2031300000230123-2113303100023233-1310201220030010-2102030210123201-0312131312201322", "registry_path": "docs/guides/data-sources--crl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_crl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["crlCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_crl/data-source.tf`; digest `sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a`.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
