---
page_title: "http_health_check.use_origin_server_name"
subcategory: "Monitoring"
description: "http_health_check.use_origin_server_name for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1004, "body_sha256": "sha256:321b171e7425ea13c38541d3a735603083ea345dc0b7773004d7cdf7a53f6d87", "canonical_id": "xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name", "parent_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "path": "docs/guides/resources--healthcheck--properties--http_health_check--use_origin_server_name.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_health_check", "use_origin_server_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/http_health_check/use_origin_server_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_health_check.use_origin_server_name for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_health_check.use_origin_server_name

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Property reference](resources--healthcheck--reference.md)
- [http_health_check](resources--healthcheck--properties--http_health_check.md)
- http_health_check.use_origin_server_name

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
use_origin_server_name = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_health_check](resources--healthcheck--properties--http_health_check.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
