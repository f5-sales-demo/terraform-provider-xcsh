---
page_title: "custom_storage_config.static_routes.static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["custom storage config static routes static routes"], "body_bytes": 8143, "body_sha256": "sha256:67ffdf8f014c4ee21c4b85a8ae10ce201d5ec62274a00ffaa8af430fe06adaa8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:default_gateway", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230", "registry_path": "docs/guides/resources--voltstack_site--reference--group-005.md", "relationships": [{"anchor": "schema-custom_storage_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:default_gateway", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "type": "choice"}, {"anchor": "schema-custom_storage_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--static_routes--static_routes--ip_prefixes", "enforcement": "provider-schema", "group": "custom_storage_config.static_routes.static_routes:RequiredListObjectAttributes:ip_prefixes", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "static_routes", "static_routes"], "schema_version": 1, "sections": [{"aliases": ["attrs"], "anchor": "schema-custom_storage_config--static_routes--static_routes--attrs", "description": "List of attributes that control forwarding, dynamic routing and control plane (host) reachability.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "attrs"], "syntax": "attribute", "type": "list"}, {"aliases": ["default gateway"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:default_gateway", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "default_gateway"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip address"], "anchor": "schema-custom_storage_config--static_routes--static_routes--ip_address", "description": "Exclusive with Traffic matching the IP prefixes is sent to this IP Address.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip prefixes"], "anchor": "schema-custom_storage_config--static_routes--static_routes--ip_prefixes", "description": "List of route prefixes that have common next hop and attributes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "ip_prefixes"], "syntax": "attribute", "type": "list"}, {"aliases": ["node interface"], "anchor": "section", "description": "On multinode site, this type holds the information about per node interfaces.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "static_routes", "static_routes", "node_interface"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.static_routes.static_routes

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/)
- custom_storage_config.static_routes.static_routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--static_routes--static_routes--attrs"></a>

### attrs property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/default_gateway/): complete subsection reference.

<a id="schema-custom_storage_config--static_routes--static_routes--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="schema-custom_storage_config--static_routes--static_routes--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/): complete subsection reference.

## Next pages

- [custom_storage_config.static_routes.static_routes.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/default_gateway/)
- [custom_storage_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/)
- [custom_storage_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/static_routes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
