---
page_title: "oci.not_managed.node_list"
subcategory: ""
description: "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section."
xcsh_docs: {"aliases": ["oci not managed node list"], "body_bytes": 4542, "body_sha256": "sha256:fbf00e92308b55dd76c6c09d369b5d522441533988a4fe1b6ce0d747f560bc3c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed", "path": "documentation/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list"], "schema_version": 1, "sections": [{"aliases": ["oci not managed node list hostname"], "anchor": "schema-oci--not_managed--node_list--hostname", "description": "Hostname for this Node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["oci not managed node list interface list"], "anchor": "section", "description": "Manage interfaces belonging to this node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["oci not managed node list public ip"], "anchor": "schema-oci--not_managed--node_list--public_ip", "description": "Public IP for this Node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "public_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["oci not managed node list type"], "anchor": "schema-oci--not_managed--node_list--type", "description": "Type for this Node, can be Control or Worker.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/)
- [oci.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/)
- oci.not_managed.node_list

<a id="section"></a>

Type: `"list"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-oci--not_managed--node_list--hostname"></a>

### hostname property

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/): complete subsection reference.

<a id="schema-oci--not_managed--node_list--public_ip"></a>

### public_ip property

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-oci--not_managed--node_list--type"></a>

### type property

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```
