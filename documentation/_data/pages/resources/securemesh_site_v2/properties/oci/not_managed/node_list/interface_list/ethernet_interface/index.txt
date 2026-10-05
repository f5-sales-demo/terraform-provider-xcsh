---
page_title: "oci.not_managed.node_list.interface_list.ethernet_interface"
subcategory: ""
description: "Configuration parameter for ethernet interface."
xcsh_docs: {"aliases": ["oci not managed node list interface list ethernet interface"], "body_bytes": 4532, "body_sha256": "sha256:0243399eaab7288b9e162ed2f37c68a8e641ffaf246f1acc5f0555efd4c09167", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ethernet_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ethernet_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0102100102313212-0321212212322322-1201012200002323-1233303312311112-1131103313332322-2000330320123022-2032313010223323-2133321132021131", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-014.md", "relationships": [{"anchor": "schema-oci--not_managed--node_list--interface_list--ethernet_interface--mac", "enforcement": "provider-schema", "group": "oci.not_managed.node_list.interface_list.ethernet_interface:RequiredObjectAttributes:mac", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ethernet_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ethernet_interface"], "schema_version": 1, "sections": [{"aliases": ["oci not managed node list interface list ethernet interface device"], "anchor": "schema-oci--not_managed--node_list--interface_list--ethernet_interface--device", "description": "Select an Ethernet device from the discovered interfaces to configure. Once configured, this interface will be part of this sites dataplane and can participate in the networking services configured on this site.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ethernet_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ethernet_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["oci not managed node list interface list ethernet interface mac"], "anchor": "schema-oci--not_managed--node_list--interface_list--ethernet_interface--mac", "description": "Configuration parameter for mac", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ethernet_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ethernet_interface", "mac"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ethernet_interface/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for ethernet interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/)
- [oci.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/)
- [oci.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- oci.not_managed.node_list.interface_list.ethernet_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
```

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
ethernet_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-oci--not_managed--node_list--interface_list--ethernet_interface--device"></a>

### device property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-oci--not_managed--node_list--interface_list--ethernet_interface--mac"></a>

### mac property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

## Next pages

- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
