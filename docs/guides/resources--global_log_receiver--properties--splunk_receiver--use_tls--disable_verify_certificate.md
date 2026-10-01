---
page_title: "splunk_receiver.use_tls.disable_verify_certificate"
subcategory: ""
description: "splunk_receiver.use_tls.disable_verify_certificate for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:cbabede14c5650f3d9fabd7b2316535c791d821c57985a0a6c8f7c8a0e4228e7", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_certificate", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_certificate", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver--use_tls--disable_verify_certificate.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver", "use_tls", "disable_verify_certificate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/use_tls/disable_verify_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver.use_tls.disable_verify_certificate for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver.use_tls.disable_verify_certificate

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- [splunk_receiver.use_tls](resources--global_log_receiver--properties--splunk_receiver--use_tls.md)
- splunk_receiver.use_tls.disable_verify_certificate

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [splunk_receiver.use_tls](resources--global_log_receiver--properties--splunk_receiver--use_tls.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
