---
page_title: "xcsh_network_global_controller_sso_egress"
subcategory: ""
description: "xcsh_network_global_controller_sso_egress for xcsh_network_global_controller_sso_egress."
xcsh_docs: {"aliases": [], "body_bytes": 1375, "body_sha256": "sha256:af97b4acbdeff09ca15d4a80eefbfedc245b4f06b7bf21d92ecef4383467745b", "canonical_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:fundamentals", "child_ids": ["xcsh-docs:data-sources:network_global_controller_sso_egress:reference", "xcsh-docs:data-sources:network_global_controller_sso_egress:examples"], "collection_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_global_controller_sso_egress:fundamentals", "parent_id": null, "path": "docs/data-sources/network_global_controller_sso_egress.md", "provider_name": "network_global_controller_sso_egress", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_controller_sso_egress/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_global_controller_sso_egress for xcsh_network_global_controller_sso_egress.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_network_global_controller_sso_egress

Breadcrumbs:

- xcsh_network_global_controller_sso_egress

Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--network_global_controller_sso_egress--reference.md)
- [Examples](../guides/data-sources--network_global_controller_sso_egress--examples.md)
