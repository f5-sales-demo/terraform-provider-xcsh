---
page_title: "xcsh_securemesh_site"
subcategory: ""
description: "Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security."
xcsh_docs: {"aliases": ["securemesh site"], "body_bytes": 1727, "body_sha256": "sha256:ab42d9c4934668a2e6587396d9031ffa4973346455a88690d891044c3dab67e8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:reference", "xcsh-docs:resources:securemesh_site:examples", "xcsh-docs:resources:securemesh_site:import", "xcsh-docs:resources:securemesh_site:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/securemesh_site/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330", "registry_path": "docs/resources/securemesh_site.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_securemesh_site

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/lifecycle/timeouts/)
