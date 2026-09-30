---
page_title: "bgp_parameters"
subcategory: ""
description: "bgp_parameters for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 3067, "body_sha256": "sha256:d36c0058a3a0e23115256c394a8e92d333cc9e4b0a44578033b7833b8603657b", "canonical_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "child_ids": ["xcsh-docs:data-sources:bgp:properties:bgp_parameters:from_site", "xcsh-docs:data-sources:bgp:properties:bgp_parameters:local_address"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "parent_id": "xcsh-docs:data-sources:bgp:reference", "path": "docs/guides/data-sources--bgp--properties--bgp_parameters.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bgp_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/bgp_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bgp_parameters for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bgp_parameters

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- bgp_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for bgp parameters.

Upstream description:

BGP parameters for the local site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

## Direct properties

<a id="schema-bgp_parameters--asn"></a>

### asn property

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](data-sources--bgp--properties--bgp_parameters--from_site.md): complete subsection reference.

<a id="schema-bgp_parameters--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Upstream description:

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

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

- [local_address](data-sources--bgp--properties--bgp_parameters--local_address.md): complete subsection reference.

## Next pages

- [bgp_parameters.from_site](data-sources--bgp--properties--bgp_parameters--from_site.md)
- [bgp_parameters.local_address](data-sources--bgp--properties--bgp_parameters--local_address.md)
- [Property reference](data-sources--bgp--reference.md)
- [xcsh_bgp](../data-sources/bgp.md)
