---
page_title: "xcsh_malicious_user_mitigation"
subcategory: ""
description: "Reads Malicious User Mitigation information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["malicious user mitigation"], "body_bytes": 1458, "body_sha256": "sha256:72276248e56a2986a4b040872f78fe481bb10425b7a6ba5b06bd569474fb2310", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:reference", "xcsh-docs:data-sources:malicious_user_mitigation:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/malicious_user_mitigation/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211", "registry_path": "docs/data-sources/malicious_user_mitigation.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reads Malicious User Mitigation information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_malicious_user_mitigation

Breadcrumbs:

- xcsh_malicious_user_mitigation

Reads Malicious User Mitigation information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/examples/)
