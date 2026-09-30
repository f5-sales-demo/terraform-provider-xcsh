---
page_title: "origin_servers.private_name"
subcategory: "Load Balancing"
description: "origin_servers.private_name for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 4327, "body_sha256": "sha256:864d71465c002491c8a0e165cfd5075cd91995b13f95b4da2ca2e7972fa86b80", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:inside_network", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:outside_network", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:segment", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:site_locator", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:snat_pool"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "path": "docs/guides/data-sources--origin_pool--properties--origin_servers--private_name.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "private_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/private_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.private_name for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.private_name

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- origin_servers.private_name

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with private or public DNS name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

## Direct properties

<a id="schema-origin_servers--private_name--dns_name"></a>

### dns_name property

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_network](data-sources--origin_pool--properties--origin_servers--private_name--inside_network.md): complete subsection reference.

- [outside_network](data-sources--origin_pool--properties--origin_servers--private_name--outside_network.md): complete subsection reference.

<a id="schema-origin_servers--private_name--refresh_interval"></a>

### refresh_interval property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](data-sources--origin_pool--properties--origin_servers--private_name--segment.md): complete subsection reference.

- [site_locator](data-sources--origin_pool--properties--origin_servers--private_name--site_locator.md): complete subsection reference.

- [snat_pool](data-sources--origin_pool--properties--origin_servers--private_name--snat_pool.md): complete subsection reference.

## Next pages

- [origin_servers.private_name.inside_network](data-sources--origin_pool--properties--origin_servers--private_name--inside_network.md)
- [origin_servers.private_name.outside_network](data-sources--origin_pool--properties--origin_servers--private_name--outside_network.md)
- [origin_servers.private_name.segment](data-sources--origin_pool--properties--origin_servers--private_name--segment.md)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--properties--origin_servers--private_name--site_locator.md)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--properties--origin_servers--private_name--snat_pool.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
