---
page_title: "xcsh_application_profiles"
subcategory: ""
description: "Manages Application Profiles in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["application profiles"], "body_bytes": 1647, "body_sha256": "sha256:a7dcc4033a8447d14f307e0c43778abf50fd9dcfe492de952ce4a40ae0e2e96e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:reference", "xcsh-docs:resources:application_profiles:examples", "xcsh-docs:resources:application_profiles:import", "xcsh-docs:resources:application_profiles:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/application_profiles/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130", "registry_path": "docs/resources/application_profiles.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Manages Application Profiles in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["application_profilesCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_application_profiles

Breadcrumbs:

- xcsh_application_profiles

Manages Application Profiles in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/lifecycle/timeouts/)
