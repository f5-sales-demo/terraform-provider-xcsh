---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:d4595832c0e0a2d9b444002fbac7b234d5b74f984fef0bfe60acb74bbda30d28", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c", "source_path": "examples/resources/xcsh_protocol_policer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protocol_policer:example:resource", "parent_id": "xcsh-docs:resources:protocol_policer:examples", "path": "documentation/resources/protocol_policer/examples/resource/index.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/examples/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
