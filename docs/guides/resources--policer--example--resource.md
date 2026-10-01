---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1054, "body_sha256": "sha256:6c86261675ba26b524a6b778a253c6a8d7b0af4df9436f13c2ea5f99310839c2", "canonical_id": "xcsh-docs:resources:policer:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443", "source_path": "examples/resources/xcsh_policer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:policer:example:resource", "parent_id": "xcsh-docs:resources:policer:examples", "path": "docs/guides/resources--policer--example--resource.md", "provider_name": "policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_policer](../resources/policer.md)
- [Examples](resources--policer--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policer/resource.tf`; digest `sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443`.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```

## Next pages

- [Examples](resources--policer--examples.md)
- [xcsh_policer](../resources/policer.md)
