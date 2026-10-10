---
page_title: "xcsh_swagger_object"
subcategory: ""
description: "Owns one immutable, content-verified Swagger object version. Use a content-addressed name: content changes require a new name and replace only the owned version. XC may reuse a deleted version label, so same-name content replacement is rejected. Import uses namespace/name/version; latest and external presigned URLs"
xcsh_docs: {"aliases": ["adopt existing object", "import existing resource", "swagger object"], "body_bytes": 1926, "body_sha256": "sha256:08a7e2b6a7545fca8004218cafbeb20c80f9949bc7143293c6d75caf1f9f84bb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:swagger_object:reference", "xcsh-docs:resources:swagger_object:examples", "xcsh-docs:resources:swagger_object:import"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:swagger_object:collection", "completeness": "complete", "id": "xcsh-docs:resources:swagger_object:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/swagger_object/index.md", "product": "distributed-cloud", "provider_name": "swagger_object", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0113122122131303-3312011120012020-2331301002012223-3033302331132230-0311201302013200-2313211231330002-3123103203213300-2123202121003133", "registry_path": "docs/resources/swagger_object.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/swagger_object/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Owns one immutable, content-verified Swagger object version. Use a content-addressed name: content changes require a new name and replace only the owned version. XC may reuse a deleted version label, so same-name content replacement is rejected. Import uses namespace/name/version; latest and external presigned URLs", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_swagger_object

Breadcrumbs:

- xcsh_swagger_object

Owns one immutable, content-verified Swagger object version. Use a content-addressed name: content
changes require a new name and replace only the owned version. XC may reuse a deleted version label,
so same-name content replacement is rejected. Import uses namespace/name/version; latest and
external presigned URLs are prohibited. State contains the complete document.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `content`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/lifecycle/import/)
