---
page_title: "xcsh_securemesh_site"
subcategory: ""
description: "Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security."
xcsh_docs: {"aliases": ["securemesh site"], "body_bytes": 1399, "body_sha256": "sha256:979b38292fc2712b42443b1c18aece3b79851d806aa8da308ec1ba2e602a3d70", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:reference", "xcsh-docs:data-sources:securemesh_site:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/securemesh_site/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230", "registry_path": "docs/data-sources/securemesh_site.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/examples/)
