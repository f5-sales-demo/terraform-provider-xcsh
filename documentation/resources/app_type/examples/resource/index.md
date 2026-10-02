---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_type."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1236, "body_sha256": "sha256:84419f9e693791cd79ab905d6579b58ee428ef7cb971cacb1a7cf628e926eb81", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d", "source_path": "examples/resources/xcsh_app_type/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_type:example:resource", "parent_id": "xcsh-docs:resources:app_type:examples", "path": "documentation/resources/app_type/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1223322221000032-2221330233112230-1021112133211022-3011302012032303-3313202032110100-3133201300220323-0311011032013311-0331032330111201", "registry_path": "docs/guides/resources--app_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_app_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_type/resource.tf`; digest `sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d`.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/examples/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
