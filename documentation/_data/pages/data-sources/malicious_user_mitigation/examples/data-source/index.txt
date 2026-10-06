---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1179, "body_sha256": "sha256:e3c5130b08577f72ed31e5898437246f0e03d5d8cf4dee95b107bcf6da227036", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:48191839abab419f70a338b5b737300e10e8068606e35e5a0829c82bf6fe5846", "source_path": "examples/data-sources/xcsh_malicious_user_mitigation/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:malicious_user_mitigation:example:data-source", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:examples", "path": "documentation/data-sources/malicious_user_mitigation/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2011013332031323-0002030011020021-2220122132030012-1011320320132210-2201121322030303-2010320300231203-0002323323202132-2323203033132212", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_malicious_user_mitigation/data-source.tf`; digest `sha256:48191839abab419f70a338b5b737300e10e8068606e35e5a0829c82bf6fe5846`.

```terraform
# MaliciousUserMitigation Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MaliciousUserMitigation by name
data "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}

output "malicious_user_mitigation_id" {
  value = data.xcsh_malicious_user_mitigation.example.id
}
```
