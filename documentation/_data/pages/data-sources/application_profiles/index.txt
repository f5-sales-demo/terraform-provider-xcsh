---
page_title: "xcsh_application_profiles"
subcategory: ""
description: "Reads Application Profiles information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["application profiles"], "body_bytes": 1405, "body_sha256": "sha256:293705bc25e0e7ecbc32fb5fec46a0d06b71c980aaec1d48a856e00ff7cd67bc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:reference", "xcsh-docs:data-sources:application_profiles:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/application_profiles/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330", "registry_path": "docs/data-sources/application_profiles.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads Application Profiles information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["application_profilesCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_application_profiles

Breadcrumbs:

- xcsh_application_profiles

Reads Application Profiles information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ApplicationProfiles by name
data "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}

output "application_profiles_id" {
  value = data.xcsh_application_profiles.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/examples/)
