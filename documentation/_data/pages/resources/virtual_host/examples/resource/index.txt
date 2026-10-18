---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_virtual_host."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1021, "body_sha256": "sha256:9601cd64ba8ef8880565162c26d70314f11fbb927ffc080c438b59e0216ef218", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7132f3f3fb88713f102679821dabd8f6cf4e11c9765e5f1be76eb1b01ecae4a0", "source_path": "examples/resources/xcsh_virtual_host/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_host:example:resource", "parent_id": "xcsh-docs:resources:virtual_host:examples", "path": "documentation/resources/virtual_host/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0021210100211230-0220022122312333-1311203322311022-1321013213023310-1022123303321020-3333121131313222-3103010321300233-3323002200013011", "registry_path": "docs/guides/resources--virtual_host--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/examples/resource/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Resource for xcsh_virtual_host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_host/resource.tf`; digest `sha256:7132f3f3fb88713f102679821dabd8f6cf4e11c9765e5f1be76eb1b01ecae4a0`.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```
