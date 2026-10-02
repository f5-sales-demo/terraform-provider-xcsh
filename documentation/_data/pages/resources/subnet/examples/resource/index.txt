---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_subnet."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1233, "body_sha256": "sha256:f6a9f7b3bdb5e9da363772303d6323059255fecb95975fd96a9e40054d26f4a4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0", "source_path": "examples/resources/xcsh_subnet/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:subnet:example:resource", "parent_id": "xcsh-docs:resources:subnet:examples", "path": "documentation/resources/subnet/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203011010313201-2233101131233311-2003031210233013-0312222212231323-3203110323201001-2231000231202022-2110212322310003-3131202203002022", "registry_path": "docs/guides/resources--subnet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_subnet/resource.tf`; digest `sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0`.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/examples/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
