---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1054, "body_sha256": "sha256:90177814af45da4b8e35c116db9d50166710a5a69747629a36ac0a0c63526588", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4", "source_path": "examples/resources/xcsh_bot_infrastructure/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bot_infrastructure:example:resource", "parent_id": "xcsh-docs:resources:bot_infrastructure:examples", "path": "documentation/resources/bot_infrastructure/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2222320201313333-3321231120231130-1320121030130311-1021311232121221-2131002020323331-3203232230001003-1023233102213300-0323233122330130", "registry_path": "docs/guides/resources--bot_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/examples/resource/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Resource for xcsh_bot_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_infrastructure/resource.tf`; digest `sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4`.

```terraform
# BotInfrastructure Resource Example
# Manages Bot Infrastructure in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotInfrastructure configuration
resource "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}
```
