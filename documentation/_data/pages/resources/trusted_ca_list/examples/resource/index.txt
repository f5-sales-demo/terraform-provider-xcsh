---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1317, "body_sha256": "sha256:d9e9cab79e96ac2459302e46a8bd23ae91ad8429e1fe64340a1627d47455a65d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328", "source_path": "examples/resources/xcsh_trusted_ca_list/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:trusted_ca_list:example:resource", "parent_id": "xcsh-docs:resources:trusted_ca_list:examples", "path": "documentation/resources/trusted_ca_list/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2103021213110323-1232123021100101-0100111023203031-2120302130122313-1013031001112213-0112213033332111-0122031310202201-0001213323302301", "registry_path": "docs/guides/resources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
