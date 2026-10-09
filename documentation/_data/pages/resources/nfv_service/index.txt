---
page_title: "xcsh_nfv_service"
subcategory: ""
description: "Manages new NFV service with configured parameters in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["nfv service"], "body_bytes": 1538, "body_sha256": "sha256:3e85c88d6440c9dd578688a9d3d52a2844484da9a5b6147847d2a898174d300d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:reference", "xcsh-docs:resources:nfv_service:examples", "xcsh-docs:resources:nfv_service:import", "xcsh-docs:resources:nfv_service:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/nfv_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222", "registry_path": "docs/resources/nfv_service.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Manages new NFV service with configured parameters in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nfv_service

Breadcrumbs:

- xcsh_nfv_service

Manages new NFV service with configured parameters in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/lifecycle/timeouts/)
