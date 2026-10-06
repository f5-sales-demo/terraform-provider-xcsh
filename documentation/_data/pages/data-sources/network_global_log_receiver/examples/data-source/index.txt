---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_global_log_receiver."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1267, "body_sha256": "sha256:92ee5df0e443d2cab1bf0ef178aff19aef3d0f3137ccf26277b10b7ac3762fec", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e6481bf54911d3f66c740203a948d54d5c759ec25e22c585e36cf6e6833c987c", "source_path": "examples/data-sources/xcsh_network_global_log_receiver/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_global_log_receiver:example:data-source", "parent_id": "xcsh-docs:data-sources:network_global_log_receiver:examples", "path": "documentation/data-sources/network_global_log_receiver/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2023332312132030-2000303032010022-0221201301001233-3322023311121200-1122002033203110-0111320333210323-2313211303221202-0232212201013121", "registry_path": "docs/guides/data-sources--network_global_log_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_log_receiver/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_network_global_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_log_receiver/data-source.tf`; digest `sha256:e6481bf54911d3f66c740203a948d54d5c759ec25e22c585e36cf6e6833c987c`.

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

data "xcsh_network_global_log_receiver" "receivers" {}

# This example chooses TLS syslog on TCP 6514. The manifest supplies only
# destinations; choose the port required by the configured log receiver.
output "tls_syslog_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 6514
    destinations = data.xcsh_network_global_log_receiver.receivers.cidr_blocks
  }
}
```
