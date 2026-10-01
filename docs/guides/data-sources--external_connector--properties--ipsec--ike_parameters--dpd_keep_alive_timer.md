---
page_title: "ipsec.ike_parameters.dpd_keep_alive_timer"
subcategory: ""
description: "ipsec.ike_parameters.dpd_keep_alive_timer for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1917, "body_sha256": "sha256:17044149db5c46aadd0e49d7f5acc48e736c8db0218218128c28fd5f8875c24a", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "child_ids": [], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ike_parameters--dpd_keep_alive_timer.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.dpd_keep_alive_timer for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.dpd_keep_alive_timer

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dpd keep alive timer.

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

## Direct properties

<a id="schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout"></a>

### timeout property

Type: `"number"`. Computed.

Keepalive Timer. Operation timeout duration

Upstream description:

Operation timeout duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

## Next pages

- [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
