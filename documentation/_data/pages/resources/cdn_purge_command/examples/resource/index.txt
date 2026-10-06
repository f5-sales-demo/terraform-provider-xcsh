---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1090, "body_sha256": "sha256:10c3e8204217ec08f37a4fdd1e29d5d93c9d8e51c1a0f8d97ebab78c24504f0a", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:07fdb3a832901b7ce2b8f7a3292052a09eb46eebbc60fc38d6035bae249459aa", "source_path": "examples/resources/xcsh_cdn_purge_command/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cdn_purge_command:example:resource", "parent_id": "xcsh-docs:resources:cdn_purge_command:examples", "path": "documentation/resources/cdn_purge_command/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3333202021103021-2310121310221301-3233123123101103-2031330021132323-0022203300020120-3201221132110221-3210103202020020-1210220303100112", "registry_path": "docs/guides/resources--cdn_purge_command--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_cdn_purge_command.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_purge_command/resource.tf`; digest `sha256:07fdb3a832901b7ce2b8f7a3292052a09eb46eebbc60fc38d6035bae249459aa`.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```
