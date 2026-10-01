---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "xcsh_mitigated_domain for xcsh_mitigated_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1251, "body_sha256": "sha256:c702183d6d64053890f2a7c89c490ff3ee7928fc3b5b5995db35eb54902eb991", "canonical_id": "xcsh-docs:data-sources:mitigated_domain:fundamentals", "child_ids": ["xcsh-docs:data-sources:mitigated_domain:reference", "xcsh-docs:data-sources:mitigated_domain:examples"], "collection_id": "xcsh-docs:data-sources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:mitigated_domain:fundamentals", "parent_id": null, "path": "docs/data-sources/mitigated_domain.md", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_mitigated_domain for xcsh_mitigated_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/data-sources--mitigated_domain--reference.md)
- [Examples](../guides/data-sources--mitigated_domain--examples.md)
