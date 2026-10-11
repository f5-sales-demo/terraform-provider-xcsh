---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_allowed_domain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1054, "body_sha256": "sha256:f1f377787182362c82f3b653bd117f910c9c237ee73507a81689de2dbc36fc2b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e", "source_path": "examples/resources/xcsh_allowed_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:allowed_domain:example:resource", "parent_id": "xcsh-docs:resources:allowed_domain:examples", "path": "documentation/resources/allowed_domain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3131212011321233-2111320300320200-2323102123033320-2000223303201032-2010330322312102-3301211022032010-2232031213111110-2320112201230211", "registry_path": "docs/guides/resources--allowed_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_allowed_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_allowed_domain/resource.tf`; digest `sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e`.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```
