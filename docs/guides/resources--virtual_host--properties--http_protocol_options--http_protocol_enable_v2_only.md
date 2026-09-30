---
page_title: "http_protocol_options.http_protocol_enable_v2_only"
subcategory: ""
description: "http_protocol_options.http_protocol_enable_v2_only for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1019, "body_sha256": "sha256:01fd3a864ecf094afcfdea08fbe87888659f32e330bee198faf75f44100fc56a", "canonical_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "path": "docs/guides/resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v2_only.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v2_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v2_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options.http_protocol_enable_v2_only for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_protocol_options.http_protocol_enable_v2_only

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [http_protocol_options](resources--virtual_host--properties--http_protocol_options.md)
- http_protocol_options.http_protocol_enable_v2_only

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_protocol_options](resources--virtual_host--properties--http_protocol_options.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
