---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_application_profiles."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1383, "body_sha256": "sha256:55ea783aba40d0435ff996936a10413f5400f3cc60a7fd9315d28e78ec6909e2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc", "source_path": "examples/data-sources/xcsh_application_profiles/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:application_profiles:example:data-source", "parent_id": "xcsh-docs:data-sources:application_profiles:examples", "path": "documentation/data-sources/application_profiles/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0332101230303031-2111330303202023-1003220102010333-1302010131120302-3320322131011313-3010302133013302-3030302001113320-0213010310210231", "registry_path": "docs/guides/data-sources--application_profiles--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/examples/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
