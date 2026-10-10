---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_api_testing."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1041, "body_sha256": "sha256:603e608165fe43183bf4027094ece84167039c7cda1cb61e76ddb081a3487ece", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd", "source_path": "examples/data-sources/xcsh_api_testing/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_testing:example:data-source", "parent_id": "xcsh-docs:data-sources:api_testing:examples", "path": "documentation/data-sources/api_testing/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3330001211231130-1030221022230232-2223323322000123-0133313130201001-0132022222302222-2221123013220202-0231300013303320-0121231221110120", "registry_path": "docs/guides/data-sources--api_testing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_api_testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_testingCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_testing/data-source.tf`; digest `sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd`.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```
