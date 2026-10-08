---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nfv_service."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1022, "body_sha256": "sha256:1698d12988e81f062d84b76122ce574696a7070aa8a8fad652cf4f109bdba9d6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9", "source_path": "examples/resources/xcsh_nfv_service/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nfv_service:example:resource", "parent_id": "xcsh-docs:resources:nfv_service:examples", "path": "documentation/resources/nfv_service/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1213221123222201-0030223100121310-2210301222223030-2232002320212033-0031131003022211-2120300230211333-1333011030220133-3133220011303213", "registry_path": "docs/guides/resources--nfv_service--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_nfv_service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nfv_service/resource.tf`; digest `sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9`.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```
