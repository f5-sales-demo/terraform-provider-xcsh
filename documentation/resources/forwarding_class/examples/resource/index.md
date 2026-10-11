---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_forwarding_class."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1108, "body_sha256": "sha256:5251913723b9bcc0dee892f389c456dde1cf2d9130a751183e200b8c6eac16e2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8", "source_path": "examples/resources/xcsh_forwarding_class/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:forwarding_class:example:resource", "parent_id": "xcsh-docs:resources:forwarding_class:examples", "path": "documentation/resources/forwarding_class/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3033303113103310-0131012220003123-2322112013321232-0102023201210323-0131131012012333-0020000210022302-2023212223111123-2210111103001110", "registry_path": "docs/guides/resources--forwarding_class--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_forwarding_class.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forwarding_class/resource.tf`; digest `sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8`.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```
