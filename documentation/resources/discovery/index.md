---
page_title: "xcsh_discovery"
subcategory: ""
description: "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration."
xcsh_docs: {"aliases": ["discovery"], "body_bytes": 1642, "body_sha256": "sha256:55bed9d6f7a02b52724f635c2c5b72d1072a900b57090dbe9b0303e8b4b0997e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:reference", "xcsh-docs:resources:discovery:examples", "xcsh-docs:resources:discovery:import", "xcsh-docs:resources:discovery:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/discovery/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202", "registry_path": "docs/resources/discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
