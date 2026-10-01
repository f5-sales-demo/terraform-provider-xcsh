---
page_title: "enable_disable_signatures"
subcategory: ""
description: "enable_disable_signatures for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1503, "body_sha256": "sha256:b2d5e1ef4b2087cd779edb5a5444c6c7dddcf4545990304b1c5242636dad44a2", "canonical_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures", "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature"], "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "docs/guides/data-sources--protocol_inspection--properties--enable_disable_signatures.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_disable_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_disable_signatures for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_signatures

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md)
- [Property reference](data-sources--protocol_inspection--reference.md)
- enable_disable_signatures

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for enable disable signatures.

Upstream description:

Enable Disable Signature Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

## Direct properties

- [disable_signature](data-sources--protocol_inspection--properties--enable_disable_signatures--disable_signature.md): complete subsection reference.

- [enable_signature](data-sources--protocol_inspection--properties--enable_disable_signatures--enable_signature.md): complete subsection reference.

## Next pages

- [enable_disable_signatures.disable_signature](data-sources--protocol_inspection--properties--enable_disable_signatures--disable_signature.md)
- [enable_disable_signatures.enable_signature](data-sources--protocol_inspection--properties--enable_disable_signatures--enable_signature.md)
- [Property reference](data-sources--protocol_inspection--reference.md)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md)
