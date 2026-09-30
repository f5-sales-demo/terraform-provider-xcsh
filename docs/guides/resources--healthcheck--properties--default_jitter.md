---
page_title: "default_jitter"
subcategory: "Monitoring"
description: "default_jitter for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:b3c4579a1c4aa06170e0764e7f5abf5d69ac7583196737ebaf08c15cf52e9c09", "canonical_id": "xcsh-docs:resources:healthcheck:properties:default_jitter", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:default_jitter", "parent_id": "xcsh-docs:resources:healthcheck:reference", "path": "docs/guides/resources--healthcheck--properties--default_jitter.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_jitter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/default_jitter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_jitter for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_jitter

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Property reference](resources--healthcheck--reference.md)
- default_jitter

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

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

OneOf alternatives in this subsection:

- [default_jitter](resources--healthcheck--properties--default_jitter.md#section)
- [jitter_percent](resources--healthcheck--reference.md#schema-jitter_percent)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_jitter = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--healthcheck--reference.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
