---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "Manages Mitigated Domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["mitigated domain"], "body_bytes": 1594, "body_sha256": "sha256:fb9080d00b3e42ba294a290cbb01679732fdde77d4936e0a380e2f46f9b6aa5e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:mitigated_domain:reference", "xcsh-docs:resources:mitigated_domain:examples", "xcsh-docs:resources:mitigated_domain:import", "xcsh-docs:resources:mitigated_domain:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:mitigated_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/mitigated_domain/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102", "registry_path": "docs/resources/mitigated_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Manages Mitigated Domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/timeouts/)
