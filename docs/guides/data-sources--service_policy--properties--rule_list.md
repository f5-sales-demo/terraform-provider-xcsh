---
page_title: "rule_list"
subcategory: "Security"
description: "rule_list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1482, "body_sha256": "sha256:5f22ca3526d1c8ca5a5e48188fed473b5414260daf94b0a789d610026da453d1", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "docs/guides/data-sources--service_policy--properties--rule_list.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source..

Upstream description:

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source
matching.

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

- [rules](data-sources--service_policy--properties--rule_list--rules.md): complete subsection reference.

## Next pages

- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [Property reference](data-sources--service_policy--reference.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
