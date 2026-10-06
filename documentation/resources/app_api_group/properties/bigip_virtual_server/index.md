---
page_title: "bigip_virtual_server"
subcategory: ""
description: "Set the scope of the API Group to a specific BIG-IP Virtual Server."
xcsh_docs: {"aliases": ["bigip virtual server"], "body_bytes": 1626, "body_sha256": "sha256:332cb05720e90547a154555ca0b9831f66de3fc6e5a002dffeb6d60e83da6e28", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:bigip_virtual_server", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "documentation/resources/app_api_group/properties/bigip_virtual_server/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1031222130211103-1021233123310101-3110112321222203-0031313230202331-2233320301312110-2022232123223302-2101022333132131-1321310320121211", "registry_path": "docs/guides/resources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bigip_virtual_server"], "schema_version": 1, "sections": [{"aliases": ["bigip virtual server bigip virtual server"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bigip_virtual_server--bigip_virtual_server--name", "enforcement": "provider-schema", "group": "bigip_virtual_server.bigip_virtual_server:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server", "type": "requires"}], "schema_path": ["bigip_virtual_server", "bigip_virtual_server"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Set the scope of the API Group to a specific BIG-IP Virtual Server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bigip_virtual_server

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- bigip_virtual_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

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

OneOf alternatives in this subsection:

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/#section)
- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/#section)
- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bigip_virtual_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/): complete subsection reference.
