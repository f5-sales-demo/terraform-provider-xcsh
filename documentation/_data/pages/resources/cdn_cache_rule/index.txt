---
page_title: "xcsh_cdn_cache_rule"
subcategory: ""
description: "Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification. configuration."
xcsh_docs: {"aliases": ["cdn cache rule"], "body_bytes": 1604, "body_sha256": "sha256:38f89ad102db06b62826fee69cdde8f989806b76ea85c4b51886beb31b2e3984", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:reference", "xcsh-docs:resources:cdn_cache_rule:examples", "xcsh-docs:resources:cdn_cache_rule:import", "xcsh-docs:resources:cdn_cache_rule:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cdn_cache_rule/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001", "registry_path": "docs/resources/cdn_cache_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/lifecycle/timeouts/)
