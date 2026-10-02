---
page_title: "xcsh_site"
subcategory: "Infrastructure"
description: "Manages a Site resource in F5 Distributed Cloud for get of site. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["site"], "body_bytes": 1398, "body_sha256": "sha256:6593a08ef04f3a74fa97fd5fd9a638c0c79feb7adb008090fa275bbf0d05a9f4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:reference", "xcsh-docs:data-sources:site:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3103121123130212-2311101033212231-0302320211230210-2003212033020122-0110200112201132-1223021030222031-3113123130220130-2002110333013120", "registry_path": "docs/data-sources/site.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:virtual_site:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Site resource in F5 Distributed Cloud for get of site. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site

Breadcrumbs:

- xcsh_site

Manages a Site resource in F5 Distributed Cloud for get of site. configuration. (read-only data
source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `virtual_site`.

- virtual_site: Logical grouping of physical sites

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/examples/)
