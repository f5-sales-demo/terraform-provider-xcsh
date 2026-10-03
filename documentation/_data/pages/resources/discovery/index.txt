---
page_title: "xcsh_discovery"
subcategory: ""
description: "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration."
xcsh_docs: {"aliases": ["discovery"], "body_bytes": 1642, "body_sha256": "sha256:55bed9d6f7a02b52724f635c2c5b72d1072a900b57090dbe9b0303e8b4b0997e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:reference", "xcsh-docs:resources:discovery:examples", "xcsh-docs:resources:discovery:import", "xcsh-docs:resources:discovery:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/discovery/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202", "registry_path": "docs/resources/discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/lifecycle/timeouts/)
