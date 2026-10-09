---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_route."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 943, "body_sha256": "sha256:7a9f2ba9d685f0c98b874b17b1b3fc79359d4d636375a99130ed173bf7b910b7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3", "source_path": "examples/resources/xcsh_route/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:route:example:resource", "parent_id": "xcsh-docs:resources:route:examples", "path": "documentation/resources/route/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1133233331113213-3131330001133202-1322003033201213-3003331311023122-3031101022221102-1021211012333310-3030033111030203-2220313212013230", "registry_path": "docs/guides/resources--route--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["routeCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_route/resource.tf`; digest `sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3`.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```
