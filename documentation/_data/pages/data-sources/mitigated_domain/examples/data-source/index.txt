---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_mitigated_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1331, "body_sha256": "sha256:4c88c91fc8940393673f8efd7a091c0e422d7c7b770b4ee3774eee328e94b8b4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:mitigated_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16", "source_path": "examples/data-sources/xcsh_mitigated_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:mitigated_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:mitigated_domain:examples", "path": "documentation/data-sources/mitigated_domain/examples/data-source/index.md", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/mitigated_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_mitigated_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_mitigated_domain/data-source.tf`; digest `sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16`.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/)
