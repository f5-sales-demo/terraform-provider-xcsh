---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "Manages Mitigated Domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["mitigated domain"], "body_bytes": 1581, "body_sha256": "sha256:b39f8a89f19272b5a254692e985061d488fa9a20627a75b26f36ab6cc3d00a14", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:mitigated_domain:reference", "xcsh-docs:resources:mitigated_domain:examples", "xcsh-docs:resources:mitigated_domain:import", "xcsh-docs:resources:mitigated_domain:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:mitigated_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/mitigated_domain/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102", "registry_path": "docs/resources/mitigated_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages Mitigated Domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_mitigated_domain

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```

## Root configuration

Required root properties: `mitigated_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/timeouts/)
