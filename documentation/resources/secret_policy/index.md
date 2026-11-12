---
page_title: "xcsh_secret_policy"
subcategory: "Security"
description: "Owns a native Secret Policy with ordered inline ALLOW/DENY rules. Evaluation is FIRST_MATCH (top to bottom); no invented API algorithm field is sent. Deletion is soft: destroy remains pending while a tombstone exists and never repeats DELETE for a marked policy. Import namespace/name; legacy child references are"
xcsh_docs: {"aliases": ["adopt existing object", "import existing resource", "secret policy"], "body_bytes": 2115, "body_sha256": "sha256:f234511cb27efc8e946e6533e650574daed7abd13e53b4bb26b3ce3608876e70", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:secret_policy:reference", "xcsh-docs:resources:secret_policy:examples", "xcsh-docs:resources:secret_policy:import"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/secret_policy/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111", "registry_path": "docs/resources/secret_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Owns a native Secret Policy with ordered inline ALLOW/DENY rules. Evaluation is FIRST_MATCH (top to bottom); no invented API algorithm field is sent. Deletion is soft: destroy remains pending while a tombstone exists and never repeats DELETE for a marked policy. Import namespace/name; legacy child references are", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["secret_policyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_secret_policy

Breadcrumbs:

- xcsh_secret_policy

Owns a native Secret Policy with ordered inline ALLOW/DENY rules. Evaluation is FIRST\_MATCH (top to
bottom); no invented API algorithm field is sent. Deletion is soft: destroy remains pending while a
tombstone exists and never repeats DELETE for a marked policy. Import namespace/name; legacy child
references are normalized into inline rules. No secret or credential belongs in this resource.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.14.0"
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 15.8.3"
    }
  }
}

# Install the current published release containing xcsh_secret_policy.
# Keep the namespace independently owned. The application pod must carry app=demo-tls.
resource "xcsh_secret_policy" "application" {
  name                     = "demo-wingman-policy"
  namespace                = "demo-app"
  allow_f5xc               = true
  decrypt_cache_timeout    = "30s"
  rule_combining_algorithm = "FIRST_MATCH"
  rules = [
    { name = "allow-application", action = "ALLOW", client_selector = ["app=demo-tls"] },
    { name = "deny-all", action = "DENY" }
  ]
}
```

## Root configuration

Required root properties: `name`, `namespace`, `rules`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/lifecycle/import/)
