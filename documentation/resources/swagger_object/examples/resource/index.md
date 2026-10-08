---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_swagger_object."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1061, "body_sha256": "sha256:bf7821c8ccaeaa676031e0cc04fa19e9734d690c276b9e92bddb197eb028dffe", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:swagger_object:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ac6ca388f618a440a05eeef1c5d8d8a3ed3b1914604de5ad200997bcadce42a7", "source_path": "examples/resources/xcsh_swagger_object/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:swagger_object:example:resource", "parent_id": "xcsh-docs:resources:swagger_object:examples", "path": "documentation/resources/swagger_object/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "swagger_object", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0112212111233001-3021103220033002-3101233330202101-1102101332032210-3312030123003320-2010310212111112-2032111312312230-3313132120222133", "registry_path": "docs/guides/resources--swagger_object--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/swagger_object/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_swagger_object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_swagger_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_swagger_object/resource.tf`; digest `sha256:ac6ca388f618a440a05eeef1c5d8d8a3ed3b1914604de5ad200997bcadce42a7`.

```terraform
terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema"
  content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
```
