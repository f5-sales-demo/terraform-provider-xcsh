---
page_title: "xcsh_crl"
subcategory: ""
description: "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration."
xcsh_docs: {"aliases": ["crl"], "body_bytes": 1648, "body_sha256": "sha256:0a921e297e39b45012107df4f1dd0d8c672853fdbeea231c05fac5b8ed779339", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:crl:reference", "xcsh-docs:resources:crl:examples", "xcsh-docs:resources:crl:import", "xcsh-docs:resources:crl:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "id": "xcsh-docs:resources:crl:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/crl/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010", "registry_path": "docs/resources/crl.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["crlCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_crl

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

## Root configuration

Required root properties: `name`, `namespace`, `refresh_interval`, `server_address`, `server_port`, `timeout`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/lifecycle/timeouts/)
