---
page_title: "ddos_profile"
subcategory: ""
description: "BIG-IP DDoS Protection Rules."
xcsh_docs: {"aliases": ["ddos profile"], "body_bytes": 1534, "body_sha256": "sha256:0fc3040f293bb6c7191e377d03b46f68cf46d92b8aeabf742e681f2c95024deb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:enable_ddos_mitigation"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/ddos_profile/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3223222331220311-2021021213301233-3110010022113211-1113021011323210-1100223221332333-3200220332110002-3301021103103321-3002222130011323", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:disable_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:enable_ddos_mitigation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "sections": [{"aliases": ["ddos profile disable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:disable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos profile enable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:enable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "BIG-IP DDoS Protection Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- ddos_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Additional upstream details:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/ddos_profile/disable_ddos_mitigation/): complete subsection reference.

- [enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/ddos_profile/enable_ddos_mitigation/): complete subsection reference.
