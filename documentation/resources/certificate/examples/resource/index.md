---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_certificate."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1276, "body_sha256": "sha256:53c9043b42c674c3141b03d66109a2a437a216ae364eb3d3ed2c63c86736db67", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092", "source_path": "examples/resources/xcsh_certificate/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:certificate:example:resource", "parent_id": "xcsh-docs:resources:certificate:examples", "path": "documentation/resources/certificate/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1112221302211031-2031103122010300-1001020001123022-3203200030003230-3021230330303200-2232230301223022-2311312301202222-1003332233301021", "registry_path": "docs/guides/resources--certificate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate/resource.tf`; digest `sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092`.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
