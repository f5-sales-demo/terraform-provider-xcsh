---
page_title: "Deny all"
subcategory: "Security"
description: "Deny all for xcsh_service_policy."
xcsh_docs: {"aliases": ["deny-all"], "body_bytes": 1070, "body_sha256": "sha256:7170fed17a487d99f11805cc8b6e432a53d48263020c9ef024a32a09e403b1ce", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:9da9688512df86906769c34d1e33802dc2db45bff21587219ced7d18b6172282", "source_path": "examples/resources/xcsh_service_policy/deny-all.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:deny-all", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "documentation/resources/service_policy/examples/deny-all/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1010220320132302-1032132231102012-0003122002201032-1230100323323112-3023322323032203-2100130100220132-2002320201331033-0211301312101222", "registry_path": "docs/guides/resources--service_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["deny-all"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/deny-all/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Deny all for xcsh_service_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Deny all

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
- Deny all

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-all.tf`; digest `sha256:9da9688512df86906769c34d1e33802dc2db45bff21587219ced7d18b6172282`.

```terraform
# DenyAll — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  # Deny all requests
  deny_all_requests = {}

  # Apply to any server
  any_server = {}
}
```
