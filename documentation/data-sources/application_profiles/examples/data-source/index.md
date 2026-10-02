---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_application_profiles."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1383, "body_sha256": "sha256:55ea783aba40d0435ff996936a10413f5400f3cc60a7fd9315d28e78ec6909e2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc", "source_path": "examples/data-sources/xcsh_application_profiles/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:application_profiles:example:data-source", "parent_id": "xcsh-docs:data-sources:application_profiles:examples", "path": "documentation/data-sources/application_profiles/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0332101230303031-2111330303202023-1003220102010333-1302010131120302-3320322131011313-3010302133013302-3030302001113320-0213010310210231", "registry_path": "docs/guides/data-sources--application_profiles--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
