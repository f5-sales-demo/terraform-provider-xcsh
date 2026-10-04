---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike1."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1176, "body_sha256": "sha256:5c746ddec6f64fc94512f9517dedd40caa4ce903f933e6cfccea898a13caba91", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e", "source_path": "examples/resources/xcsh_ike1/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike1:example:resource", "parent_id": "xcsh-docs:resources:ike1:examples", "path": "documentation/resources/ike1/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0212202023121020-3323020003311120-1210331132100101-3102212110230210-3030213020330112-1112320212102311-2002202120100300-0012001023112221", "registry_path": "docs/guides/resources--ike1--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_ike1.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["ike1CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike1/resource.tf`; digest `sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e`.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/examples/)
- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
