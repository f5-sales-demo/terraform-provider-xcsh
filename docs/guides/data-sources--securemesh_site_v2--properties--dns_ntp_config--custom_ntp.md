---
page_title: "dns_ntp_config.custom_ntp"
subcategory: ""
description: "dns_ntp_config.custom_ntp for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1815, "body_sha256": "sha256:75f3b9615eb4173c4cddd789e36d0f54531519e025ceb7a05132b349c63530a1", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:dns_ntp_config", "path": "docs/guides/data-sources--securemesh_site_v2--properties--dns_ntp_config--custom_ntp.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dns_ntp_config", "custom_ntp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dns_ntp_config.custom_ntp for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config.custom_ntp

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [dns_ntp_config](data-sources--securemesh_site_v2--properties--dns_ntp_config.md)
- dns_ntp_config.custom_ntp

<a id="section"></a>

Type: `"single"`. Computed.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

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

<a id="schema-dns_ntp_config--custom_ntp--ntp_servers"></a>

### ntp_servers property

Type: `["list", "string"]`. Computed.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [dns_ntp_config](data-sources--securemesh_site_v2--properties--dns_ntp_config.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
