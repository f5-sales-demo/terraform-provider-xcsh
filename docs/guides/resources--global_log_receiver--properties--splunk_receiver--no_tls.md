---
page_title: "splunk_receiver.no_tls"
subcategory: ""
description: "splunk_receiver.no_tls for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 927, "body_sha256": "sha256:7fa76f41a5075ea07b61970bdc7978e1da3c96537c58ab56abbca9cd1d52cf57", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:no_tls", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:no_tls", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver--no_tls.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver", "no_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/no_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver.no_tls for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# splunk_receiver.no_tls

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- splunk_receiver.no_tls

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
no_tls = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
