---
page_title: "xcsh_swagger_object"
subcategory: ""
description: "Owns one immutable, content-verified Swagger object version. Content changes replace only the owned version. Import uses namespace/name/version; latest and external presigned URLs are prohibited. State contains the complete document."
xcsh_docs: {"aliases": ["adopt existing object", "import existing resource", "swagger object"], "body_bytes": 1637, "body_sha256": "sha256:89d3999b5828e281bb2b538f82cffca57d9ce65102c91867011804e7c5c9f8c6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:swagger_object:reference", "xcsh-docs:resources:swagger_object:examples", "xcsh-docs:resources:swagger_object:import"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:swagger_object:collection", "completeness": "complete", "id": "xcsh-docs:resources:swagger_object:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/swagger_object/index.md", "product": "distributed-cloud", "provider_name": "swagger_object", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0113122122131303-3312011120012020-2331301002012223-3033302331132230-0311201302013200-2313211231330002-3123103203213300-2123202121003133", "registry_path": "docs/resources/swagger_object.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/swagger_object/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Owns one immutable, content-verified Swagger object version. Content changes replace only the owned version. Import uses namespace/name/version; latest and external presigned URLs are prohibited. State contains the complete document.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_swagger_object

Breadcrumbs:

- xcsh_swagger_object

Owns one immutable, content-verified Swagger object version. Content changes replace only the owned
version. Import uses namespace/name/version; latest and external presigned URLs are prohibited.
State contains the complete document.

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

## Root configuration

Required root properties: `content`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/lifecycle/import/)
