---
page_title: "custom_anonymization"
subcategory: "Security"
description: "custom_anonymization for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1645, "body_sha256": "sha256:11a8c86827ee5fb6151944d686131a78bc894e44b2c7daaf14c399697761b5b2", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "docs/guides/data-sources--app_firewall--properties--custom_anonymization.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_anonymization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- custom_anonymization

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Upstream description:

Anonymization settings which is a list of HTTP headers, parameters and cookies.

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

- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md#section)
- [default_anonymization](data-sources--app_firewall--properties--default_anonymization.md#section)
- [disable_anonymization](data-sources--app_firewall--properties--disable_anonymization.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [anonymization_config](data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md): complete subsection reference.

## Next pages

- [custom_anonymization.anonymization_config](data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md)
- [Property reference](data-sources--app_firewall--reference.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
