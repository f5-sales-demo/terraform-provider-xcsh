---
page_title: "custom_network_config.sli_config"
subcategory: ""
description: "Site local network configuration."
xcsh_docs: {"aliases": ["custom network config sli config"], "body_bytes": 8629, "body_sha256": "sha256:785c6bdf8be63e91ccc625596d83091c6f695995ced05bc018024e4ca85c62ad", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:dc_cluster_group", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_dc_cluster_group", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_static_routes", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_v6_static_routes", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "path": "documentation/resources/securemesh_site/properties/custom_network_config/sli_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "sli_config"], "schema_version": 1, "sections": [{"aliases": ["custom network config sli config dc cluster group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:dc_cluster_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_network_config--sli_config--dc_cluster_group--name", "enforcement": "provider-schema", "group": "custom_network_config.sli_config.dc_cluster_group:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:dc_cluster_group", "type": "requires"}], "schema_path": ["custom_network_config", "sli_config", "dc_cluster_group"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config sli config labels"], "anchor": "schema-custom_network_config--sli_config--labels", "description": "Add Labels for this network, these labels can be used in firewall policy.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["custom network config sli config nameserver"], "anchor": "schema-custom_network_config--sli_config--nameserver", "description": "Optional DNS V4 server IP to be used for name resolution.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config sli config no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_dc_cluster_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config no v6 static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_v6_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "no_v6_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "sli_config", "static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config sli config static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config.static_v6_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "sli_config", "static_v6_routes"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config sli config vip"], "anchor": "schema-custom_network_config--sli_config--vip", "description": "Optional common virtual V4 IP across all nodes to be used as automatic VIP.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "vip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/sli_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Site local network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- custom_network_config.sli_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
sli_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/): complete subsection reference.

<a id="schema-custom_network_config--sli_config--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-custom_network_config--sli_config--nameserver"></a>

### nameserver property

Type: `"string"`. Optional.

Optional DNS V4 server IP to be used for name resolution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_dc_cluster_group/): complete subsection reference.

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_static_routes/): complete subsection reference.

- [no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_v6_static_routes/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/): complete subsection reference.

<a id="schema-custom_network_config--sli_config--vip"></a>

### vip property

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [custom_network_config.sli_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/)
- [custom_network_config.sli_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_dc_cluster_group/)
- [custom_network_config.sli_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_static_routes/)
- [custom_network_config.sli_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_v6_static_routes/)
- [custom_network_config.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/)
- [custom_network_config.sli_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
