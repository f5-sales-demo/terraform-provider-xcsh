---
page_title: "local_ip.ip_address.virtual_network_type"
subcategory: ""
description: "Different types of virtual networks understood by the system."
xcsh_docs: {"aliases": ["local ip ip address virtual network type"], "body_bytes": 2882, "body_sha256": "sha256:636542a9bedf8df930ef9ab73438034d50d8d2bc02c018493b83b8e1fb99e824", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:site_local,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:site_local,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "schema_version": 1, "sections": [{"aliases": ["public"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "public"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "site_local"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local inside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "site_local_inside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Different types of virtual networks understood by the system.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.virtual_network_type

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- local_ip.ip_address.virtual_network_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Different types of virtual networks understood by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("public",
    "site_local"),
  validators.ConflictingObjectAttributes("public",
    "site_local_inside"),
  validators.ConflictingObjectAttributes("site_local",
    "site_local_inside")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

Terraform syntax:

```terraform
virtual_network_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/): complete subsection reference.

- [site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/): complete subsection reference.

- [site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/): complete subsection reference.

## Next pages

- [local_ip.ip_address.virtual_network_type.public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/)
- [local_ip.ip_address.virtual_network_type.site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/)
- [local_ip.ip_address.virtual_network_type.site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
