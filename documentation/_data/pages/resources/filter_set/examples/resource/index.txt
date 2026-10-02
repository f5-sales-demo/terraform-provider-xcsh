---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_filter_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1234, "body_sha256": "sha256:1bfd7d5be4bee888f2d3ffad954c506e3bace1cda55d3b214be2fe4139632452", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610", "source_path": "examples/resources/xcsh_filter_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:filter_set:example:resource", "parent_id": "xcsh-docs:resources:filter_set:examples", "path": "documentation/resources/filter_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2313012200213310-2312010001331203-2320031131200101-3010310321330001-2010300010123111-0202311102003120-3021311011232101-0012112122113133", "registry_path": "docs/guides/resources--filter_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_filter_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["filter_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_filter_set/resource.tf`; digest `sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610`.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/examples/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
