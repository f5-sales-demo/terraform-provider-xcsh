---
page_title: "Resource"
subcategory: "API Management"
description: "Resource for xcsh_api_definition."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1018, "body_sha256": "sha256:2163d16427e37ca893222bd692e2f9735d3c3cc21d8a990f0a4db304cd2cef08", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64", "source_path": "examples/resources/xcsh_api_definition/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_definition:example:resource", "parent_id": "xcsh-docs:resources:api_definition:examples", "path": "documentation/resources/api_definition/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1213001112132303-3333011121103100-3210233222333023-0331300210320303-1230003312130322-2100032100011211-2201033221322101-2011122021010300", "registry_path": "docs/guides/resources--api_definition--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/examples/resource/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Resource for xcsh_api_definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["api_definitionCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
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
