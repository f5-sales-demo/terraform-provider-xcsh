---
page_title: "https.http_protocol_options.http_protocol_enable_v1_only"
subcategory: "Load Balancing"
description: "https.http_protocol_options.http_protocol_enable_v1_only for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1440, "body_sha256": "sha256:9e09c66615400a874762b53e74816938f0663ce114bac662b3382cebe027b43f", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.http_protocol_options.http_protocol_enable_v1_only for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- [https.http_protocol_options](data-sources--http_loadbalancer--properties--https--http_protocol_options.md)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="section"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

- [header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md): complete subsection reference.

## Next pages

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- [https.http_protocol_options](data-sources--http_loadbalancer--properties--https--http_protocol_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
