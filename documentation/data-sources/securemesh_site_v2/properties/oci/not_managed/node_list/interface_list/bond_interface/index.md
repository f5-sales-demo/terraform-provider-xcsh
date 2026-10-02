---
page_title: "oci.not_managed.node_list.interface_list.bond_interface"
subcategory: ""
description: "Bond devices configuration for fleet."
xcsh_docs: {"aliases": ["oci not managed node list interface list bond interface"], "body_bytes": 7660, "body_sha256": "sha256:a45160c5c387eeb22ece9c33bce574638ffba76398f8abe15a98b5b02198556c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface:active_backup", "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface:lacp"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3102311321011003-1130113132323220-0222031322212102-0230111001200300-3203320302233030-1003021220223303-1330201303230212-3011122212022133", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface"], "schema_version": 1, "sections": [{"aliases": ["active backup"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface:active_backup", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "active_backup"], "syntax": "attribute", "type": "object"}, {"aliases": ["devices"], "anchor": "schema-oci--not_managed--node_list--interface_list--bond_interface--devices", "description": "Ethernet devices that will make up this bond.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "devices"], "syntax": "attribute", "type": "list"}, {"aliases": ["lacp"], "anchor": "section", "description": "LACP parameters for the bond device.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface:lacp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "syntax": "attribute", "type": "object"}, {"aliases": ["link polling interval"], "anchor": "schema-oci--not_managed--node_list--interface_list--bond_interface--link_polling_interval", "description": "Link polling interval in milliseconds.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "link_polling_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["link up delay"], "anchor": "schema-oci--not_managed--node_list--interface_list--bond_interface--link_up_delay", "description": "Milliseconds wait before link is declared up.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "link_up_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-oci--not_managed--node_list--interface_list--bond_interface--name", "description": "Name for the Bond. Ex 'bond0'", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:bond_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "bond_interface", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bond devices configuration for fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.bond_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/)
- [oci.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/)
- [oci.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- oci.not_managed.node_list.interface_list.bond_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

## Direct properties

- [active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/active_backup/): complete subsection reference.

<a id="schema-oci--not_managed--node_list--interface_list--bond_interface--devices"></a>

### devices property

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/lacp/): complete subsection reference.

<a id="schema-oci--not_managed--node_list--interface_list--bond_interface--link_polling_interval"></a>

### link_polling_interval property

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="schema-oci--not_managed--node_list--interface_list--bond_interface--link_up_delay"></a>

### link_up_delay property

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="schema-oci--not_managed--node_list--interface_list--bond_interface--name"></a>

### name property

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [oci.not_managed.node_list.interface_list.bond_interface.active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/active_backup/)
- [oci.not_managed.node_list.interface_list.bond_interface.lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/bond_interface/lacp/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
