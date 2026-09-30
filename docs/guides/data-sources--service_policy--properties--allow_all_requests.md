---
page_title: "allow_all_requests"
subcategory: "Security"
description: "allow_all_requests for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1404, "body_sha256": "sha256:73aeb3742c82112d145ed4f5288cf5f0706903112e54d8480c33dd5b2e5e9f4b", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:allow_all_requests", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:allow_all_requests", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "docs/guides/data-sources--service_policy--properties--allow_all_requests.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_all_requests"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/allow_all_requests/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all_requests for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_all_requests

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- allow_all_requests

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

- [allow_all_requests](data-sources--service_policy--properties--allow_all_requests.md#section)
- [allow_list](data-sources--service_policy--properties--allow_list.md#section)
- [deny_all_requests](data-sources--service_policy--properties--deny_all_requests.md#section)
- [deny_list](data-sources--service_policy--properties--deny_list.md#section)
- [rule_list](data-sources--service_policy--properties--rule_list.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--service_policy--reference.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
