---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_global_controller_sso_egress."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1475, "body_sha256": "sha256:13573e63d0210438872617dd1c7c0d5f12685c9587cb37dbbf2f5d455c738115", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00", "source_path": "examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_global_controller_sso_egress:example:data-source", "parent_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:examples", "path": "documentation/data-sources/network_global_controller_sso_egress/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_global_controller_sso_egress", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1200211013301333-1222212313103000-2302100211003332-3011222110113330-2123221303110130-3332332310330221-3221212022321122-0313131012302212", "registry_path": "docs/guides/data-sources--network_global_controller_sso_egress--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_controller_sso_egress/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_global_controller_sso_egress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/examples/)
- [xcsh_network_global_controller_sso_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/)
