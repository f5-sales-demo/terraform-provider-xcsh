---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_forwarding_class."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:719bf67a57476adc67bc53f9ed8876cea919888175745dd766e3363415be305d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:12747b7c8fc5a62067b4a0cf603f5a2dd940ca02738f214a627abb11e786118d", "source_path": "examples/data-sources/xcsh_forwarding_class/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:forwarding_class:example:data-source", "parent_id": "xcsh-docs:data-sources:forwarding_class:examples", "path": "documentation/data-sources/forwarding_class/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211130101120122-3130023111212113-1301010213200202-2030102011033201-3201111120220223-0032001333031030-3311023311310211-2023233100000331", "registry_path": "docs/guides/data-sources--forwarding_class--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_forwarding_class.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
