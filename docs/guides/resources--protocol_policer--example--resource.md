---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1079, "body_sha256": "sha256:42f55f20dabce11206de26b2a72dbde0673dce4f8a2f373d19f53684cd160f25", "canonical_id": "xcsh-docs:resources:protocol_policer:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c", "source_path": "examples/resources/xcsh_protocol_policer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protocol_policer:example:resource", "parent_id": "xcsh-docs:resources:protocol_policer:examples", "path": "docs/guides/resources--protocol_policer--example--resource.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Examples](resources--protocol_policer--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_policer/resource.tf`; digest `sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c`.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```

## Next pages

- [Examples](resources--protocol_policer--examples.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
