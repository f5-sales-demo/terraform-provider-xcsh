---
page_title: "xcsh_site_signatures_update"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["site signatures update"], "body_bytes": 1297, "body_sha256": "sha256:f4fdb777161ac92730baa5d5240e4c58b39e7313c9b6f967e87a6295238719b7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:site_signatures_update:reference", "xcsh-docs:actions:site_signatures_update:examples", "xcsh-docs:actions:site_signatures_update:lifecycle"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_signatures_update:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/site_signatures_update/index.md", "product": "distributed-cloud", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "registry_anchor": "canonical-2121133033301232-0011100131212313-0013203021322030-2021023332100333-1321010310023012-3332003100213031-1333010220333010-1320230310323133", "registry_path": "docs/actions/site_signatures_update.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_signatures_update

Breadcrumbs:

- xcsh_site_signatures_update

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/lifecycle/)
