---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_application_profiles."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1069, "body_sha256": "sha256:158447fc0fd91a83547410006d328a7236b19c76dab75e1da78837f75bc4a278", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f", "source_path": "examples/resources/xcsh_application_profiles/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:application_profiles:example:resource", "parent_id": "xcsh-docs:resources:application_profiles:examples", "path": "documentation/resources/application_profiles/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0320321111331220-0331023100121133-3000212131200102-2302232203333031-3221123230331223-2102011000000010-1100112021001032-2031132130301202", "registry_path": "docs/guides/resources--application_profiles--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_application_profiles/resource.tf`; digest `sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f`.

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
