---
page_title: "enabled_ssh_access"
subcategory: ""
description: "SSH based configuration."
xcsh_docs: {"aliases": ["enabled ssh access"], "body_bytes": 3442, "body_sha256": "sha256:67a3d120348ab933956bc62b2cdd154e69638f0023e971c2c5473ccdd212a055", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:node_ssh_ports"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "documentation/data-sources/nfv_service/properties/enabled_ssh_access/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled_ssh_access"], "schema_version": 1, "sections": [{"aliases": ["advertise on sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise on slo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise on slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain suffix"], "anchor": "schema-enabled_ssh_access--domain_suffix", "description": "Domain suffix will be used along with node name to form the hostname for SSH node management.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "domain_suffix"], "syntax": "attribute", "type": "string"}, {"aliases": ["node ssh ports"], "anchor": "section", "description": "Enter TCP port and node name per node.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["enabled_ssh_access", "node_ssh_ports"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/enabled_ssh_access/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SSH based configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled_ssh_access

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- enabled_ssh_access

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for enabled ssh access.

Upstream description:

SSH based configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_sli\",\"advertise_on_slo\",\"advertise_on_slo_sli\"]"
}
```

## Direct properties

- [advertise_on_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_sli/): complete subsection reference.

- [advertise_on_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_slo/): complete subsection reference.

- [advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_slo_sli/): complete subsection reference.

<a id="schema-enabled_ssh_access--domain_suffix"></a>

### domain_suffix property

Type: `"string"`. Computed.

Domain suffix will be used along with node name to form the hostname for SSH node management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [node_ssh_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/): complete subsection reference.

## Next pages

- [enabled_ssh_access.advertise_on_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_sli/)
- [enabled_ssh_access.advertise_on_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_slo/)
- [enabled_ssh_access.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/advertise_on_slo_sli/)
- [enabled_ssh_access.node_ssh_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
