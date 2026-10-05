---
page_title: "enabled"
subcategory: ""
description: "CloudLink ADN Network Config."
xcsh_docs: {"aliases": ["enabled"], "body_bytes": 2344, "body_sha256": "sha256:c0c0226b002e33853b78e2f5ce1ec778452c86e08b1e11056da2b37d59a1e388", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:enabled", "parent_id": "xcsh-docs:resources:cloud_link:reference", "path": "documentation/resources/cloud_link/properties/enabled/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2320130011300110-3021312032122332-3011210131012030-1310310010301231-2131130022302300-3011131030311231-0221010200100112-3232312022020312", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "schema-enabled--cloudlink_network_name", "enforcement": "provider-schema", "group": "enabled:RequiredObjectAttributes:cloudlink_network_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:enabled", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled"], "schema_version": 1, "sections": [{"aliases": ["enabled cloudlink network name"], "anchor": "schema-enabled--cloudlink_network_name", "description": "Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "document_id": "xcsh-docs:resources:cloud_link:properties:enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled", "cloudlink_network_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/enabled/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "CloudLink ADN Network Config.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- enabled

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
enabled {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enabled--cloudlink_network_name"></a>

### cloudlink_network_name property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
