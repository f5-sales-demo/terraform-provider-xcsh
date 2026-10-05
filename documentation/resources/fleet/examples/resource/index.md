---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fleet."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1234, "body_sha256": "sha256:82b2fbb1f01c402606e9b49a1bc20ac025e65911b35387674d294471eff47c8c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4d32cab1b63bbbd4184a815cf7049069724ae120ad87141da3ed9def32e6d8f3", "source_path": "examples/resources/xcsh_fleet/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fleet:example:resource", "parent_id": "xcsh-docs:resources:fleet:examples", "path": "documentation/resources/fleet/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2212131230021011-2120022302100122-3213222103232310-2322313112333112-0111200130300320-2002101113202230-0202111303112333-1310122332322322", "registry_path": "docs/guides/resources--fleet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/examples/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
