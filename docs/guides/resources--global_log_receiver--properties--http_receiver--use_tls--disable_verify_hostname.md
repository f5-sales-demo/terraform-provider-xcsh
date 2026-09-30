---
page_title: "http_receiver.use_tls.disable_verify_hostname"
subcategory: ""
description: "http_receiver.use_tls.disable_verify_hostname for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1096, "body_sha256": "sha256:5690591b3c7cb15b069e73abd84bfb5127e720f0a70c72858b48fcc82effd1f9", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls", "path": "docs/guides/resources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_hostname.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "use_tls", "disable_verify_hostname"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_hostname/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.use_tls.disable_verify_hostname for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_receiver.use_tls.disable_verify_hostname

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- [http_receiver.use_tls](resources--global_log_receiver--properties--http_receiver--use_tls.md)
- http_receiver.use_tls.disable_verify_hostname

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_verify_hostname = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_receiver.use_tls](resources--global_log_receiver--properties--http_receiver--use_tls.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
