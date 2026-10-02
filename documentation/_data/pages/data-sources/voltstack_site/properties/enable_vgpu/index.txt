---
page_title: "enable_vgpu"
subcategory: ""
description: "Licensing configuration for NVIDIA vGPU."
xcsh_docs: {"aliases": ["enable vgpu"], "body_bytes": 3485, "body_sha256": "sha256:1c08b39c72078baa8fb909023ae3142ca4fa48e7e256e1a9f9b833d59b7a5485", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:enable_vgpu", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/enable_vgpu/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1131330000323001-2000332111122333-3330122310233213-2112201320320113-3011213113023220-1211302100220022-2031013022001131-1002210211033323", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_vgpu"], "schema_version": 1, "sections": [{"aliases": ["feature type"], "anchor": "schema-enable_vgpu--feature_type", "description": "Set feature to be enabled Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:enable_vgpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "feature_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["server address"], "anchor": "schema-enable_vgpu--server_address", "description": "Set License Server Address.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:enable_vgpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "server_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["server port"], "anchor": "schema-enable_vgpu--server_port", "description": "Set License Server port number.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:enable_vgpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "server_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/enable_vgpu/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Licensing configuration for NVIDIA vGPU.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_vgpu

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- enable_vgpu

<a id="section"></a>

Type: `"single"`. Computed.

Licensing configuration for NVIDIA vGPU.

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

## Direct properties

<a id="schema-enable_vgpu--feature_type"></a>

### feature_type property

Type: `"string"`. Computed.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Upstream description:

Set feature to be enabled

Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation
Enable NVIDIA Virtual Compute Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNLICENSED",
  "enum": [
    "UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-enable_vgpu--server_address"></a>

### server_address property

Type: `"string"`. Computed.

License Server Address. Set License Server Address.

Upstream description:

Set License Server Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="schema-enable_vgpu--server_port"></a>

### server_port property

Type: `"number"`. Computed.

License Server Port Number. Set License Server port number.

Upstream description:

Set License Server port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
