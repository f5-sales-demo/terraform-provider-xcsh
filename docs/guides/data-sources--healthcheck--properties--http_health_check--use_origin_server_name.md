---
page_title: "http_health_check.use_origin_server_name"
subcategory: "Monitoring"
description: "http_health_check.use_origin_server_name for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1043, "body_sha256": "sha256:1960ccfef80885eb3645f5b8df0cc4011b933474b4db370fc10a693a7f49845b", "canonical_id": "xcsh-docs:data-sources:healthcheck:properties:http_health_check:use_origin_server_name", "child_ids": [], "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:properties:http_health_check:use_origin_server_name", "parent_id": "xcsh-docs:data-sources:healthcheck:properties:http_health_check", "path": "docs/guides/data-sources--healthcheck--properties--http_health_check--use_origin_server_name.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_health_check", "use_origin_server_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/properties/http_health_check/use_origin_server_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_health_check.use_origin_server_name for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check.use_origin_server_name

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md)
- [Property reference](data-sources--healthcheck--reference.md)
- [http_health_check](data-sources--healthcheck--properties--http_health_check.md)
- http_health_check.use_origin_server_name

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_health_check](data-sources--healthcheck--properties--http_health_check.md)
- [xcsh_healthcheck](../data-sources/healthcheck.md)
