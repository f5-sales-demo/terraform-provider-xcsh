---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "Reads Allowed Domain information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["allowed domain"], "body_bytes": 1339, "body_sha256": "sha256:3dfa1203cbbc3e84332abd9cde62fb418830110a698893bb5e7f6bf340c8a4f8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:allowed_domain:reference", "xcsh-docs:data-sources:allowed_domain:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:allowed_domain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:allowed_domain:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/allowed_domain/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220", "registry_path": "docs/data-sources/allowed_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/allowed_domain/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads Allowed Domain information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_allowed_domain

Breadcrumbs:

- xcsh_allowed_domain

Reads Allowed Domain information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/examples/)
