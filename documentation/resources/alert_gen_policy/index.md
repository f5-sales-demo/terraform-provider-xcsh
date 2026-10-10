---
page_title: "xcsh_alert_gen_policy"
subcategory: ""
description: "Manages Alert Generation Policy in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["alert gen policy"], "body_bytes": 1548, "body_sha256": "sha256:5f8c93e7f81f86c85111e7126cc00151c685f8e449d10dc7019474a4128450ab", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_gen_policy:reference", "xcsh-docs:resources:alert_gen_policy:examples", "xcsh-docs:resources:alert_gen_policy:import", "xcsh-docs:resources:alert_gen_policy:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_gen_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/alert_gen_policy/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001", "registry_path": "docs/resources/alert_gen_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages Alert Generation Policy in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/lifecycle/timeouts/)
