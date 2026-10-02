---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_gen_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1275, "body_sha256": "sha256:6a87d4cae118e9aa23f199e616cdbe464d5c1cab1547bedf00219c60dd9dd487", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eea28da7e92ac8598036f86f1095005a92f369536d4496a75262a4b141db489a", "source_path": "examples/resources/xcsh_alert_gen_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_gen_policy:example:resource", "parent_id": "xcsh-docs:resources:alert_gen_policy:examples", "path": "documentation/resources/alert_gen_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3222011002002223-1330220300113113-0131301101312110-0102323310112300-2332203333312303-2321210023123022-1130211011232310-2231110332222322", "registry_path": "docs/guides/resources--alert_gen_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_alert_gen_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/examples/)
- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/)
