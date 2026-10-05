---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1327, "body_sha256": "sha256:a1b9e2c551fa8f92e466fb8394a2481dcf7740b765b2510ab65584ae1bc88748", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:07fdb3a832901b7ce2b8f7a3292052a09eb46eebbc60fc38d6035bae249459aa", "source_path": "examples/resources/xcsh_cdn_purge_command/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cdn_purge_command:example:resource", "parent_id": "xcsh-docs:resources:cdn_purge_command:examples", "path": "documentation/resources/cdn_purge_command/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3333202021103021-2310121310221301-3233123123101103-2031330021132323-0022203300020120-3201221132110221-3210103202020020-1210220303100112", "registry_path": "docs/guides/resources--cdn_purge_command--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_cdn_purge_command.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/examples/)
- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/)
