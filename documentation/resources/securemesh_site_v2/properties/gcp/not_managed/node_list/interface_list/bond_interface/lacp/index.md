---
page_title: "gcp.not_managed.node_list.interface_list.bond_interface.lacp"
subcategory: ""
description: "LACP parameters for the bond device."
xcsh_docs: {"aliases": ["gcp not managed node list interface list bond interface lacp"], "body_bytes": 2908, "body_sha256": "sha256:017225e25c4ba23c6c2935ff62653f6c57dd9ef910b543a7a9cb31f18c086f92", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface:lacp", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface", "path": "documentation/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/lacp/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-009.md", "relationships": [{"anchor": "schema-gcp--not_managed--node_list--interface_list--bond_interface--lacp--rate", "enforcement": "provider-schema", "group": "gcp.not_managed.node_list.interface_list.bond_interface.lacp:RequiredObjectAttributes:rate", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface:lacp", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "schema_version": 1, "sections": [{"aliases": ["gcp not managed node list interface list bond interface lacp rate"], "anchor": "schema-gcp--not_managed--node_list--interface_list--bond_interface--lacp--rate", "description": "Interval in seconds to transmit LACP packets.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface:lacp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "bond_interface", "lacp", "rate"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/lacp/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "LACP parameters for the bond device.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.bond_interface.lacp

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/)
- [gcp.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/)
- [gcp.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- [gcp.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-gcp--not_managed--node_list--interface_list--bond_interface--lacp--rate"></a>

### rate property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```
