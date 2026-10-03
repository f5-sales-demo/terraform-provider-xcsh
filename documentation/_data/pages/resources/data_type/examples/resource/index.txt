---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_type."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1253, "body_sha256": "sha256:fe49316b20d23b941b5d25be29ca51efd27b7f1d2c9c4da28b25e897c461bbc9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee", "source_path": "examples/resources/xcsh_data_type/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_type:example:resource", "parent_id": "xcsh-docs:resources:data_type:examples", "path": "documentation/resources/data_type/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3123010221000010-1300110032110002-0020113210323203-3323003033130211-3223132211302022-1011103232001133-3101011012212000-2123200123233222", "registry_path": "docs/guides/resources--data_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_data_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_type/resource.tf`; digest `sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee`.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/examples/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
