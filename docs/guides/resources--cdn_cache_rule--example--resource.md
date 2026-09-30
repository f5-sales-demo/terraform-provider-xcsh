---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 985, "body_sha256": "sha256:0d4e83ce5ad1ad5940aebca7a936031cb53d35edf59679de1ba7b1cad39c3db5", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a602b3888b8d6105a3cf057267d0d93b04e1fb4a8ec89a85c41a5bb720a38a3e", "source_path": "examples/resources/xcsh_cdn_cache_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cdn_cache_rule:example:resource", "parent_id": "xcsh-docs:resources:cdn_cache_rule:examples", "path": "docs/guides/resources--cdn_cache_rule--example--resource.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Examples](resources--cdn_cache_rule--examples.md)
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

- [Examples](resources--cdn_cache_rule--examples.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
