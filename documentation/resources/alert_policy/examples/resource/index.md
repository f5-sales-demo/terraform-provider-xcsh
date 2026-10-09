---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_alert_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1011, "body_sha256": "sha256:ea43b6054779da442ee1b85a689c3391e63df03c4d565c971bc6d29c99acc665", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e", "source_path": "examples/resources/xcsh_alert_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_policy:example:resource", "parent_id": "xcsh-docs:resources:alert_policy:examples", "path": "documentation/resources/alert_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1222003231100012-1132113102332331-2310331010301201-1020303102110002-3222210113331200-3021303323321120-0222023310120020-1101032113332303", "registry_path": "docs/guides/resources--alert_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_alert_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["alert_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_policy/resource.tf`; digest `sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e`.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```
