---
page_title: "rules.spec"
subcategory: "Security"
description: "rules.spec for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4799, "body_sha256": "sha256:4654e4c5f72a56daffc190d6ccdbeaae24632cc000acb28671909f34e17e48d4", "canonical_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_asn", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_country", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_ip", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:country_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:domain_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:http_method", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:path", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules", "path": "docs/guides/data-sources--rate_limiter_policy--properties--rules--spec.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
- [Property reference](data-sources--rate_limiter_policy--reference.md)
- [rules](data-sources--rate_limiter_policy--properties--rules.md)
- rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

## Direct properties

- [any_asn](data-sources--rate_limiter_policy--properties--rules--spec--any_asn.md): complete subsection reference.

- [any_country](data-sources--rate_limiter_policy--properties--rules--spec--any_country.md): complete subsection reference.

- [any_ip](data-sources--rate_limiter_policy--properties--rules--spec--any_ip.md): complete subsection reference.

- [apply_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--apply_rate_limiter.md): complete subsection reference.

- [asn_list](data-sources--rate_limiter_policy--properties--rules--spec--asn_list.md): complete subsection reference.

- [asn_matcher](data-sources--rate_limiter_policy--properties--rules--spec--asn_matcher.md): complete subsection reference.

- [bypass_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--bypass_rate_limiter.md): complete subsection reference.

- [country_list](data-sources--rate_limiter_policy--properties--rules--spec--country_list.md): complete subsection reference.

- [custom_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--custom_rate_limiter.md): complete subsection reference.

- [domain_matcher](data-sources--rate_limiter_policy--properties--rules--spec--domain_matcher.md): complete subsection reference.

- [headers](data-sources--rate_limiter_policy--properties--rules--spec--headers.md): complete subsection reference.

- [http_method](data-sources--rate_limiter_policy--properties--rules--spec--http_method.md): complete subsection reference.

- [ip_matcher](data-sources--rate_limiter_policy--properties--rules--spec--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](data-sources--rate_limiter_policy--properties--rules--spec--ip_prefix_list.md): complete subsection reference.

- [path](data-sources--rate_limiter_policy--properties--rules--spec--path.md): complete subsection reference.

- [segment_policy](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy.md): complete subsection reference.

## Next pages

- [rules.spec.any_asn](data-sources--rate_limiter_policy--properties--rules--spec--any_asn.md)
- [rules.spec.any_country](data-sources--rate_limiter_policy--properties--rules--spec--any_country.md)
- [rules.spec.any_ip](data-sources--rate_limiter_policy--properties--rules--spec--any_ip.md)
- [rules.spec.apply_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--apply_rate_limiter.md)
- [rules.spec.asn_list](data-sources--rate_limiter_policy--properties--rules--spec--asn_list.md)
- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--properties--rules--spec--asn_matcher.md)
- [rules.spec.bypass_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--bypass_rate_limiter.md)
- [rules.spec.country_list](data-sources--rate_limiter_policy--properties--rules--spec--country_list.md)
- [rules.spec.custom_rate_limiter](data-sources--rate_limiter_policy--properties--rules--spec--custom_rate_limiter.md)
- [rules.spec.domain_matcher](data-sources--rate_limiter_policy--properties--rules--spec--domain_matcher.md)
- [rules.spec.headers](data-sources--rate_limiter_policy--properties--rules--spec--headers.md)
- [rules.spec.http_method](data-sources--rate_limiter_policy--properties--rules--spec--http_method.md)
- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--properties--rules--spec--ip_matcher.md)
- [rules.spec.ip_prefix_list](data-sources--rate_limiter_policy--properties--rules--spec--ip_prefix_list.md)
- [rules.spec.path](data-sources--rate_limiter_policy--properties--rules--spec--path.md)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- [rules](data-sources--rate_limiter_policy--properties--rules.md)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
