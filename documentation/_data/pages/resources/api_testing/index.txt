---
page_title: "xcsh_api_testing"
subcategory: ""
description: "xcsh_api_testing for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1485, "body_sha256": "sha256:6d38954883209e9dc9facdc1dbc417b43823164397c21745b7e5949b20c9116f", "child_ids": ["xcsh-docs:resources:api_testing:reference", "xcsh-docs:resources:api_testing:examples", "xcsh-docs:resources:api_testing:import", "xcsh-docs:resources:api_testing:timeouts"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:fundamentals", "parent_id": null, "path": "documentation/resources/api_testing/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_testing for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_testing

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/lifecycle/timeouts/)
