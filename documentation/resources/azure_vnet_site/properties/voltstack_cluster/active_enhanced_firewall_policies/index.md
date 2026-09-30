---
page_title: "voltstack_cluster.active_enhanced_firewall_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_enhanced_firewall_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2182, "body_sha256": "sha256:fca592d00c30586219172eab1bb8150a1a9c3e7db0b9fd9512fa224271dd1509", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["voltstack_cluster", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_enhanced_firewall_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
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

- [enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
