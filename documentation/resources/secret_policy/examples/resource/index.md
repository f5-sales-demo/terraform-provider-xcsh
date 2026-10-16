---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_secret_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1347, "body_sha256": "sha256:301bf87f827f72a20d5cdd5081e1602a589cd768817cf7ecb6bb4471ca876f47", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f59eb002c8cb50585be56dfb7682ee9d3bd976a6ce27601c7cb63562b64c6c5f", "source_path": "examples/resources/xcsh_secret_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:secret_policy:example:resource", "parent_id": "xcsh-docs:resources:secret_policy:examples", "path": "documentation/resources/secret_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3331221202302033-1321203231013113-2131303202012320-2312133000103013-3010222033113123-2320221223301003-3301300302210021-0012212322100010", "registry_path": "docs/guides/resources--secret_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Resource for xcsh_secret_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["secret_policyCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_secret_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_policy/resource.tf`; digest `sha256:f59eb002c8cb50585be56dfb7682ee9d3bd976a6ce27601c7cb63562b64c6c5f`.

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
