---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_forwarding_class."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:719bf67a57476adc67bc53f9ed8876cea919888175745dd766e3363415be305d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:12747b7c8fc5a62067b4a0cf603f5a2dd940ca02738f214a627abb11e786118d", "source_path": "examples/data-sources/xcsh_forwarding_class/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:forwarding_class:example:data-source", "parent_id": "xcsh-docs:data-sources:forwarding_class:examples", "path": "documentation/data-sources/forwarding_class/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1211130101120122-3130023111212113-1301010213200202-2030102011033201-3201111120220223-0032001333031030-3311023311310211-2023233100000331", "registry_path": "docs/guides/data-sources--forwarding_class--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_forwarding_class.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forwarding_class/data-source.tf`; digest `sha256:12747b7c8fc5a62067b4a0cf603f5a2dd940ca02738f214a627abb11e786118d`.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/examples/)
- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
