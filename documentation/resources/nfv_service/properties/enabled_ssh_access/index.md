---
page_title: "enabled_ssh_access"
subcategory: ""
description: "SSH based configuration."
xcsh_docs: {"aliases": ["enabled ssh access"], "body_bytes": 2637, "body_sha256": "sha256:8115047e98030084fac7c470b13f24c9c68a843eef36825422f87ab6c27860d5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/enabled_ssh_access/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled_ssh_access"], "schema_version": 1, "sections": [{"aliases": ["enabled ssh access advertise on sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access advertise on slo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access advertise on slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "advertise_on_slo_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["enabled ssh access domain suffix"], "anchor": "schema-enabled_ssh_access--domain_suffix", "description": "Domain suffix will be used along with node name to form the hostname for SSH node management.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "domain_suffix"], "syntax": "attribute", "type": "string"}, {"aliases": ["enabled ssh access node ssh ports"], "anchor": "section", "description": "Enter TCP port and node name per node.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["enabled_ssh_access", "node_ssh_ports"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/enabled_ssh_access/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "SSH based configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
