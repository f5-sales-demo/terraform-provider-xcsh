---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fleet."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1033, "body_sha256": "sha256:147b34902200358bddfa30cd040d5e89481d034dd868fe2f7dd7cf50e3e3c92c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4d32cab1b63bbbd4184a815cf7049069724ae120ad87141da3ed9def32e6d8f3", "source_path": "examples/resources/xcsh_fleet/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fleet:example:resource", "parent_id": "xcsh-docs:resources:fleet:examples", "path": "documentation/resources/fleet/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2212131230021011-2120022302100122-3213222103232310-2322313112333112-0111200130300320-2002101113202230-0202111303112333-1310122332322322", "registry_path": "docs/guides/resources--fleet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fleet/resource.tf`; digest `sha256:4d32cab1b63bbbd4184a815cf7049069724ae120ad87141da3ed9def32e6d8f3`.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```
