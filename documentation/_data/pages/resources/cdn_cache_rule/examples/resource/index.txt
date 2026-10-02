---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1290, "body_sha256": "sha256:94ec9a66b51b9f36de56424a9088cac61cd8f7e3589cc62875d67aeb1d0a1405", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a602b3888b8d6105a3cf057267d0d93b04e1fb4a8ec89a85c41a5bb720a38a3e", "source_path": "examples/resources/xcsh_cdn_cache_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cdn_cache_rule:example:resource", "parent_id": "xcsh-docs:resources:cdn_cache_rule:examples", "path": "documentation/resources/cdn_cache_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0212021203112122-1321202003113303-1102113122112103-1201233203003222-2123203031122302-0022223302321100-3001231230330001-1311112111231021", "registry_path": "docs/guides/resources--cdn_cache_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cdn_cache_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_cache_rule/resource.tf`; digest `sha256:a602b3888b8d6105a3cf057267d0d93b04e1fb4a8ec89a85c41a5bb720a38a3e`.

```terraform
# CDNCacheRule Resource Example
# Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNCacheRule configuration
resource "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/examples/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
