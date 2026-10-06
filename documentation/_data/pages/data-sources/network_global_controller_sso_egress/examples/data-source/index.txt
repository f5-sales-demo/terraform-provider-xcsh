---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_global_controller_sso_egress."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1175, "body_sha256": "sha256:eaeec6670327f7427ff600048fbecdd92f9c1d8737994a771ade3e1f93d13bb9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00", "source_path": "examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_global_controller_sso_egress:example:data-source", "parent_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:examples", "path": "documentation/data-sources/network_global_controller_sso_egress/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_global_controller_sso_egress", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-1200211013301333-1222212313103000-2302100211003332-3011222110113330-2123221303110130-3332332310330221-3221212022321122-0313131012302212", "registry_path": "docs/guides/data-sources--network_global_controller_sso_egress--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_controller_sso_egress/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_network_global_controller_sso_egress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
