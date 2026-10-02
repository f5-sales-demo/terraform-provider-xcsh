---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike2."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1176, "body_sha256": "sha256:50087717f3b89c650d22fa07220570a7150702a2456b2e892865262a05e59c48", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9af0540e9ef3cddbd2cf33f981bb394c77fd860b251d272dd666ef02cd5bc017", "source_path": "examples/resources/xcsh_ike2/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike2:example:resource", "parent_id": "xcsh-docs:resources:ike2:examples", "path": "documentation/resources/ike2/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3211023310301031-0330202110310232-1122212202130031-3311003323310133-1320332122302023-0213113311301321-1312233300101022-1122221312211022", "registry_path": "docs/guides/resources--ike2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_ike2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike2/resource.tf`; digest `sha256:9af0540e9ef3cddbd2cf33f981bb394c77fd860b251d272dd666ef02cd5bc017`.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/examples/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
