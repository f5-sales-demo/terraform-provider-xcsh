---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_crl."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1068, "body_sha256": "sha256:0efcab36d59ce57eb963457d3df738ba9dd47a6687c34e1b409d5a86932e2f88", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3", "source_path": "examples/resources/xcsh_crl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:crl:example:resource", "parent_id": "xcsh-docs:resources:crl:examples", "path": "documentation/resources/crl/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1222233233230320-2223120312320112-2110210320101320-1002302222132120-0230233333021011-3312133333020310-0132322131300302-2012300022303313", "registry_path": "docs/guides/resources--crl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_crl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["crlCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_crl/resource.tf`; digest `sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3`.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```
