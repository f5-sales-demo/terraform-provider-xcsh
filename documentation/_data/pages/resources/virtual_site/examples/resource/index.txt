---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_virtual_site."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1248, "body_sha256": "sha256:e747336fc9efd1c64d84ba6938d99e278a8dce64b9ab6bbf0011d6123c4f4fac", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872", "source_path": "examples/resources/xcsh_virtual_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_site:example:resource", "parent_id": "xcsh-docs:resources:virtual_site:examples", "path": "documentation/resources/virtual_site/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "virtual_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0212333203312330-3320301102212300-3032103331311011-3321002330020023-2003221033111101-3000301001100322-1220212313313310-1101030333323000", "registry_path": "docs/guides/resources--virtual_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_site/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_virtual_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_site/resource.tf`; digest `sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872`.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/examples/)
- [xcsh_virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/)
