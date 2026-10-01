---
page_title: "custom_network_config.slo_config.labels"
subcategory: ""
description: "custom_network_config.slo_config.labels for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1172, "body_sha256": "sha256:23869f9f75c2e1875562603d1dddd6a62d320ccf8cc1b3c449061c31de8e8c64", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:labels", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--slo_config--labels.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "slo_config", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config.labels for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.labels

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.slo_config](resources--voltstack_site--properties--custom_network_config--slo_config.md)
- custom_network_config.slo_config.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this network, these labels can be used in firewall policy.

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
labels {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_network_config.slo_config](resources--voltstack_site--properties--custom_network_config--slo_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
