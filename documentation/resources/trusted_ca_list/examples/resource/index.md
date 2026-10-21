---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1086, "body_sha256": "sha256:8e10ff113a361a0beb261dedad9553f02dac929b881ae4419368aa52bf5acdc3", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328", "source_path": "examples/resources/xcsh_trusted_ca_list/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:trusted_ca_list:example:resource", "parent_id": "xcsh-docs:resources:trusted_ca_list:examples", "path": "documentation/resources/trusted_ca_list/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2103021213110323-1232123021100101-0100111023203031-2120302130122313-1013031001112213-0112213033332111-0122031310202201-0001213323302301", "registry_path": "docs/guides/resources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/examples/resource/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Resource for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```
