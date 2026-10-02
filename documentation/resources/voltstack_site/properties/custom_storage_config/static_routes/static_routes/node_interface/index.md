---
page_title: "custom_storage_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["custom storage config static routes static routes node interface"], "body_bytes": 2197, "body_sha256": "sha256:5ceb07dc6015f9aa25b9b77992fccbdb9a52c5e59993f47ea238d9ef6577c947", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface:list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "node_interface", "list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/)
- [custom_storage_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/)
- custom_storage_config.static_routes.static_routes.node_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [custom_storage_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/list/)
- [custom_storage_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
