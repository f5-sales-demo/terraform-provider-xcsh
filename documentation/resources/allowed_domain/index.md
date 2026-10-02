---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "Manages allowed domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["allowed domain"], "body_bytes": 1553, "body_sha256": "sha256:7cc57e0d99e02f77b2ff2e36a79d3af49d381176e4e6e686ebbf2374d0427abb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:allowed_domain:reference", "xcsh-docs:resources:allowed_domain:examples", "xcsh-docs:resources:allowed_domain:import", "xcsh-docs:resources:allowed_domain:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:allowed_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/allowed_domain/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300", "registry_path": "docs/resources/allowed_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages allowed domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_allowed_domain

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

## Root configuration

Required root properties: `allowed_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/timeouts/)
