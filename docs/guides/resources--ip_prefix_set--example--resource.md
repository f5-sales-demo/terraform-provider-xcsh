---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": [], "body_bytes": 1093, "body_sha256": "sha256:bae84382c3495439b83982c2bd4c656d4131c3f12b531dbd72f57f4f8ce73c8e", "canonical_id": "xcsh-docs:resources:ip_prefix_set:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a", "source_path": "examples/resources/xcsh_ip_prefix_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ip_prefix_set:example:resource", "parent_id": "xcsh-docs:resources:ip_prefix_set:examples", "path": "docs/guides/resources--ip_prefix_set--example--resource.md", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_ip_prefix_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md)
- [Examples](resources--ip_prefix_set--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ip_prefix_set/resource.tf`; digest `sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a`.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--ip_prefix_set--examples.md)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md)
