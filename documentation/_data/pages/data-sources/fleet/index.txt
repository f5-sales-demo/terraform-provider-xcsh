---
page_title: "xcsh_fleet"
subcategory: ""
description: "Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["fleet"], "body_bytes": 1278, "body_sha256": "sha256:3079ee185b3658fe92bea20fe0b38b7cd065091e5b657f4f5da945c9d279b830", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:reference", "xcsh-docs:data-sources:fleet:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/fleet/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222", "registry_path": "docs/data-sources/fleet.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fleet

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/examples/)
