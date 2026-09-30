---
page_title: "oci.not_managed.node_list.interface_list.static_ip"
subcategory: ""
description: "oci.not_managed.node_list.interface_list.static_ip for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3334, "body_sha256": "sha256:106ef6944af03db7e37f6c952bd8bfad121dc956b075859da6857fb021b1ee8c", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ip", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ip", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ip.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oci.not_managed.node_list.interface_list.static_ip for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# oci.not_managed.node_list.interface_list.static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [oci](data-sources--securemesh_site_v2--properties--oci.md)
- [oci.not_managed](data-sources--securemesh_site_v2--properties--oci--not_managed.md)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list.md)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- oci.not_managed.node_list.interface_list.static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="schema-oci--not_managed--node_list--interface_list--static_ip--default_gw"></a>

### default_gw property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-oci--not_managed--node_list--interface_list--static_ip--dns_server"></a>

### dns_server property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="schema-oci--not_managed--node_list--interface_list--static_ip--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

## Next pages

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
