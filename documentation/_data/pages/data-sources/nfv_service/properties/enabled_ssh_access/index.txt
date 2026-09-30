---
page_title: "enabled_ssh_access"
subcategory: ""
description: "enabled_ssh_access for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 3343, "body_sha256": "sha256:7d28af3c94c32d3666c9325fc63d301dbec2cff1c8d7dd9c2d02b4d2e168c232", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_sli", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:advertise_on_slo_sli", "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access:node_ssh_ports"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:enabled_ssh_access", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "documentation/data-sources/nfv_service/properties/enabled_ssh_access/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["enabled_ssh_access"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/enabled_ssh_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enabled_ssh_access for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
