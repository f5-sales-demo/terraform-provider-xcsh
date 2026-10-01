---
page_title: "https.http_protocol_options.http_protocol_enable_v1_only"
subcategory: "Load Balancing"
description: "https.http_protocol_options.http_protocol_enable_v1_only for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1553, "body_sha256": "sha256:d3e7774d93cacf7dbc59e7f97d501c11a9859de89568df6233fc6d4eb2ede73d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options", "path": "docs/guides/resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.http_protocol_options.http_protocol_enable_v1_only for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.http_protocol_options](resources--http_loadbalancer--properties--https--http_protocol_options.md)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md): complete subsection reference.

## Next pages

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- [https.http_protocol_options](resources--http_loadbalancer--properties--https--http_protocol_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
