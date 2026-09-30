---
page_title: "rule_list"
subcategory: "Security"
description: "rule_list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1383, "body_sha256": "sha256:6ec01f119e070c9038a414607b00250691c9b30e5ec6eaf63be0e57042c7b5de", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "docs/guides/data-sources--service_policy--properties--rule_list.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
