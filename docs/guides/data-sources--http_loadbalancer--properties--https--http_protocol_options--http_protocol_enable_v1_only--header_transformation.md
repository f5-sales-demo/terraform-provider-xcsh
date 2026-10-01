---
page_title: "https.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Load Balancing"
description: "https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3044, "body_sha256": "sha256:1f927fee2338e8759bc67d7f47e15eee4fe355aff80b4fc1fe0a8a9b6c1e5192", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- [https.http_protocol_options](data-sources--http_loadbalancer--properties--https--http_protocol_options.md)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

## Direct properties

- [default_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
