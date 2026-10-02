---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_public_ip_binding."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1435, "body_sha256": "sha256:a6fdb6bfa9da48fe6bf2b9652c5eeeaf4439cc868c61c5c8a21d7d94aa83b6a2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:public_ip_binding:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:193afa89349b62b6fa18ffe2657f80b77590d2dd479611710090adfb88effdd9", "source_path": "examples/resources/xcsh_public_ip_binding/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:public_ip_binding:example:resource", "parent_id": "xcsh-docs:resources:public_ip_binding:examples", "path": "documentation/resources/public_ip_binding/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "public_ip_binding", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2223123313212300-3023032303202012-2012211231031331-1311020231001102-3213120012231210-1013001311303121-0003302020113300-2213222303130231", "registry_path": "docs/guides/resources--public_ip_binding--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/public_ip_binding/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_public_ip_binding.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_public_ip_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_public_ip_binding/resource.tf`; digest `sha256:193afa89349b62b6fa18ffe2657f80b77590d2dd479611710090adfb88effdd9`.

```terraform
# PublicIPBinding Resource Example
# Manage the regional virtual-site binding of an already allocated public IP.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PublicIPBinding configuration
resource "xcsh_public_ip_binding" "example" {
  name      = "example-public-ip-binding"
  namespace = "staging"

  expected_ip            = "example-value"
  virtual_site           = "example-value"
  virtual_site_namespace = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/examples/)
- [xcsh_public_ip_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/)
