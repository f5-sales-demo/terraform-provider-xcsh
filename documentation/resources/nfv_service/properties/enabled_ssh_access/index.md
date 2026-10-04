---
page_title: "enabled_ssh_access"
subcategory: ""
description: "SSH based configuration."
xcsh_docs: {"aliases": ["enabled ssh access"], "body_bytes": 4112, "body_sha256": "sha256:bc84eb13d42992604da46ff8b80e2cc33caf597c9058540e766e7b745e3695a3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/enabled_ssh_access/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_sli,advertise_on_slo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_sli,advertise_on_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_sli,advertise_on_slo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_slo,advertise_on_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_sli,advertise_on_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:ConflictingObjectAttributes:advertise_on_slo,advertise_on_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "type": "conflicts"}, {"anchor": "schema-enabled_ssh_access--domain_suffix", "enforcement": "provider-schema", "group": "enabled_ssh_access:RequiredObjectAttributes:domain_suffix,node_ssh_ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enabled_ssh_access:RequiredObjectAttributes:domain_suffix,node_ssh_ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled_ssh_access"], "schema_version": 1, "sections": [{"aliases": ["enabled ssh access advertise on sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access advertise on slo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access advertise on slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access domain suffix"], "anchor": "schema-enabled_ssh_access--domain_suffix", "description": "Domain suffix will be used along with node name to form the hostname for SSH node management.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "domain_suffix"], "syntax": "attribute", "type": "string"}, {"aliases": ["enabled ssh access node ssh ports"], "anchor": "section", "description": "Enter TCP port and node name per node.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-enabled_ssh_access--node_ssh_ports--node_name", "enforcement": "provider-schema", "group": "enabled_ssh_access.node_ssh_ports:RequiredListObjectAttributes:node_name,ssh_port", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "type": "requires"}, {"anchor": "schema-enabled_ssh_access--node_ssh_ports--ssh_port", "enforcement": "provider-schema", "group": "enabled_ssh_access.node_ssh_ports:RequiredListObjectAttributes:node_name,ssh_port", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "type": "requires"}], "schema_path": ["enabled_ssh_access", "node_ssh_ports"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/enabled_ssh_access/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SSH based configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled_ssh_access

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- enabled_ssh_access

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enabled ssh access.

Upstream description:

SSH based configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domain_suffix",
    "node_ssh_ports"),
  validators.ConflictingObjectAttributes("advertise_on_sli",
    "advertise_on_slo"),
  validators.ConflictingObjectAttributes("advertise_on_sli",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_slo",
    "advertise_on_slo_sli")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_sli\",\"advertise_on_slo\",\"advertise_on_slo_sli\"]"
}
```

Terraform syntax:

```terraform
enabled_ssh_access {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_on_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_sli/): complete subsection reference.

- [advertise_on_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_slo/): complete subsection reference.

- [advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_slo_sli/): complete subsection reference.

<a id="schema-enabled_ssh_access--domain_suffix"></a>

### domain_suffix property

Type: `"string"`. Optional.

Domain suffix will be used along with node name to form the hostname for SSH node management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [node_ssh_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/): complete subsection reference.

## Next pages

- [enabled_ssh_access.advertise_on_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_sli/)
- [enabled_ssh_access.advertise_on_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_slo/)
- [enabled_ssh_access.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/advertise_on_slo_sli/)
- [enabled_ssh_access.node_ssh_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
