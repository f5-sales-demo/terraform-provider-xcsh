---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_gen_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1041, "body_sha256": "sha256:8189197a0e6ee8a429a3e45957240eef3875a36449a0e405d83148ae806c1815", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eea28da7e92ac8598036f86f1095005a92f369536d4496a75262a4b141db489a", "source_path": "examples/resources/xcsh_alert_gen_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_gen_policy:example:resource", "parent_id": "xcsh-docs:resources:alert_gen_policy:examples", "path": "documentation/resources/alert_gen_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3222011002002223-1330220300113113-0131301101312110-0102323310112300-2332203333312303-2321210023123022-1130211011232310-2231110332222322", "registry_path": "docs/guides/resources--alert_gen_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_alert_gen_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_gen_policy/resource.tf`; digest `sha256:eea28da7e92ac8598036f86f1095005a92f369536d4496a75262a4b141db489a`.

```terraform
# AlertGenPolicy Resource Example
# Manages Alert Generation Policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertGenPolicy configuration
resource "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}
```
