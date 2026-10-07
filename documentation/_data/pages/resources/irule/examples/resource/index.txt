---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_irule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1011, "body_sha256": "sha256:671252a451549847ad14fafa14c7ed81a1ee2ba9f30b974110137e5e5ce041d8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b", "source_path": "examples/resources/xcsh_irule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:irule:example:resource", "parent_id": "xcsh-docs:resources:irule:examples", "path": "documentation/resources/irule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0121033310322201-2113232211001202-2212002230230323-1023202203302203-0131202322320320-0022210011030332-2011130333023221-3113320313120331", "registry_path": "docs/guides/resources--irule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_irule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["iruleCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_irule/resource.tf`; digest `sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b`.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```
