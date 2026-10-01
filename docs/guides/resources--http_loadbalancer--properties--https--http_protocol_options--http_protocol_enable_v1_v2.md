---
page_title: "https.http_protocol_options.http_protocol_enable_v1_v2"
subcategory: "Load Balancing"
description: "https.http_protocol_options.http_protocol_enable_v1_v2 for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1245, "body_sha256": "sha256:6e5d0f48572437775eeefe2377588a9b6c39c2cdbed4d819beb45dbd4a8f2706", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_v2", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_v2", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options", "path": "docs/guides/resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_v2.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "http_protocol_options", "http_protocol_enable_v1_v2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.http_protocol_options.http_protocol_enable_v1_v2 for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options.http_protocol_enable_v1_v2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.http_protocol_options](resources--http_loadbalancer--properties--https--http_protocol_options.md)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https.http_protocol_options](resources--http_loadbalancer--properties--https--http_protocol_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
