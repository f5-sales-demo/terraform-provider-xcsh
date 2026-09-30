---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1076, "body_sha256": "sha256:78faadb098edd3c5db8bb664a249853ef3f41bad6f871325820b4bfd7632ffb3", "canonical_id": "xcsh-docs:resources:policy_based_routing:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968", "source_path": "examples/resources/xcsh_policy_based_routing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:policy_based_routing:example:resource", "parent_id": "xcsh-docs:resources:policy_based_routing:examples", "path": "docs/guides/resources--policy_based_routing--example--resource.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Examples](resources--policy_based_routing--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policy_based_routing/resource.tf`; digest `sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968`.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--policy_based_routing--examples.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
