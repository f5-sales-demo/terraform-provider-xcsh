---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1195, "body_sha256": "sha256:bee7c8651cfb517fe95c2225c3721f410dcd3389a6b919beb86c133c4af34888", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690", "source_path": "examples/resources/xcsh_data_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_group:example:resource", "parent_id": "xcsh-docs:resources:data_group:examples", "path": "documentation/resources/data_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2203123231131002-1010333110013022-0002230320220313-2231031310211333-1211333111103021-2310133322223011-1021233301312012-0111201321002133", "registry_path": "docs/guides/resources--data_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_data_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_group/resource.tf`; digest `sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690`.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
