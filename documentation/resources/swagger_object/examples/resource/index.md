---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_swagger_object."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1213, "body_sha256": "sha256:c0288006d656ae8a61b2500cbd1e053e797f09d8a59aeebf4e33187aadf96588", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:swagger_object:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3e856ac623c4fd62ed456c51c40b6d7398bf0e4040635eb274d1802a30e02ed2", "source_path": "examples/resources/xcsh_swagger_object/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:swagger_object:example:resource", "parent_id": "xcsh-docs:resources:swagger_object:examples", "path": "documentation/resources/swagger_object/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "swagger_object", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0112212111233001-3021103220033002-3101233330202101-1102101332032210-3312030123003320-2010310212111112-2032111312312230-3313132120222133", "registry_path": "docs/guides/resources--swagger_object--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/swagger_object/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_swagger_object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

Source: `examples/resources/xcsh_swagger_object/resource.tf`; digest `sha256:3e856ac623c4fd62ed456c51c40b6d7398bf0e4040635eb274d1802a30e02ed2`.

```terraform
terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

locals {
  schema_content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema-${substr(sha256(local.schema_content), 0, 32)}"
  content   = local.schema_content
  lifecycle {
    create_before_destroy = true
  }
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
```
