---
page_title: "xcsh_protocol_policer"
subcategory: ""
description: "xcsh_protocol_policer for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1340, "body_sha256": "sha256:268e7112d54839219087f3aa94c33b7f43d834331e049f6ccc933ed58ab3e782", "child_ids": ["xcsh-docs:data-sources:protocol_policer:reference", "xcsh-docs:data-sources:protocol_policer:examples"], "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:fundamentals", "parent_id": null, "path": "documentation/data-sources/protocol_policer/index.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_protocol_policer for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_protocol_policer

Breadcrumbs:

- xcsh_protocol_policer

Manages protocol\_policer object, protocol\_policer object contains list of L4 protocol match
condition and corresponding traffic rate limits in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/examples/)
