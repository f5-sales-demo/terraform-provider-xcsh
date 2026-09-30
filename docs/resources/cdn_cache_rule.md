---
page_title: "xcsh_cdn_cache_rule"
subcategory: ""
description: "xcsh_cdn_cache_rule for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1316, "body_sha256": "sha256:2ebc04fac7e97cf419c2dd475cb7932b2a40e62461df38d986634059095c553c", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:fundamentals", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:reference", "xcsh-docs:resources:cdn_cache_rule:examples", "xcsh-docs:resources:cdn_cache_rule:import", "xcsh-docs:resources:cdn_cache_rule:timeouts"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:fundamentals", "parent_id": null, "path": "docs/resources/cdn_cache_rule.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cdn_cache_rule for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_cdn_cache_rule

Breadcrumbs:

- xcsh_cdn_cache_rule

Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cdn_cache_rule--reference.md)
- [Examples](../guides/resources--cdn_cache_rule--examples.md)
- [Import](../guides/resources--cdn_cache_rule--import.md)
- [Timeouts](../guides/resources--cdn_cache_rule--timeouts.md)
