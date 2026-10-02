---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1394, "body_sha256": "sha256:14f53a43f91c79bcc8c8c2d26b3070ca14c8904b66c127a214def0989108451d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0", "source_path": "examples/resources/xcsh_srv6_network_slice/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:srv6_network_slice:example:resource", "parent_id": "xcsh-docs:resources:srv6_network_slice:examples", "path": "documentation/resources/srv6_network_slice/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1331111120302201-1301121320111201-3320103121320102-0111112211101230-0010013013200103-1200330021230010-3021100201320131-2110302311023232", "registry_path": "docs/guides/resources--srv6_network_slice--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/srv6_network_slice/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_srv6_network_slice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_srv6_network_slice/resource.tf`; digest `sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0`.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/examples/)
- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/)
