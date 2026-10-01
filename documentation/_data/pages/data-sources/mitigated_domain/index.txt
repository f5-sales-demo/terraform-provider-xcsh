---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "xcsh_mitigated_domain for xcsh_mitigated_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1336, "body_sha256": "sha256:baf813c24f97b0e25fcff4f487aacbe64ec9fe100725b40f7ef917ade7e2b379", "child_ids": ["xcsh-docs:data-sources:mitigated_domain:reference", "xcsh-docs:data-sources:mitigated_domain:examples"], "collection_id": "xcsh-docs:data-sources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:mitigated_domain:fundamentals", "parent_id": null, "path": "documentation/data-sources/mitigated_domain/index.md", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_mitigated_domain for xcsh_mitigated_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_mitigated_domain

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
