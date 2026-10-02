---
page_title: "custom_network_config.global_network_list"
subcategory: ""
description: "List of global network connections."
xcsh_docs: {"aliases": ["custom network config global network list"], "body_bytes": 1966, "body_sha256": "sha256:ee7cd6c1ac1a27d0e3bc1b206812818997f3f84c185160c47390b682edad6ac7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "documentation/resources/voltstack_site/properties/custom_network_config/global_network_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3331330211330330-0322011301330323-1010231000103122-0030200003111010-1033201021201031-3323003112130013-2312331321222302-3033110210212020", "registry_path": "docs/guides/resources--voltstack_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.global_network_list:RequiredObjectAttributes:global_network_connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "global_network_list"], "schema_version": 1, "sections": [{"aliases": ["global network connections"], "anchor": "section", "description": "Global network connections.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr", "type": "conflicts"}], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- custom_network_config.global_network_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
