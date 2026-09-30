---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_global_controller_sso_egress."
xcsh_docs: {"aliases": [], "body_bytes": 1170, "body_sha256": "sha256:72385adf8ddd914d85d446c88059b90f6d0c1ac1ea0e3ba03467574a24edc387", "canonical_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00", "source_path": "examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_global_controller_sso_egress:example:data-source", "parent_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:examples", "path": "docs/guides/data-sources--network_global_controller_sso_egress--example--data-source.md", "provider_name": "network_global_controller_sso_egress", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_controller_sso_egress/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_global_controller_sso_egress.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md)
- [Examples](data-sources--network_global_controller_sso_egress--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf`; digest `sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00`.

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

## Next pages

- [Examples](data-sources--network_global_controller_sso_egress--examples.md)
- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md)
