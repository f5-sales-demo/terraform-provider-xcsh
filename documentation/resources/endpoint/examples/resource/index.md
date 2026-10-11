---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_endpoint."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1045, "body_sha256": "sha256:b886ae3914f511716382bd6d3afc4f0b5f3a11f08caf7762d84776902330d1f3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e", "source_path": "examples/resources/xcsh_endpoint/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:endpoint:example:resource", "parent_id": "xcsh-docs:resources:endpoint:examples", "path": "documentation/resources/endpoint/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2212202010123110-2131002232331010-0321202122203111-2021233032121011-1211231100220030-2031111213233103-0132231032232021-2200211010231232", "registry_path": "docs/guides/resources--endpoint--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["endpointCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_endpoint/resource.tf`; digest `sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e`.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```
