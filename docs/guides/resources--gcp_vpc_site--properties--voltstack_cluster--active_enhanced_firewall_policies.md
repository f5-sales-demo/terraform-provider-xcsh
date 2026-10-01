---
page_title: "voltstack_cluster.active_enhanced_firewall_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_enhanced_firewall_policies for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1899, "body_sha256": "sha256:55025ea2ad6551d2cbdac925f26ecd771f2d1bce6ad25038867111ad1871c7cc", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "path": "docs/guides/resources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_enhanced_firewall_policies for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [voltstack_cluster](resources--gcp_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enhanced_firewall_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [voltstack_cluster](resources--gcp_vpc_site--properties--voltstack_cluster.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
