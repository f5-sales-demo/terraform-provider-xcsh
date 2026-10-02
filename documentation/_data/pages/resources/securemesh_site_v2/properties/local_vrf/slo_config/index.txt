---
page_title: "local_vrf.slo_config"
subcategory: ""
description: "Site local network configuration."
xcsh_docs: {"aliases": ["local vrf slo config"], "body_bytes": 7084, "body_sha256": "sha256:4ffcdcec00590f7b206eb879313674e711214364ab6b5086ca51e7218fab7611", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf", "path": "documentation/resources/securemesh_site_v2/properties/local_vrf/slo_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "slo_config"], "schema_version": 1, "sections": [{"aliases": ["labels"], "anchor": "schema-local_vrf--slo_config--labels", "description": "Add Labels for this network, these labels can be used in firewall policy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["nameserver"], "anchor": "schema-local_vrf--slo_config--nameserver", "description": "Optional IPv4 DNS server to be used for name resolution.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["no v6 static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "no_v6_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["secondary nameserver"], "anchor": "schema-local_vrf--slo_config--secondary_nameserver", "description": "Optional Secondary IPv4 DNS server to be used for name resolution.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "secondary_nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["static routes"], "anchor": "section", "description": "Configuration parameter for static routes.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["local_vrf", "slo_config", "static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_v6_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes", "type": "requires"}], "schema_path": ["local_vrf", "slo_config", "static_v6_routes"], "syntax": "block", "type": "object"}, {"aliases": ["vip"], "anchor": "schema-local_vrf--slo_config--vip", "description": "Optional common virtual V4 IP across all nodes to be used as automatic VIP.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "vip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/slo_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Site local network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/)
- local_vrf.slo_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_vrf--slo_config--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
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

<a id="schema-local_vrf--slo_config--nameserver"></a>

### nameserver property

Type: `"string"`. Optional.

Optional IPv4 DNS server to be used for name resolution.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/no_static_routes/): complete subsection reference.

- [no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/no_v6_static_routes/): complete subsection reference.

<a id="schema-local_vrf--slo_config--secondary_nameserver"></a>

### secondary_nameserver property

Type: `"string"`. Optional.

Optional Secondary IPv4 DNS server to be used for name resolution.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/): complete subsection reference.

<a id="schema-local_vrf--slo_config--vip"></a>

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [local_vrf.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/no_static_routes/)
- [local_vrf.slo_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/no_v6_static_routes/)
- [local_vrf.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/)
- [local_vrf.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
