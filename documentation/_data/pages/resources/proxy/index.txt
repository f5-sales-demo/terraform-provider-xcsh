---
page_title: "xcsh_proxy"
subcategory: ""
description: "Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification. configuration."
xcsh_docs: {"aliases": ["proxy"], "body_bytes": 1527, "body_sha256": "sha256:d25a88865d946015b8d132db3ca896f52f132667a328f73898c09b28d38208b9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:reference", "xcsh-docs:resources:proxy:examples", "xcsh-docs:resources:proxy:import", "xcsh-docs:resources:proxy:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332", "registry_path": "docs/resources/proxy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_proxy

Breadcrumbs:

- xcsh_proxy

Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/lifecycle/timeouts/)
