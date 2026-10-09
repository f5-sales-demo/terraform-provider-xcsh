---
page_title: "xcsh_discovery"
subcategory: ""
description: "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration."
xcsh_docs: {"aliases": ["discovery"], "body_bytes": 1655, "body_sha256": "sha256:6a48eee905ba5bdb3c934b41813ab344427aa30a2cd34d400277155cc0bb0c27", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:reference", "xcsh-docs:resources:discovery:examples", "xcsh-docs:resources:discovery:import", "xcsh-docs:resources:discovery:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/discovery/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202", "registry_path": "docs/resources/discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["discoveryCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_discovery

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/lifecycle/timeouts/)
