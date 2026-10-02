---
page_title: "xcsh_voltstack_site"
subcategory: ""
description: "Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites."
xcsh_docs: {"aliases": ["voltstack site"], "body_bytes": 1679, "body_sha256": "sha256:f2d93199403bf07697ae86e642ae44279d94d45d074ba969265080a8d3b02413", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:reference", "xcsh-docs:resources:voltstack_site:examples", "xcsh-docs:resources:voltstack_site:import", "xcsh-docs:resources:voltstack_site:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/voltstack_site/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113", "registry_path": "docs/resources/voltstack_site.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_voltstack_site

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/lifecycle/timeouts/)
