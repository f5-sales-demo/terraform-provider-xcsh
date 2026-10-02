---
page_title: "xcsh_nat_policy"
subcategory: ""
description: "Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,. configuration."
xcsh_docs: {"aliases": ["nat policy"], "body_bytes": 1646, "body_sha256": "sha256:90169e42ac2628e001a578313c49ab3b5a6bc5e0f81fb612cb26bd2d88ff3b03", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:reference", "xcsh-docs:resources:nat_policy:examples", "xcsh-docs:resources:nat_policy:import", "xcsh-docs:resources:nat_policy:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/nat_policy/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310", "registry_path": "docs/resources/nat_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nat_policy

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/lifecycle/timeouts/)
