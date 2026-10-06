---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_subnet."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1029, "body_sha256": "sha256:78f15337981ac75f9ea8835ac6d48a3d33cb1ddbf56b18c4c8f1b697d199c6b6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0", "source_path": "examples/resources/xcsh_subnet/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:subnet:example:resource", "parent_id": "xcsh-docs:resources:subnet:examples", "path": "documentation/resources/subnet/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-0203011010313201-2233101131233311-2003031210233013-0312222212231323-3203110323201001-2231000231202022-2110212322310003-3131202203002022", "registry_path": "docs/guides/resources--subnet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
