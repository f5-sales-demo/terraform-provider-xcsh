---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike2."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 978, "body_sha256": "sha256:56852d078c45c0c2694a4cffa2390745a500c5c9f12cd85faa4ff4c497e39292", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9af0540e9ef3cddbd2cf33f981bb394c77fd860b251d272dd666ef02cd5bc017", "source_path": "examples/resources/xcsh_ike2/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike2:example:resource", "parent_id": "xcsh-docs:resources:ike2:examples", "path": "documentation/resources/ike2/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3211023310301031-0330202110310232-1122212202130031-3311003323310133-1320332122302023-0213113311301321-1312233300101022-1122221312211022", "registry_path": "docs/guides/resources--ike2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_ike2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
