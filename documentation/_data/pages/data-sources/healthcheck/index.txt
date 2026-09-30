---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "xcsh_healthcheck for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1405, "body_sha256": "sha256:023e8462752bef94ea5e429283a0dc4a1186fe7c7f15740c2e2715f468533756", "child_ids": ["xcsh-docs:data-sources:healthcheck:reference", "xcsh-docs:data-sources:healthcheck:examples"], "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:fundamentals", "parent_id": null, "path": "documentation/data-sources/healthcheck/index.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_healthcheck for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_healthcheck

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Healthcheck by name
data "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"
}

output "healthcheck_id" {
  value = data.xcsh_healthcheck.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/examples/)
