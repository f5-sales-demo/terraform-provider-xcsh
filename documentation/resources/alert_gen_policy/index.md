---
page_title: "xcsh_alert_gen_policy"
subcategory: ""
description: "Manages Alert Generation Policy in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["alert gen policy"], "body_bytes": 1535, "body_sha256": "sha256:71f7eeb1458fdb7a0de3bcc8624c05b5f3ec39c1e3083c2dc33c5291a3054a67", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_gen_policy:reference", "xcsh-docs:resources:alert_gen_policy:examples", "xcsh-docs:resources:alert_gen_policy:import", "xcsh-docs:resources:alert_gen_policy:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_gen_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/alert_gen_policy/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001", "registry_path": "docs/resources/alert_gen_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages Alert Generation Policy in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_alert_gen_policy

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/lifecycle/timeouts/)
