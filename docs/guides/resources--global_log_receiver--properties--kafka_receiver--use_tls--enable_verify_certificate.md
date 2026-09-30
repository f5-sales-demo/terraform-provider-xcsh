---
page_title: "kafka_receiver.use_tls.enable_verify_certificate"
subcategory: ""
description: "kafka_receiver.use_tls.enable_verify_certificate for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1146, "body_sha256": "sha256:4bf6739e5b7711f104a6f22ed759f4019666ff827a21f861c8d1b1d917653437", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "path": "docs/guides/resources--global_log_receiver--properties--kafka_receiver--use_tls--enable_verify_certificate.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver", "use_tls", "enable_verify_certificate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver.use_tls.enable_verify_certificate for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kafka_receiver.use_tls.enable_verify_certificate

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [kafka_receiver](resources--global_log_receiver--properties--kafka_receiver.md)
- [kafka_receiver.use_tls](resources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kafka_receiver.use_tls](resources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
