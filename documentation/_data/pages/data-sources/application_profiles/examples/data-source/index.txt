---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_application_profiles."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1131, "body_sha256": "sha256:44ad164eda880db87f74adce87e5b61678e388d850f6ea11d571976e3ff15eda", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc", "source_path": "examples/data-sources/xcsh_application_profiles/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:application_profiles:example:data-source", "parent_id": "xcsh-docs:data-sources:application_profiles:examples", "path": "documentation/data-sources/application_profiles/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0332101230303031-2111330303202023-1003220102010333-1302010131120302-3320322131011313-3010302133013302-3030302001113320-0213010310210231", "registry_path": "docs/guides/data-sources--application_profiles--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_application_profiles/data-source.tf`; digest `sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc`.

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
