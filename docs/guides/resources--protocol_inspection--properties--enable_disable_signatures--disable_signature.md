---
page_title: "enable_disable_signatures.disable_signature"
subcategory: ""
description: "enable_disable_signatures.disable_signature for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1048, "body_sha256": "sha256:0c4511659933a3efdb3dca43e652b6e7bc1eba7511b960366f5131e39afa9e5f", "canonical_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "parent_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures", "path": "docs/guides/resources--protocol_inspection--properties--enable_disable_signatures--disable_signature.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_disable_signatures", "disable_signature"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_signatures/disable_signature/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_disable_signatures.disable_signature for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_disable_signatures.disable_signature

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md)
- [Property reference](resources--protocol_inspection--reference.md)
- [enable_disable_signatures](resources--protocol_inspection--properties--enable_disable_signatures.md)
- enable_disable_signatures.disable_signature

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable signature.

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
disable_signature = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_disable_signatures](resources--protocol_inspection--properties--enable_disable_signatures.md)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md)
