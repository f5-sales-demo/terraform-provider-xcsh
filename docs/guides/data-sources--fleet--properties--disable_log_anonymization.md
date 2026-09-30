---
page_title: "disable_log_anonymization"
subcategory: ""
description: "disable_log_anonymization for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1160, "body_sha256": "sha256:d56879759174c8a62266cc62397580e015db9b2f19b0af4057590a321aa27c92", "canonical_id": "xcsh-docs:data-sources:fleet:properties:disable_log_anonymization", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:disable_log_anonymization", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "docs/guides/data-sources--fleet--properties--disable_log_anonymization.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_log_anonymization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/disable_log_anonymization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_log_anonymization for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_log_anonymization

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- disable_log_anonymization

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

- [disable_log_anonymization](data-sources--fleet--properties--disable_log_anonymization.md#section)
- [enable_log_anonymization](data-sources--fleet--properties--enable_log_anonymization.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--fleet--reference.md)
- [xcsh_fleet](../data-sources/fleet.md)
