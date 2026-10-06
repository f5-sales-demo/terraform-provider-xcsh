---
page_title: "Resource"
subcategory: "Identity"
description: "Resource for xcsh_token."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 974, "body_sha256": "sha256:e3797bbd737cc8d96cdffa3af8882e25c846889224291a9d8b2ca259aa1e7ab5", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9", "source_path": "examples/resources/xcsh_token/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:token:example:resource", "parent_id": "xcsh-docs:resources:token:examples", "path": "documentation/resources/token/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2101330310331202-0332121202233333-3312120320221103-2222113332330330-0203101000211121-0112012123211223-2321302231333312-3120102111310332", "registry_path": "docs/guides/resources--token--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_token.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tokenCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_token/resource.tf`; digest `sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9`.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```
