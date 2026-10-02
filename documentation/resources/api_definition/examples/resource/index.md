---
page_title: "Resource"
subcategory: "API Management"
description: "Resource for xcsh_api_definition."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1246, "body_sha256": "sha256:f2c8e1b99039dd69d3a49dc0868ab420bf1202a6bb6c2a7b90d0ab98e018f13c", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64", "source_path": "examples/resources/xcsh_api_definition/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_definition:example:resource", "parent_id": "xcsh-docs:resources:api_definition:examples", "path": "documentation/resources/api_definition/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1213001112132303-3333011121103100-3210233222333023-0331300210320303-1230003312130322-2100032100011211-2201033221322101-2011122021010300", "registry_path": "docs/guides/resources--api_definition--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_api_definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_definitionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_definition/resource.tf`; digest `sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64`.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/examples/)
- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
