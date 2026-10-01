---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules"
subcategory: ""
description: "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 4524, "body_sha256": "sha256:3377f7f3f893a9689abde3cb0470f67b018d1053bc218eb2812a0e3d4edcf2a7", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "path": "docs/guides/data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="section"></a>

Type: `"list"`. Computed.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Upstream description:

Network(L3/L4) routing policy rules.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [all_destinations](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_destinations.md): complete subsection reference.

- [all_sources](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_sources.md): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md): complete subsection reference.

- [http_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list.md): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md): complete subsection reference.

- [metadata](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md): complete subsection reference.

- [tls_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list.md): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_destinations.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_sources.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list.md)
- [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
