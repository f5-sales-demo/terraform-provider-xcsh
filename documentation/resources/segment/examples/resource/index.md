---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_segment."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 979, "body_sha256": "sha256:6dfa74be8544d93fe50a14272ca6afa0c1dc6a1bd84531abd5ee72be98e4ce22", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232", "source_path": "examples/resources/xcsh_segment/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:segment:example:resource", "parent_id": "xcsh-docs:resources:segment:examples", "path": "documentation/resources/segment/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0231123230002013-3232012203311111-1201213111100111-3033333033132300-0301320201032003-2322333203120010-2022201300222220-3203323130021100", "registry_path": "docs/guides/resources--segment--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_segment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["segmentCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_segment/resource.tf`; digest `sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232`.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```
