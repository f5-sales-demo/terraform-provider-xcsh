---
page_title: "ddos_profile"
subcategory: ""
description: "DDoS Protection Rule for DNS."
xcsh_docs: {"aliases": ["ddos profile"], "body_bytes": 2049, "body_sha256": "sha256:86978f25a52d32445936dfe0e066791c9dba8b15fe1aa3c69e2bbc55067dcbcc", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "documentation/resources/dns_proxy/properties/ddos_profile/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2110003101023121-2222231331102321-2213230320013033-1313102003322220-0010011312320132-2221021103130201-1303331100212311-3230313220231100", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "sections": [{"aliases": ["ddos profile disable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos profile enable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "DDoS Protection Rule for DNS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- ddos_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

DDoS Protection Rule for DNS.

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

- [disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/): complete subsection reference.

- [enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/): complete subsection reference.

## Next pages

- [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/)
- [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
