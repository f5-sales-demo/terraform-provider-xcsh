---
page_title: "gcp_bucket_receiver.compression"
subcategory: ""
description: "gcp_bucket_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2887, "body_sha256": "sha256:bc07051bc7b32fdd2963ce353bd9daded2986792c29201b445829bbb37583586", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_default", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_gzip", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_none"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "path": "documentation/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["gcp_bucket_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_bucket_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gcp_bucket_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/)
- gcp_bucket_receiver.compression

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

## Direct properties

- [compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_default/): complete subsection reference.

- [compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_gzip/): complete subsection reference.

- [compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_none/): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_default/)
- [gcp_bucket_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_gzip/)
- [gcp_bucket_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_none/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
