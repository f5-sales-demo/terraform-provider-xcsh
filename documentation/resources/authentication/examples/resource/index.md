---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_authentication."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1031, "body_sha256": "sha256:37d7e6f395db9f8652e0f6410fc01423e0d051fbb5968c082dfe27b8ef121a72", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876", "source_path": "examples/resources/xcsh_authentication/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:authentication:example:resource", "parent_id": "xcsh-docs:resources:authentication:examples", "path": "documentation/resources/authentication/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3102111333213010-3213112102323101-2302032301012330-3033323221301032-3101111301120310-2321221030132032-0200003312022213-0123200321031200", "registry_path": "docs/guides/resources--authentication--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_authentication.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["authenticationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authentication/resource.tf`; digest `sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876`.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```
