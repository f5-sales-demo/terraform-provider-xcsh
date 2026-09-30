---
page_title: "local_vrf.sli_config.static_v6_routes.static_routes"
subcategory: ""
description: "local_vrf.sli_config.static_v6_routes.static_routes for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 6684, "body_sha256": "sha256:9a9031ab85b6d9eb27cda8373c2d285b7fb45ef819da619c3f76aa84ee76741e", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:default_gateway", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes", "path": "docs/guides/data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_v6_routes", "static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf.sli_config.static_v6_routes.static_routes for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_vrf.sli_config.static_v6_routes.static_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [local_vrf](data-sources--securemesh_site_v2--properties--local_vrf.md)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes.md)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="section"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

## Direct properties

<a id="schema-local_vrf--sli_config--static_v6_routes--static_routes--attrs"></a>

### attrs property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--default_gateway.md): complete subsection reference.

<a id="schema-local_vrf--sli_config--static_v6_routes--static_routes--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-local_vrf--sli_config--static_v6_routes--static_routes--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface.md): complete subsection reference.

## Next pages

- [local_vrf.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--default_gateway.md)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface.md)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
