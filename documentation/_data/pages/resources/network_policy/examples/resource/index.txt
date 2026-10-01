---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:e4ca4234430eafdf28f5530f5cb4b063b0c1d728afcd8fd9fbc87f2cb0765687", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887", "source_path": "examples/resources/xcsh_network_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy:example:resource", "parent_id": "xcsh-docs:resources:network_policy:examples", "path": "documentation/resources/network_policy/examples/resource/index.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy/resource.tf`; digest `sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887`.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
